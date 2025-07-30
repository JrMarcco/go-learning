# GORM 分页搜索查询 - 快速开始

## 1. 环境准备

### 使用Docker Compose（推荐）

```bash
# 在gorm目录下启动MySQL和Redis
docker-compose up -d

# 查看服务状态
docker-compose ps

# 停止服务
docker-compose down
```

### 手动安装

- MySQL 8.0+
- Redis 6.0+ (可选)
- Go 1.24+

## 2. 安装依赖

```bash
cd gorm
go mod tidy
```

## 3. 快速运行

### 方式一：运行完整的Web API示例

```bash
# 修改 example/main.go 中的数据库配置
go run example/main.go
```

然后访问：
- API文档: http://localhost:8080/api/v1/examples
- 健康检查: http://localhost:8080/health
- 数据库管理: http://localhost:8081 (如果使用Docker Compose)

### 方式二：在你的项目中使用

```go
package main

import (
    "fmt"
    "log"
    
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    
    "github.com/JrMarcco/go-learning/gorm/models"
    "github.com/JrMarcco/go-learning/gorm/services"
    "github.com/JrMarcco/go-learning/gorm/types"
)

func main() {
    // 1. 连接数据库
    dsn := "user:password@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("连接数据库失败:", err)
    }
    
    // 2. 自动迁移
    db.AutoMigrate(&models.User{}, &models.UserProfile{}, &models.Role{})
    
    // 3. 创建服务
    userService := services.NewUserService(db, nil)
    
    // 4. 使用分页搜索
    params := &types.UserSearchParams{
        PaginationParams: types.PaginationParams{
            Page:     1,
            PageSize: 10,
            OrderBy:  "created_at",
            Sort:     "desc",
        },
        Name: "张", // 搜索姓名包含"张"的用户
    }
    
    result, err := userService.SearchUsers(params)
    if err != nil {
        log.Fatal("搜索失败:", err)
    }
    
    fmt.Printf("总数: %d, 当前页: %d\n", result.Total, result.Page)
}
```

## 4. 常用API示例

### 基础搜索

```bash
# 搜索姓名包含"张"的用户
curl "http://localhost:8080/api/v1/users/search?name=张&page=1&page_size=10"

# 按城市和状态搜索
curl "http://localhost:8080/api/v1/users/search?city=北京&status=1&order_by=age&sort=asc"
```

### 高级搜索

```bash
curl -X POST "http://localhost:8080/api/v1/users/search/advanced" \
  -H "Content-Type: application/json" \
  -d '{
    "name_like": "张",
    "status_in": [1, 2],
    "age_range": {"min": 20, "max": 30}
  }'
```

### 创建用户

```bash
curl -X POST "http://localhost:8080/api/v1/users" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "测试用户",
    "email": "test@example.com",
    "phone": "13800138000",
    "status": 1,
    "age": 25,
    "city": "北京"
  }'
```

## 5. 关键特性

### ✅ 安全性
- SQL注入防护
- 参数验证
- 字段白名单

### ✅ 性能优化
- 数据库索引
- 查询缓存
- 连接池

### ✅ 灵活搜索
- 模糊搜索
- 范围查询
- 复合条件
- 游标分页

### ✅ 开发友好
- 类型安全
- 错误处理
- 结构化响应

## 6. 目录结构说明

```
gorm/
├── models/          # 数据模型
├── types/           # 类型定义
├── utils/           # 工具函数
├── cache/           # 缓存层
├── services/        # 业务逻辑
├── handlers/        # HTTP处理器
├── example/         # 示例代码
├── docker-compose.yml  # Docker配置
├── init.sql         # 数据库初始化
└── README.md        # 详细文档
```

## 7. 故障排查

### 数据库连接失败
```bash
# 检查MySQL是否启动
docker-compose ps

# 查看MySQL日志
docker-compose logs mysql
```

### 依赖问题
```bash
# 清理并重新下载依赖
go mod tidy
go mod download
```

### 权限问题
```bash
# 确保数据库用户有足够权限
GRANT ALL PRIVILEGES ON gorm_example.* TO 'gorm_user'@'%';
FLUSH PRIVILEGES;
```

## 8. 下一步

- 查看 [README.md](README.md) 了解详细文档
- 阅读源码了解实现原理
- 根据需求自定义搜索条件
- 添加认证和权限控制
- 实现搜索高亮和建议功能

## 联系我们

如有问题，请提交Issue或Pull Request。 