package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/JrMarcco/go-learning/gorm/models"
	"github.com/JrMarcco/go-learning/gorm/services"
	"github.com/JrMarcco/go-learning/gorm/types"
	"github.com/JrMarcco/go-learning/gorm/utils"
)

// UserHandler 用户处理器
type UserHandler struct {
	userService *services.UserService
}

// NewUserHandler 创建用户处理器
func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// Response 统一响应结构
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// ErrorResponse 错误响应
type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

// Success 成功响应
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    0,
		Message: "success",
		Data:    data,
	})
}

// Error 错误响应
func Error(c *gin.Context, statusCode int, message string, err error) {
	response := ErrorResponse{
		Code:    statusCode,
		Message: message,
	}

	if err != nil {
		response.Error = err.Error()
	}

	c.JSON(statusCode, response)
}

// CreateUser 创建用户
// @Summary 创建用户
// @Description 创建新用户
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.User true "用户信息"
// @Success 200 {object} Response
// @Failure 400 {object} ErrorResponse
// @Router /users [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		Error(c, http.StatusBadRequest, "invalid request body", err)
		return
	}

	// 验证邮箱格式
	if !utils.ValidateEmail(user.Email) {
		Error(c, http.StatusBadRequest, "invalid email format", nil)
		return
	}

	// 验证手机号格式
	if !utils.ValidatePhone(user.Phone) {
		Error(c, http.StatusBadRequest, "invalid phone format", nil)
		return
	}

	if err := h.userService.CreateUser(&user); err != nil {
		Error(c, http.StatusInternalServerError, "failed to create user", err)
		return
	}

	Success(c, user)
}

// GetUser 获取用户详情
// @Summary 获取用户详情
// @Description 根据ID获取用户详情
// @Tags users
// @Produce json
// @Param id path int true "用户ID"
// @Success 200 {object} Response
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /users/{id} [get]
func (h *UserHandler) GetUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		Error(c, http.StatusBadRequest, "invalid user id", err)
		return
	}

	user, err := h.userService.GetUserByID(uint(id))
	if err != nil {
		Error(c, http.StatusNotFound, "user not found", err)
		return
	}

	Success(c, user)
}

// UpdateUser 更新用户
// @Summary 更新用户
// @Description 更新用户信息
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "用户ID"
// @Param user body models.User true "用户信息"
// @Success 200 {object} Response
// @Failure 400 {object} ErrorResponse
// @Router /users/{id} [put]
func (h *UserHandler) UpdateUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		Error(c, http.StatusBadRequest, "invalid user id", err)
		return
	}

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		Error(c, http.StatusBadRequest, "invalid request body", err)
		return
	}

	user.ID = uint(id)

	// 验证邮箱格式
	if !utils.ValidateEmail(user.Email) {
		Error(c, http.StatusBadRequest, "invalid email format", nil)
		return
	}

	// 验证手机号格式
	if !utils.ValidatePhone(user.Phone) {
		Error(c, http.StatusBadRequest, "invalid phone format", nil)
		return
	}

	if err := h.userService.UpdateUser(&user); err != nil {
		Error(c, http.StatusInternalServerError, "failed to update user", err)
		return
	}

	Success(c, user)
}

// DeleteUser 删除用户
// @Summary 删除用户
// @Description 删除用户（软删除）
// @Tags users
// @Produce json
// @Param id path int true "用户ID"
// @Success 200 {object} Response
// @Failure 400 {object} ErrorResponse
// @Router /users/{id} [delete]
func (h *UserHandler) DeleteUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		Error(c, http.StatusBadRequest, "invalid user id", err)
		return
	}

	if err := h.userService.DeleteUser(uint(id)); err != nil {
		Error(c, http.StatusInternalServerError, "failed to delete user", err)
		return
	}

	Success(c, gin.H{"message": "user deleted successfully"})
}

