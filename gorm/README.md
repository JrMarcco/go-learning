# GORM 分页搜索查询最佳实践

这是一个完整的 GORM 分页搜索查询示例项目，展示了如何构建高效、安全、可维护的分页搜索功能。

## 项目结构

```
gorm/
├── models/           # 数据模型
│   └── user.go      # 用户模型定义
├── types/           # 类型定义
│   └── pagination.go # 分页和搜索相关类型
├── utils/           # 工具函数
│   ├── validation.go # 验证和安全工具
│   └── pagination.go # 分页工具函数
├── cache/           # 缓存层
│   └── cache.go     # 缓存接口和实现
├── services/        # 服务层
│   └── user_service.go # 用户业务逻辑
├── handlers/        # HTTP处理器
│   └── user_handler.go # 用户API处理器
├── example/         # 示例程序
│   └── main.go      # 完整示例应用
└── README.md        # 项目文档
```

## 核心功能特性

### 1. 安全分页查询
- ✅ SQL注入防护
- ✅ 参数验证和清理
- ✅ 安全的排序字段映射
- ✅ 分页参数限制

### 2. 多样化搜索支持
- ✅ 基础分页搜索
- ✅ 模糊搜索（姓名、邮箱、城市等）
- ✅ 精确搜索（状态、角色等）
- ✅ 范围搜索（年龄、日期等）
- ✅ 游标分页（适用于大数据量）
- ✅ 高级条件搜索

### 3. 性能优化
- ✅ 数据库索引优化
- ✅ 预加载关联数据
- ✅ 选择性字段查询
- ✅ Redis缓存支持
- ✅ 连接池配置

### 4. 开发友好
- ✅ 统一错误处理
- ✅ 结构化响应格式
- ✅ 完整的API文档注释
- ✅ 可配置的中间件
- ✅ 详细的示例代码

## 快速开始

### 1. 环境要求

- Go 1.24+
- MySQL 8.0+
- Redis 6.0+ (可选，用于缓存)

### 2. 安装依赖

```bash
cd gorm
go mod tidy
```

### 3. 数据库配置

创建MySQL数据库：

```sql
CREATE DATABASE gorm_example CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

### 4. 运行示例

```bash
go run example/main.go
```

### 5. 测试API

访问 http://localhost:8080/api/v1/examples 查看所有可用的API端点示例。

## API 端点说明

### 基础CRUD操作

| 方法 | 端点 | 描述 |
|------|------|------|
| POST | `/api/v1/users` | 创建用户 |
| GET | `/api/v1/users/:id` | 获取用户详情 |
| PUT | `/api/v1/users/:id` | 更新用户 |
| DELETE | `/api/v1/users/:id` | 删除用户 |

### 搜索和分页

| 方法 | 端点 | 描述 |
|------|------|------|
| GET | `/api/v1/users/search` | 基础搜索（含关联数据） |
| GET | `/api/v1/users/search/simple` | 简单搜索（仅基本字段） |
| GET | `/api/v1/users/search/cursor` | 游标分页搜索 |
| POST | `/api/v1/users/search/advanced` | 高级条件搜索 |

### 其他功能

| 方法 | 端点 | 描述 |
|------|------|------|
| GET | `/api/v1/users/stats` | 用户统计信息 |
| POST | `/api/v1/users/batch` | 批量创建用户 |
| GET | `/api/v1/users/role/:role_id` | 按角色获取用户 |

## 使用示例

### 1. 基础分页搜索

```bash
# 搜索姓名包含"张"的用户，按创建时间降序排列
curl "http://localhost:8080/api/v1/users/search?name=张&page=1&page_size=10&order_by=created_at&sort=desc"
```

### 2. 复合条件搜索

```bash
# 搜索活跃状态、年龄20-30岁、在北京的用户
curl "http://localhost:8080/api/v1/users/search?status=1&min_age=20&max_age=30&city=北京"
```

### 3. 游标分页

```bash
# 首次请求
curl "http://localhost:8080/api/v1/users/search/cursor?limit=10&direction=next"

