package utils

// IntPtr 返回int的指针
func IntPtr(i int) *int {
	return &i
}

// UintPtr 返回uint的指针
func UintPtr(u uint) *uint {
	return &u
}

// StringPtr 返回string的指针
func StringPtr(s string) *string {
	return &s
}

// BoolPtr 返回bool的指针
func BoolPtr(b bool) *bool {
	return &b
}

// Float64Ptr 返回float64的指针
func Float64Ptr(f float64) *float64 {
	return &f
}

// SafeString 安全地从指针获取字符串值，如果为nil则返回默认值
func SafeString(ptr *string, defaultValue string) string {
	if ptr == nil {
		return defaultValue
	}
	return *ptr
}

// SafeInt 安全地从指针获取int值，如果为nil则返回默认值
func SafeInt(ptr *int, defaultValue int) int {
	if ptr == nil {
		return defaultValue
	}
	return *ptr
}

// SafeUint 安全地从指针获取uint值，如果为nil则返回默认值
func SafeUint(ptr *uint, defaultValue uint) uint {
	if ptr == nil {
		return defaultValue
	}
	return *ptr
}

// SafeBool 安全地从指针获取bool值，如果为nil则返回默认值
func SafeBool(ptr *bool, defaultValue bool) bool {
	if ptr == nil {
		return defaultValue
	}
	return *ptr
}

// Contains 检查字符串切片是否包含指定字符串
func Contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

// ContainsInt 检查int切片是否包含指定int
func ContainsInt(slice []int, item int) bool {
	for _, i := range slice {
		if i == item {
			return true
		}
	}
	return false
}

// RemoveDuplicates 移除字符串切片中的重复项
func RemoveDuplicates(slice []string) []string {
	keys := make(map[string]bool)
	var result []string

	for _, item := range slice {
		if !keys[item] {
			keys[item] = true
			result = append(result, item)
		}
	}

	return result
}

// Min 返回两个int中的较小值
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Max 返回两个int中的较大值
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Clamp 将值限制在指定范围内
func Clamp(value, min, max int) int {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}