// SearchUsers 搜索用户
// @Summary 搜索用户
// @Description 根据条件搜索用户，支持分页
// @Tags users
// @Produce json
// @Param page query int false "页码" default(1)
// @Param page_size query int false "每页数量" default(10)
// @Param order_by query string false "排序字段" default(created_at)
// @Param sort query string false "排序方向" Enums(asc, desc) default(desc)
// @Param name query string false "姓名（模糊搜索）"
// @Param email query string false "邮箱（模糊搜索）"
// @Param phone query string false "手机号（模糊搜索）"
// @Param status query int false "状态"
// @Param min_age query int false "最小年龄"
// @Param max_age query int false "最大年龄"
// @Param city query string false "城市（模糊搜索）"
// @Param start_date query string false "开始日期"
// @Param end_date query string false "结束日期"
// @Param role_id query int false "角色ID"
// @Success 200 {object} Response
// @Failure 400 {object} ErrorResponse
// @Router /users/search [get]
func (h *UserHandler) SearchUsers(c *gin.Context) {
	var params types.UserSearchParams

	// 绑定查询参数
	if err := c.ShouldBindQuery(&params); err != nil {
		Error(c, http.StatusBadRequest, "invalid query parameters", err)
		return
	}

	// 验证分页参数
	if err := utils.ValidatePageNumber(params.Page); err != nil {
		Error(c, http.StatusBadRequest, "invalid page number", err)
		return
	}

	if err := utils.ValidatePageSize(params.PageSize); err != nil {
		Error(c, http.StatusBadRequest, "invalid page size", err)
		return
	}

	// 验证排序参数
	if err := utils.ValidateOrderBy(params.OrderBy); err != nil {
		Error(c, http.StatusBadRequest, "invalid order by field", err)
		return
	}

	if err := utils.ValidateSortDirection(params.Sort); err != nil {
		Error(c, http.StatusBadRequest, "invalid sort direction", err)
		return
	}

	// 执行搜索
	result, err := h.userService.SearchUsers(&params)
	if err != nil {
		Error(c, http.StatusInternalServerError, "failed to search users", err)
		return
	}

	Success(c, result)
}

// SearchUsersSimple 简单搜索用户（不包含关联数据）
// @Summary 简单搜索用户
// @Description 简单搜索用户，不包含关联数据，响应更快
// @Tags users
// @Produce json
// @Param page query int false "页码" default(1)
// @Param page_size query int false "每页数量" default(10)
// @Param order_by query string false "排序字段" default(created_at)
// @Param sort query string false "排序方向" Enums(asc, desc) default(desc)
// @Param name query string false "姓名（模糊搜索）"
// @Param email query string false "邮箱（模糊搜索）"
// @Param status query int false "状态"
// @Success 200 {object} Response
// @Failure 400 {object} ErrorResponse
// @Router /users/search/simple [get]
func (h *UserHandler) SearchUsersSimple(c *gin.Context) {
	var params types.UserSearchParams

	if err := c.ShouldBindQuery(&params); err != nil {
		Error(c, http.StatusBadRequest, "invalid query parameters", err)
		return
	}

	result, err := h.userService.SearchUsersSimple(&params)
	if err != nil {
		Error(c, http.StatusInternalServerError, "failed to search users", err)
		return
	}

	Success(c, result)
}

// SearchUsersWithCursor 游标分页搜索
// @Summary 游标分页搜索
// @Description 使用游标分页搜索用户，适用于大数据量场景
// @Tags users
// @Produce json
// @Param limit query int false "限制数量" default(10)
// @Param cursor query string false "游标"
// @Param direction query string false "方向" Enums(next, prev) default(next)
// @Param name query string false "姓名（模糊搜索）"
// @Param email query string false "邮箱（模糊搜索）"
// @Param status query int false "状态"
// @Success 200 {object} Response
// @Failure 400 {object} ErrorResponse
// @Router /users/search/cursor [get]
func (h *UserHandler) SearchUsersWithCursor(c *gin.Context) {
	var searchParams types.UserSearchParams
	var cursorParams types.CursorPaginationParams

	if err := c.ShouldBindQuery(&searchParams); err != nil {
		Error(c, http.StatusBadRequest, "invalid search parameters", err)
		return
	}

	if err := c.ShouldBindQuery(&cursorParams); err != nil {
		Error(c, http.StatusBadRequest, "invalid cursor parameters", err)
		return
	}

	result, err := h.userService.SearchUsersWithCursor(&searchParams, &cursorParams)
	if err != nil {
		Error(c, http.StatusInternalServerError, "failed to search users with cursor", err)
		return
	}

	Success(c, result)
}

// GetUserStats 获取用户统计信息
// @Summary 获取用户统计信息
// @Description 获取用户相关的统计信息
// @Tags users
// @Produce json
// @Success 200 {object} Response
// @Router /users/stats [get]
func (h *UserHandler) GetUserStats(c *gin.Context) {
	stats, err := h.userService.GetUserStats()
	if err != nil {
		Error(c, http.StatusInternalServerError, "failed to get user stats", err)
		return
	}

	Success(c, stats)
}

