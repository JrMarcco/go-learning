package utils

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"gorm.io/gorm"

	"github.com/JrMarcco/go-learning/gorm/types"
)

// PaginateQuery 通用分页查询函数
func PaginateQuery(db *gorm.DB, params *types.PaginationParams, result interface{}) (*types.PaginationResult, error) {
	// 设置默认值
	params.SetDefaults()

	// 验证参数
	if err := ValidatePageNumber(params.Page); err != nil {
		return nil, err
	}
	if err := ValidatePageSize(params.PageSize); err != nil {
		return nil, err
	}
	if err := ValidateOrderBy(params.OrderBy); err != nil {
		return nil, err
	}
	if err := ValidateSortDirection(params.Sort); err != nil {
		return nil, err
	}

	var total int64

	// 计算总数（克隆查询以避免影响原查询）
	countDB := db.Session(&gorm.Session{})
	if err := countDB.Count(&total).Error; err != nil {
		return nil, ErrorHandler(fmt.Errorf("failed to count records: %w", err))
	}

	// 如果没有数据，直接返回
	if total == 0 {
		return types.NewPaginationResult(result, 0, params), nil
	}

	// 构建安全的排序字符串
	orderBy := SanitizeOrderBy(params.OrderBy, params.Sort)

	// 执行分页查询
	offset := params.GetOffset()
	if err := db.Order(orderBy).
		Limit(params.PageSize).
		Offset(offset).
		Find(result).Error; err != nil {
		return nil, ErrorHandler(fmt.Errorf("failed to fetch records: %w", err))
	}

	return types.NewPaginationResult(result, total, params), nil
}

// PaginateQueryWithPreload 带预加载的分页查询
func PaginateQueryWithPreload(db *gorm.DB, params *types.PaginationParams, result interface{}, preloads ...string) (*types.PaginationResult, error) {
	// 应用预加载
	for _, preload := range preloads {
		db = db.Preload(preload)
	}

	return PaginateQuery(db, params, result)
}

// PaginateQueryWithSelect 带字段选择的分页查询
func PaginateQueryWithSelect(db *gorm.DB, params *types.PaginationParams, result interface{}, selectFields string) (*types.PaginationResult, error) {
	if selectFields != "" {
		db = db.Select(selectFields)
	}

	return PaginateQuery(db, params, result)
}

// CursorInfo 游标信息
type CursorInfo struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
}

// EncodeCursor 编码游标
func EncodeCursor(id uint, createdAt time.Time) string {
	cursor := CursorInfo{
		ID:        id,
		CreatedAt: createdAt,
	}

	data, _ := json.Marshal(cursor)
	return base64.URLEncoding.EncodeToString(data)
}

// DecodeCursor 解码游标
func DecodeCursor(cursorStr string) (*CursorInfo, error) {
	if cursorStr == "" {
		return nil, nil
	}

	data, err := base64.URLEncoding.DecodeString(cursorStr)
	if err != nil {
		return nil, fmt.Errorf("invalid cursor format: %w", err)
	}

	var cursor CursorInfo
	if err := json.Unmarshal(data, &cursor); err != nil {
		return nil, fmt.Errorf("invalid cursor data: %w", err)
	}

	return &cursor, nil
}

// CursorPaginate 游标分页查询（适用于大数据量）
func CursorPaginate(db *gorm.DB, params *types.CursorPaginationParams, result interface{}) (*types.CursorPaginationResult, error) {
	params.SetDefaults()

	// 解析游标
	cursor, err := DecodeCursor(params.Cursor)
	if err != nil {
		return nil, err
	}

	query := db

	// 根据游标和方向设置查询条件
	if cursor != nil {
		if params.Direction == "next" {
			// 获取比当前游标更新的记录
			query = query.Where("(created_at < ? OR (created_at = ? AND id < ?))",
				cursor.CreatedAt, cursor.CreatedAt, cursor.ID)
		} else {
			// 获取比当前游标更旧的记录
			query = query.Where("(created_at > ? OR (created_at = ? AND id > ?))",
				cursor.CreatedAt, cursor.CreatedAt, cursor.ID)
		}
	}

	// 设置排序和限制
	if params.Direction == "next" {
		query = query.Order("created_at DESC, id DESC")
	} else {
		query = query.Order("created_at ASC, id ASC")
	}

	// 多取一条记录来判断是否还有更多数据
	limit := params.Limit + 1
	query = query.Limit(limit)

	// 执行查询
	if err := query.Find(result).Error; err != nil {
		return nil, ErrorHandler(fmt.Errorf("failed to fetch records: %w", err))
	}

	// 处理结果
	resultSlice, ok := result.(*[]interface{})
	if !ok {
		return nil, fmt.Errorf("result must be a pointer to slice")
	}

	hasMore := len(*resultSlice) > params.Limit
	if hasMore {
		// 移除多余的记录
		*resultSlice = (*resultSlice)[:params.Limit]
	}

	paginationResult := &types.CursorPaginationResult{
		Data:    *resultSlice,
		HasNext: false,
		HasPrev: cursor != nil,
	}

	// 设置游标
	if len(*resultSlice) > 0 {
		if params.Direction == "next" {
			paginationResult.HasNext = hasMore
			if hasMore {
				// 设置下一页游标（最后一条记录）
				lastRecord := (*resultSlice)[len(*resultSlice)-1]
				if record, ok := lastRecord.(map[string]interface{}); ok {
					if id, ok := record["id"].(uint); ok {
						if createdAt, ok := record["created_at"].(time.Time); ok {
							paginationResult.NextCursor = EncodeCursor(id, createdAt)
						}
					}
				}
			}
		} else {
			paginationResult.HasPrev = hasMore
			if hasMore {
				// 设置上一页游标（第一条记录）
				firstRecord := (*resultSlice)[0]
				if record, ok := firstRecord.(map[string]interface{}); ok {
					if id, ok := record["id"].(uint); ok {
						if createdAt, ok := record["created_at"].(time.Time); ok {
							paginationResult.PrevCursor = EncodeCursor(id, createdAt)
						}
					}
				}
			}
		}
	}

	return paginationResult, nil
}

// GenerateCacheKey 生成缓存键
func GenerateCacheKey(prefix string, params interface{}) string {
	data, _ := json.Marshal(params)
	return fmt.Sprintf("%s:%x", prefix, data)
}

// ParsePaginationFromQuery 从查询参数解析分页信息
func ParsePaginationFromQuery(query map[string][]string) *types.PaginationParams {
	params := &types.PaginationParams{}

	if page := query["page"]; len(page) > 0 {
		if p, err := strconv.Atoi(page[0]); err == nil {
			params.Page = p
		}
	}

	if pageSize := query["page_size"]; len(pageSize) > 0 {
		if ps, err := strconv.Atoi(pageSize[0]); err == nil {
			params.PageSize = ps
		}
	}

	if orderBy := query["order_by"]; len(orderBy) > 0 {
		params.OrderBy = orderBy[0]
	}

	if sort := query["sort"]; len(sort) > 0 {
		params.Sort = sort[0]
	}

	params.SetDefaults()
	return params
}

// CalculateOffset 计算数据库查询偏移量
func CalculateOffset(page, pageSize int) int {
	if page <= 0 {
		page = 1
	}
	return (page - 1) * pageSize
}

// CalculateTotalPages 计算总页数
func CalculateTotalPages(total int64, pageSize int) int {
	if pageSize <= 0 {
		return 0
	}

	totalPages := int(total) / pageSize
	if int(total)%pageSize > 0 {
		totalPages++
	}

	return totalPages
}