# 下一页（使用返回的next_cursor）
curl "http://localhost:8080/api/v1/users/search/cursor?limit=10&direction=next&cursor=eyJpZCI6MTAsImNyZWF0ZWRfYXQiOiIyMDI0LTAxLTE1VDA5OjMwOjAwWiJ9"
```

### 4. 高级搜索

```bash
curl -X POST "http://localhost:8080/api/v1/users/search/advanced" \
  -H "Content-Type: application/json" \
  -d '{
    "name_like": "张",
    "status_in": [1, 2],
    "age_range": {"min": 20, "max": 30},
    "created_after": "2024-01-01 00:00:00"
  }'
```

### 5. 批量创建用户

```bash
curl -X POST "http://localhost:8080/api/v1/users/batch" \
  -H "Content-Type: application/json" \
  -d '[
    {
      "name": "测试用户1",
      "email": "test1@example.com",
      "phone": "13800138001",
      "status": 1,
      "age": 25,
      "city": "北京"
    },
    {
      "name": "测试用户2",
      "email": "test2@example.com",
      "phone": "13800138002",
      "status": 1,
      "age": 28,
      "city": "上海"
    }
  ]'
```

## 响应格式

### 成功响应

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "data": [...],
    "total": 100,
    "page": 1,
    "page_size": 10,
    "total_pages": 10,
    "has_next": true,
    "has_prev": false
  }
}
```

### 错误响应

```json
{
  "code": 400,
  "message": "invalid query parameters",
  "error": "page size must not exceed 100"
}
```

## 性能优化建议

### 1. 数据库索引

确保为常用的搜索字段创建索引：

```sql
-- 用户表索引
CREATE INDEX idx_users_name ON users(name);
CREATE INDEX idx_users_status ON users(status);
CREATE INDEX idx_users_city ON users(city);
CREATE INDEX idx_users_age ON users(age);
CREATE INDEX idx_users_created_at ON users(created_at);

-- 复合索引（针对常用的搜索组合）
CREATE INDEX idx_users_status_city ON users(status, city);
CREATE INDEX idx_users_age_status ON users(age, status);
```

### 2. 缓存策略

- **查询缓存**: 对频繁的搜索查询结果进行缓存（5分钟）
- **用户缓存**: 对单个用户信息进行缓存（10分钟）
- **统计缓存**: 对统计信息进行缓存（30分钟）

### 3. 分页策略选择

- **偏移分页**: 适用于总数据量 < 10万，用户需要跳页
- **游标分页**: 适用于总数据量 > 10万，用户顺序浏览

### 4. 查询优化

- 使用 `Select` 只查询需要的字段
- 合理使用 `Preload` 避免 N+1 查询
- 对大表使用索引覆盖查询

## 安全注意事项

### 1. SQL注入防护

- ✅ 使用参数化查询
- ✅ 验证排序字段白名单
- ✅ 清理用户输入

### 2. 参数验证

- ✅ 页码和页大小限制
- ✅ 搜索词长度限制
- ✅ 特殊字符转义

### 3. 访问控制

- 🔄 添加认证中间件
- 🔄 实现权限检查
- 🔄 记录审计日志

## 扩展功能

### 1. 搜索高亮

```go
// 在服务层添加搜索词高亮功能
func (s *UserService) HighlightSearchTerms(users []User, searchTerm string) []User {
    // 实现搜索词高亮逻辑
}
```

### 2. 搜索建议

```go
// 添加搜索建议功能
func (s *UserService) GetSearchSuggestions(partial string) []string {
    // 实现搜索建议逻辑
}
```

### 3. 导出功能

```go
// 添加搜索结果导出功能
func (s *UserService) ExportSearchResults(params *SearchParams) ([]byte, error) {
    // 实现CSV/Excel导出
}
```

## 故障排查

### 常见问题

1. **查询缓慢**
   - 检查是否缺少数据库索引
   - 优化查询条件
   - 减少返回字段

2. **内存占用高**
   - 限制分页大小
   - 使用游标分页
   - 检查缓存策略

3. **缓存失效**
   - 检查Redis连接
   - 验证缓存键格式
   - 调整缓存过期时间

### 监控指标

- 查询响应时间
- 缓存命中率
- 数据库连接数
- 内存使用情况