// BatchCreateUsers 批量创建用户
// @Summary 批量创建用户
// @Description 批量创建多个用户
// @Tags users
// @Accept json
// @Produce json
// @Param users body []models.User true "用户列表"
// @Success 200 {object} Response
// @Failure 400 {object} ErrorResponse
// @Router /users/batch [post]
func (h *UserHandler) BatchCreateUsers(c *gin.Context) {
	var users []models.User

	if err := c.ShouldBindJSON(&users); err != nil {
		Error(c, http.StatusBadRequest, "invalid request body", err)
		return
	}

	// 验证用户数量限制
	if len(users) > 100 {
		Error(c, http.StatusBadRequest, "too many users, maximum 100 allowed", nil)
		return
	}

	// 验证每个用户的数据
	for i, user := range users {
		if !utils.ValidateEmail(user.Email) {
			Error(c, http.StatusBadRequest, "invalid email format at index "+strconv.Itoa(i), nil)
			return
		}

		if !utils.ValidatePhone(user.Phone) {
			Error(c, http.StatusBadRequest, "invalid phone format at index "+strconv.Itoa(i), nil)
			return
		}
	}

	if err := h.userService.BatchCreateUsers(users); err != nil {
		Error(c, http.StatusInternalServerError, "failed to batch create users", err)
		return
	}

	Success(c, gin.H{
		"message": "users created successfully",
		"count":   len(users),
	})
}

// GetUsersByRole 根据角色获取用户
// @Summary 根据角色获取用户
// @Description 获取指定角色的用户列表
// @Tags users
// @Produce json
// @Param role_id path int true "角色ID"
// @Param page query int false "页码" default(1)
// @Param page_size query int false "每页数量" default(10)
// @Param order_by query string false "排序字段" default(created_at)
// @Param sort query string false "排序方向" Enums(asc, desc) default(desc)
// @Success 200 {object} Response
// @Failure 400 {object} ErrorResponse
// @Router /users/role/{role_id} [get]
func (h *UserHandler) GetUsersByRole(c *gin.Context) {
	roleIDStr := c.Param("role_id")
	roleID, err := strconv.ParseUint(roleIDStr, 10, 32)
	if err != nil {
		Error(c, http.StatusBadRequest, "invalid role id", err)
		return
	}

	var params types.PaginationParams
	if err := c.ShouldBindQuery(&params); err != nil {
		Error(c, http.StatusBadRequest, "invalid query parameters", err)
		return
	}

	result, err := h.userService.GetUsersByRole(uint(roleID), &params)
	if err != nil {
		Error(c, http.StatusInternalServerError, "failed to get users by role", err)
		return
	}

	Success(c, result)
}

// SearchUsersAdvanced 高级搜索
// @Summary 高级搜索用户
// @Description 支持复杂条件的高级搜索
// @Tags users
// @Accept json
// @Produce json
// @Param conditions body map[string]interface{} true "搜索条件"
// @Param page query int false "页码" default(1)
// @Param page_size query int false "每页数量" default(10)
// @Param order_by query string false "排序字段" default(created_at)
// @Param sort query string false "排序方向" Enums(asc, desc) default(desc)
// @Success 200 {object} Response
// @Failure 400 {object} ErrorResponse
// @Router /users/search/advanced [post]
func (h *UserHandler) SearchUsersAdvanced(c *gin.Context) {
	var conditions map[string]interface{}
	if err := c.ShouldBindJSON(&conditions); err != nil {
		Error(c, http.StatusBadRequest, "invalid request body", err)
		return
	}

	var params types.PaginationParams
	if err := c.ShouldBindQuery(&params); err != nil {
		Error(c, http.StatusBadRequest, "invalid query parameters", err)
		return
	}

	// 处理时间类型的条件
	if createdAfter, exists := conditions["created_after"]; exists {
		if timeStr, ok := createdAfter.(string); ok {
			if t, err := time.Parse("2006-01-02 15:04:05", timeStr); err == nil {
				conditions["created_after"] = t
			}
		}
	}

	if createdBefore, exists := conditions["created_before"]; exists {
		if timeStr, ok := createdBefore.(string); ok {
			if t, err := time.Parse("2006-01-02 15:04:05", timeStr); err == nil {
				conditions["created_before"] = t
			}
		}
	}

	result, err := h.userService.SearchUsersAdvanced(conditions, &params)
	if err != nil {
		Error(c, http.StatusInternalServerError, "failed to search users", err)
		return
	}

	Success(c, result)
}

// RegisterRoutes 注册路由
func (h *UserHandler) RegisterRoutes(router *gin.RouterGroup) {
	users := router.Group("/users")
	{
		users.POST("", h.CreateUser)
		users.GET("/:id", h.GetUser)
		users.PUT("/:id", h.UpdateUser)
		users.DELETE("/:id", h.DeleteUser)

		users.GET("/search", h.SearchUsers)
		users.GET("/search/simple", h.SearchUsersSimple)
		users.GET("/search/cursor", h.SearchUsersWithCursor)
		users.POST("/search/advanced", h.SearchUsersAdvanced)

		users.GET("/stats", h.GetUserStats)
		users.POST("/batch", h.BatchCreateUsers)
		users.GET("/role/:role_id", h.GetUsersByRole)
	}
}
