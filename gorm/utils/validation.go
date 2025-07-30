package utils

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"

	"gorm.io/gorm"
)

// 允许的排序字段，防止SQL注入
var AllowedOrderFields = map[string]bool{
	"id":         true,
	"name":       true,
	"email":      true,
	"phone":      true,
	"status":     true,
	"age":        true,
	"city":       true,
	"created_at": true,
	"updated_at": true,
}

// 允许的表字段映射，用于避免SQL注入
var TableFieldMapping = map[string]string{
	"id":         "users.id",
	"name":       "users.name",
	"email":      "users.email",
	"phone":      "users.phone",
	"status":     "users.status",
	"age":        "users.age",
	"city":       "users.city",
	"created_at": "users.created_at",
	"updated_at": "users.updated_at",
}

// ValidateOrderBy 验证排序字段
func ValidateOrderBy(orderBy string) error {
	if orderBy == "" {
		return nil
	}

	// 转换为小写进行检查
	field := strings.ToLower(strings.TrimSpace(orderBy))

	if !AllowedOrderFields[field] {
		return fmt.Errorf("invalid order by field: %s", orderBy)
	}

	return nil
}

// ValidateSortDirection 验证排序方向
func ValidateSortDirection(sort string) error {
	if sort == "" {
		return nil
	}

	direction := strings.ToLower(strings.TrimSpace(sort))
	if direction != "asc" && direction != "desc" {
		return fmt.Errorf("invalid sort direction: %s, must be 'asc' or 'desc'", sort)
	}

	return nil
}

// SanitizeOrderBy 清理并获取安全的排序字符串
func SanitizeOrderBy(orderBy, sort string) string {
	// 验证字段
	if err := ValidateOrderBy(orderBy); err != nil {
		orderBy = "created_at" // 使用默认字段
	}

	// 验证排序方向
	if err := ValidateSortDirection(sort); err != nil {
		sort = "desc" // 使用默认排序
	}

	// 获取映射的完整字段名
	field := strings.ToLower(strings.TrimSpace(orderBy))
	if mappedField, exists := TableFieldMapping[field]; exists {
		return fmt.Sprintf("%s %s", mappedField, strings.ToUpper(sort))
	}

	return fmt.Sprintf("%s %s", field, strings.ToUpper(sort))
}

// ValidateEmail 验证邮箱格式
func ValidateEmail(email string) bool {
	if email == "" {
		return true // 空邮箱认为有效（可选字段）
	}

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

// ValidatePhone 验证手机号格式（简单验证）
func ValidatePhone(phone string) bool {
	if phone == "" {
		return true // 空手机号认为有效（可选字段）
	}

	// 简单的手机号验证，11位数字
	phoneRegex := regexp.MustCompile(`^1[3-9]\d{9}$`)
	return phoneRegex.MatchString(phone)
}

// ValidateDateFormat 验证日期格式
func ValidateDateFormat(dateStr string) (*time.Time, error) {
	if dateStr == "" {
		return nil, nil
	}

	// 支持多种日期格式
	formats := []string{
		"2006-01-02",
		"2006-01-02 15:04:05",
		"2006/01/02",
		"2006/01/02 15:04:05",
	}

	for _, format := range formats {
		if t, err := time.Parse(format, dateStr); err == nil {
			return &t, nil
		}
	}

	return nil, fmt.Errorf("invalid date format: %s", dateStr)
}

// ValidatePageSize 验证分页大小
func ValidatePageSize(pageSize int) error {
	if pageSize < 1 {
		return errors.New("page size must be greater than 0")
	}
	if pageSize > 100 {
		return errors.New("page size must not exceed 100")
	}
	return nil
}

// ValidatePageNumber 验证页码
func ValidatePageNumber(page int) error {
	if page < 1 {
		return errors.New("page number must be greater than 0")
	}
	return nil
}

// SanitizeStringForLike 清理字符串用于LIKE查询
func SanitizeStringForLike(s string) string {
	// 移除或转义特殊字符，防止SQL注入
	s = strings.ReplaceAll(s, "%", "\\%")
	s = strings.ReplaceAll(s, "_", "\\_")
	s = strings.ReplaceAll(s, "'", "\\'")
	s = strings.ReplaceAll(s, "\"", "\\\"")
	return strings.TrimSpace(s)
}

// BuildLikePattern 构建LIKE查询模式
func BuildLikePattern(s string) string {
	if s == "" {
		return ""
	}
	sanitized := SanitizeStringForLike(s)
	return "%" + sanitized + "%"
}

// ValidateAndSanitizeSearchTerm 验证并清理搜索词
func ValidateAndSanitizeSearchTerm(term string) (string, error) {
	if term == "" {
		return "", nil
	}

	// 检查长度
	if len(term) > 100 {
		return "", errors.New("search term too long, maximum 100 characters")
	}

	// 移除首尾空格
	term = strings.TrimSpace(term)

	// 检查是否只包含空格
	if term == "" {
		return "", nil
	}

	return SanitizeStringForLike(term), nil
}

// IsValidSQLIdentifier 检查是否为有效的SQL标识符
func IsValidSQLIdentifier(identifier string) bool {
	// SQL标识符只能包含字母、数字和下划线，且不能以数字开头
	matched, _ := regexp.MatchString(`^[a-zA-Z_][a-zA-Z0-9_]*$`, identifier)
	return matched
}

// ErrorHandler GORM错误处理
func ErrorHandler(err error) error {
	if err == nil {
		return nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("record not found")
	}

	if errors.Is(err, gorm.ErrInvalidTransaction) {
		return errors.New("invalid transaction")
	}

	// 可以添加更多特定的错误处理
	return fmt.Errorf("database error: %w", err)
}
