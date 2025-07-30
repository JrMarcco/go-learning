package services

import (
	"fmt"
	"time"

	"gorm.io/gorm"

	"github.com/JrMarcco/go-learning/gorm/cache"
	"github.com/JrMarcco/go-learning/gorm/models"
	"github.com/JrMarcco/go-learning/gorm/types"
	"github.com/JrMarcco/go-learning/gorm/utils"
)

// UserService 用户服务
type UserService struct {
	db           *gorm.DB
	cacheManager *cache.CacheManager
}

// NewUserService 创建用户服务实例
func NewUserService(db *gorm.DB, cacheManager *cache.CacheManager) *UserService {
	return &UserService{
		db:           db,
		cacheManager: cacheManager,
	}
}

// CreateUser 创建用户
func (s *UserService) CreateUser(user *models.User) error {
	if err := s.db.Create(user).Error; err != nil {
		return utils.ErrorHandler(err)
	}

	// 清除相关缓存
	if s.cacheManager != nil {
		s.cacheManager.GetCache().FlushPattern("users:*")
	}

	return nil
}

// GetUserByID 根据ID获取用户
func (s *UserService) GetUserByID(id uint) (*models.User, error) {
	var user models.User

	// 尝试从缓存获取
	if s.cacheManager != nil {
		cacheKey := fmt.Sprintf("users:id:%d", id)
		if err := s.cacheManager.GetJSON(cacheKey, &user); err == nil {
			return &user, nil
		}
	}

	// 从数据库获取
	if err := s.db.Preload("Profile").Preload("Roles").First(&user, id).Error; err != nil {
		return nil, utils.ErrorHandler(err)
	}

	// 存储到缓存
	if s.cacheManager != nil {
		cacheKey := fmt.Sprintf("users:id:%d", id)
		s.cacheManager.SetJSON(cacheKey, &user, 10*time.Minute)
	}

	return &user, nil
}

// UpdateUser 更新用户
func (s *UserService) UpdateUser(user *models.User) error {
	if err := s.db.Save(user).Error; err != nil {
		return utils.ErrorHandler(err)
	}

	// 清除相关缓存
	if s.cacheManager != nil {
		cacheKey := fmt.Sprintf("users:id:%d", user.ID)
		s.cacheManager.GetCache().Delete(cacheKey)
		s.cacheManager.GetCache().FlushPattern("users:search:*")
	}

	return nil
}

// DeleteUser 删除用户（软删除）
func (s *UserService) DeleteUser(id uint) error {
	if err := s.db.Delete(&models.User{}, id).Error; err != nil {
		return utils.ErrorHandler(err)
	}

	// 清除相关缓存
	if s.cacheManager != nil {
		cacheKey := fmt.Sprintf("users:id:%d", id)
		s.cacheManager.GetCache().Delete(cacheKey)
		s.cacheManager.GetCache().FlushPattern("users:search:*")
	}

	return nil
}

// BuildSearchQuery 构建搜索查询
func (s *UserService) BuildSearchQuery(params *types.UserSearchParams) *gorm.DB {
	query := s.db.Model(&models.User{})

	// 姓名模糊搜索
	if params.Name != "" {
		if term, err := utils.ValidateAndSanitizeSearchTerm(params.Name); err == nil && term != "" {
			pattern := utils.BuildLikePattern(term)
			query = query.Where("name ILIKE ?", pattern)
		}
	}

	// 邮箱模糊搜索
	if params.Email != "" {
		if utils.ValidateEmail(params.Email) {
			if term, err := utils.ValidateAndSanitizeSearchTerm(params.Email); err == nil && term != "" {
				pattern := utils.BuildLikePattern(term)
				query = query.Where("email ILIKE ?", pattern)
			}
		}
	}

	// 手机号模糊搜索
	if params.Phone != "" {
		if term, err := utils.ValidateAndSanitizeSearchTerm(params.Phone); err == nil && term != "" {
			pattern := utils.BuildLikePattern(term)
			query = query.Where("phone ILIKE ?", pattern)
		}
	}

	// 状态精确搜索
	if params.Status != nil {
		query = query.Where("status = ?", *params.Status)
	}

	// 年龄范围搜索
	if params.MinAge != nil {
		query = query.Where("age >= ?", *params.MinAge)
	}
	if params.MaxAge != nil {
		query = query.Where("age <= ?", *params.MaxAge)
	}

	// 城市模糊搜索
	if params.City != "" {
		if term, err := utils.ValidateAndSanitizeSearchTerm(params.City); err == nil && term != "" {
			pattern := utils.BuildLikePattern(term)
			query = query.Where("city ILIKE ?", pattern)
		}
	}

	// 日期范围搜索
	if params.StartDate != "" {
		if startTime, err := utils.ValidateDateFormat(params.StartDate); err == nil && startTime != nil {
			query = query.Where("created_at >= ?", startTime)
		}
	}
	if params.EndDate != "" {
		if endTime, err := utils.ValidateDateFormat(params.EndDate); err == nil && endTime != nil {
			query = query.Where("created_at <= ?", endTime)
		}
	}

	// 角色搜索
	if params.RoleID != nil {
		query = query.Joins("JOIN user_roles ON users.id = user_roles.user_id").
			Where("user_roles.role_id = ?", *params.RoleID)
	}

	return query
}

