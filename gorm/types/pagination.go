package types

import (
	"math"
	"strings"
)

// PaginationParams 分页参数
type PaginationParams struct {
	Page     int    `json:"page" form:"page" validate:"min=1"`                   // 页码，从1开始
	PageSize int    `json:"page_size" form:"page_size" validate:"min=1,max=100"` // 每页数量
	OrderBy  string `json:"order_by" form:"order_by"`                            // 排序字段
	Sort     string `json:"sort" form:"sort" validate:"oneof=asc desc"`          // 排序方向
}

// SetDefaults 设置默认值
func (p *PaginationParams) SetDefaults() {
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.PageSize <= 0 {
		p.PageSize = 10
	}
	if p.PageSize > 100 {
		p.PageSize = 100
	}
	if p.Sort == "" {
		p.Sort = "desc"
	}
	if p.OrderBy == "" {
		p.OrderBy = "created_at"
	}
}

// GetOffset 计算偏移量
func (p *PaginationParams) GetOffset() int {
	return (p.Page - 1) * p.PageSize
}

// PaginationResult 分页结果
type PaginationResult struct {
	Data       interface{} `json:"data"`        // 数据列表
	Total      int64       `json:"total"`       // 总记录数
	Page       int         `json:"page"`        // 当前页
	PageSize   int         `json:"page_size"`   // 每页数量
	TotalPages int         `json:"total_pages"` // 总页数
	HasNext    bool        `json:"has_next"`    // 是否有下一页
	HasPrev    bool        `json:"has_prev"`    // 是否有上一页
}

// NewPaginationResult 创建分页结果
func NewPaginationResult(data interface{}, total int64, params *PaginationParams) *PaginationResult {
	totalPages := int(math.Ceil(float64(total) / float64(params.PageSize)))

	return &PaginationResult{
		Data:       data,
		Total:      total,
		Page:       params.Page,
		PageSize:   params.PageSize,
		TotalPages: totalPages,
		HasNext:    params.Page < totalPages,
		HasPrev:    params.Page > 1,
	}
}

// UserSearchParams 用户搜索参数
type UserSearchParams struct {
	PaginationParams
	Name      string `json:"name" form:"name"`             // 姓名模糊搜索
	Email     string `json:"email" form:"email"`           // 邮箱模糊搜索
	Phone     string `json:"phone" form:"phone"`           // 手机号模糊搜索
	Status    *int   `json:"status" form:"status"`         // 状态精确搜索
	MinAge    *int   `json:"min_age" form:"min_age"`       // 最小年龄
	MaxAge    *int   `json:"max_age" form:"max_age"`       // 最大年龄
	City      string `json:"city" form:"city"`             // 城市模糊搜索
	StartDate string `json:"start_date" form:"start_date"` // 开始日期
	EndDate   string `json:"end_date" form:"end_date"`     // 结束日期
	RoleID    *uint  `json:"role_id" form:"role_id"`       // 角色ID
}

// CursorPaginationParams 游标分页参数（适用于大数据量）
type CursorPaginationParams struct {
	Limit     int    `json:"limit" form:"limit" validate:"min=1,max=100"`
	Cursor    string `json:"cursor" form:"cursor"`                                  // 游标（base64编码的时间戳或ID）
	Direction string `json:"direction" form:"direction" validate:"oneof=next prev"` // 方向
}

// SetDefaults 设置默认值
func (p *CursorPaginationParams) SetDefaults() {
	if p.Limit <= 0 {
		p.Limit = 10
	}
	if p.Limit > 100 {
		p.Limit = 100
	}
	if p.Direction == "" {
		p.Direction = "next"
	}
}

// CursorPaginationResult 游标分页结果
type CursorPaginationResult struct {
	Data       interface{} `json:"data"`
	NextCursor string      `json:"next_cursor,omitempty"` // 下一页游标
	PrevCursor string      `json:"prev_cursor,omitempty"` // 上一页游标
	HasNext    bool        `json:"has_next"`              // 是否有下一页
	HasPrev    bool        `json:"has_prev"`              // 是否有上一页
}

// SortDirection 排序方向
type SortDirection string

const (
	SortAsc  SortDirection = "asc"
	SortDesc SortDirection = "desc"
)

// IsValid 验证排序方向是否有效
func (s SortDirection) IsValid() bool {
	return s == SortAsc || s == SortDesc
}

// String 转换为字符串
func (s SortDirection) String() string {
	return string(s)
}

// NormalizeSortDirection 标准化排序方向
func NormalizeSortDirection(sort string) SortDirection {
	switch strings.ToLower(sort) {
	case "asc", "ascending":
		return SortAsc
	case "desc", "descending":
		return SortDesc
	default:
		return SortDesc
	}
}