// SearchUsers 搜索用户（带缓存）
func (s *UserService) SearchUsers(params *types.UserSearchParams) (*types.PaginationResult, error) {
	var users []models.User

	// 尝试从缓存获取
	if s.cacheManager != nil {
		cacheKey := utils.GenerateCacheKey("users:search", params)
		var cachedResult types.PaginationResult
		if err := s.cacheManager.GetJSON(cacheKey, &cachedResult); err == nil {
			return &cachedResult, nil
		}
	}

	// 构建查询
	query := s.BuildSearchQuery(params)

	// 执行分页查询
	result, err := utils.PaginateQueryWithPreload(query, &params.PaginationParams, &users, "Profile", "Roles")
	if err != nil {
		return nil, err
	}

	// 存储到缓存
	if s.cacheManager != nil {
		cacheKey := utils.GenerateCacheKey("users:search", params)
		s.cacheManager.SetJSON(cacheKey, result, 5*time.Minute)
	}

	return result, nil
}

// SearchUsersSimple 简单搜索用户（不带关联数据）
func (s *UserService) SearchUsersSimple(params *types.UserSearchParams) (*types.PaginationResult, error) {
	var users []models.User

	// 构建查询
	query := s.BuildSearchQuery(params)

	// 只选择基本字段
	selectFields := "id, name, email, phone, status, age, city, created_at, updated_at"

	return utils.PaginateQueryWithSelect(query, &params.PaginationParams, &users, selectFields)
}

// SearchUsersWithCursor 游标分页搜索（适用于大数据量）
func (s *UserService) SearchUsersWithCursor(searchParams *types.UserSearchParams, cursorParams *types.CursorPaginationParams) (*types.CursorPaginationResult, error) {
	// 构建基础搜索查询
	query := s.BuildSearchQuery(searchParams)

	// 转换为interface{}切片以供游标分页使用
	var users []interface{}

	return utils.CursorPaginate(query, cursorParams, &users)
}

// GetUserStats 获取用户统计信息
func (s *UserService) GetUserStats() (map[string]interface{}, error) {
	stats := make(map[string]interface{})

	// 总用户数
	var totalUsers int64
	if err := s.db.Model(&models.User{}).Count(&totalUsers).Error; err != nil {
		return nil, utils.ErrorHandler(err)
	}
	stats["total_users"] = totalUsers

	// 活跃用户数
	var activeUsers int64
	if err := s.db.Model(&models.User{}).Where("status = ?", 1).Count(&activeUsers).Error; err != nil {
		return nil, utils.ErrorHandler(err)
	}
	stats["active_users"] = activeUsers

	// 今日新增用户
	today := time.Now().Format("2006-01-02")
	var todayUsers int64
	if err := s.db.Model(&models.User{}).
		Where("DATE(created_at) = ?", today).
		Count(&todayUsers).Error; err != nil {
		return nil, utils.ErrorHandler(err)
	}
	stats["today_users"] = todayUsers

	// 按状态分组统计
	var statusStats []struct {
		Status int   `json:"status"`
		Count  int64 `json:"count"`
	}
	if err := s.db.Model(&models.User{}).
		Select("status, COUNT(*) as count").
		Group("status").
		Find(&statusStats).Error; err != nil {
		return nil, utils.ErrorHandler(err)
	}
	stats["status_distribution"] = statusStats

	return stats, nil
}

// BatchCreateUsers 批量创建用户
func (s *UserService) BatchCreateUsers(users []models.User) error {
	// 使用事务批量插入
	err := s.db.Transaction(func(tx *gorm.DB) error {
		for _, user := range users {
			if err := tx.Create(&user).Error; err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		return utils.ErrorHandler(err)
	}

	// 清除相关缓存
	if s.cacheManager != nil {
		s.cacheManager.GetCache().FlushPattern("users:*")
	}

	return nil
}

// GetUsersByRole 根据角色获取用户
func (s *UserService) GetUsersByRole(roleID uint, params *types.PaginationParams) (*types.PaginationResult, error) {
	var users []models.User

	query := s.db.Model(&models.User{}).
		Joins("JOIN user_roles ON users.id = user_roles.user_id").
		Where("user_roles.role_id = ?", roleID).
		Preload("Profile").
		Preload("Roles")

	return utils.PaginateQuery(query, params, &users)
}

// SearchUsersAdvanced 高级搜索（支持复杂条件）
func (s *UserService) SearchUsersAdvanced(conditions map[string]interface{}, params *types.PaginationParams) (*types.PaginationResult, error) {
	var users []models.User

	query := s.db.Model(&models.User{})

	// 动态构建查询条件
	for field, value := range conditions {
		switch field {
		case "name_like":
			if str, ok := value.(string); ok && str != "" {
				pattern := utils.BuildLikePattern(str)
				query = query.Where("name ILIKE ?", pattern)
			}
		case "email_exact":
			if str, ok := value.(string); ok && str != "" {
				query = query.Where("email = ?", str)
			}
		case "status_in":
			if statuses, ok := value.([]int); ok && len(statuses) > 0 {
				query = query.Where("status IN ?", statuses)
			}
		case "age_range":
			if ageRange, ok := value.(map[string]int); ok {
				if min, exists := ageRange["min"]; exists {
					query = query.Where("age >= ?", min)
				}
				if max, exists := ageRange["max"]; exists {
					query = query.Where("age <= ?", max)
				}
			}
		case "created_after":
			if t, ok := value.(time.Time); ok {
				query = query.Where("created_at > ?", t)
			}
		case "created_before":
			if t, ok := value.(time.Time); ok {
				query = query.Where("created_at < ?", t)
			}
		}
	}

	return utils.PaginateQueryWithPreload(query, params, &users, "Profile", "Roles")
}
