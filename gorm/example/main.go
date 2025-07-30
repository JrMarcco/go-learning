package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/JrMarcco/go-learning/gorm/cache"
	"github.com/JrMarcco/go-learning/gorm/handlers"
	"github.com/JrMarcco/go-learning/gorm/models"
	"github.com/JrMarcco/go-learning/gorm/services"
)

// DBConfig 数据库配置
type DBConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
	Charset  string `json:"charset"`
}

// RedisConfig Redis配置
type RedisConfig struct {
	Addr     string `json:"addr"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}

// initDatabase 初始化数据库连接
func initDatabase(config DBConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		config.Username, config.Password, config.Host, config.Port, config.Database, config.Charset)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // 开启SQL日志
	})

	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}

	// 设置连接池
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get sql.DB: %w", err)
	}

	sqlDB.SetMaxIdleConns(10)  // 最大空闲连接数
	sqlDB.SetMaxOpenConns(100) // 最大打开连接数

	return db, nil
}

// initCache 初始化缓存
func initCache(config RedisConfig) *cache.CacheManager {
	// 创建Redis缓存（生产环境）
	redisCache := cache.NewRedisCache(config.Addr, config.Password, config.DB)

	// 如果Redis不可用，可以降级使用内存缓存
	// memoryCache := cache.NewMemoryCache()

	return cache.NewCacheManager(redisCache)
}

// migrateDatabase 执行数据库迁移
func migrateDatabase(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.User{},
		&models.UserProfile{},
		&models.Role{},
	)
}

// seedDatabase 初始化测试数据
func seedDatabase(userService *services.UserService) error {
	// 创建测试用户
	users := []models.User{
		{
			Name:   "张三",
			Email:  "zhangsan@example.com",
			Phone:  "13800138001",
			Status: 1,
			Age:    25,
			City:   "北京",
			Profile: &models.UserProfile{
				Avatar: "https://example.com/avatar1.jpg",
				Bio:    "这是张三的个人简介",
				Gender: 1,
			},
		},
		{
			Name:   "李四",
			Email:  "lisi@example.com",
			Phone:  "13800138002",
			Status: 1,
			Age:    30,
			City:   "上海",
			Profile: &models.UserProfile{
				Avatar: "https://example.com/avatar2.jpg",
				Bio:    "这是李四的个人简介",
				Gender: 2,
			},
		},
		{
			Name:   "王五",
			Email:  "wangwu@example.com",
			Phone:  "13800138003",
			Status: 2,
			Age:    28,
			City:   "广州",
			Profile: &models.UserProfile{
				Avatar: "https://example.com/avatar3.jpg",
				Bio:    "这是王五的个人简介",
				Gender: 1,
			},
		},
	}

	// 批量创建用户
	return userService.BatchCreateUsers(users)
}

func main() {
	// 配置信息（实际项目中应该从配置文件或环境变量读取）
	dbConfig := DBConfig{
		Host:     "localhost",
		Port:     3306,
		Username: "root",
		Password: "password",
		Database: "gorm_example",
		Charset:  "utf8mb4",
	}

	redisConfig := RedisConfig{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	}

	// 初始化数据库
	db, err := initDatabase(dbConfig)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// 执行数据库迁移
	if err := migrateDatabase(db); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// 初始化缓存
	cacheManager := initCache(redisConfig)

	// 初始化服务
	userService := services.NewUserService(db, cacheManager)

	// 初始化测试数据（可选）
	if err := seedDatabase(userService); err != nil {
		log.Printf("Warning: Failed to seed database: %v", err)
	}

	// 初始化处理器
	userHandler := handlers.NewUserHandler(userService)

	// 初始化Gin路由
	gin.SetMode(gin.ReleaseMode) // 生产模式
	router := gin.Default()

	// 添加中间件
	router.Use(gin.Recovery())      // 恢复中间件
	router.Use(corsMiddleware())    // CORS中间件
	router.Use(loggingMiddleware()) // 日志中间件

	// 健康检查
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "GORM pagination example is running",
		})
	})

	// API路由组
	v1 := router.Group("/api/v1")

	// 注册用户路由
	userHandler.RegisterRoutes(v1)

	// 添加其他路由示例
	v1.GET("/examples", func(c *gin.Context) {
		examples := map[string]interface{}{
			"basic_search":        "/api/v1/users/search?name=张&page=1&page_size=10",
			"advanced_search":     "/api/v1/users/search?name=张&status=1&min_age=20&max_age=30&city=北京&order_by=created_at&sort=desc",
			"cursor_pagination":   "/api/v1/users/search/cursor?limit=10&direction=next",
			"simple_search":       "/api/v1/users/search/simple?name=李",
			"role_search":         "/api/v1/users/role/1",
			"stats":               "/api/v1/users/stats",
			"batch_create":        "POST /api/v1/users/batch",
			"advanced_conditions": "POST /api/v1/users/search/advanced",
		}

		c.JSON(http.StatusOK, gin.H{
			"message":   "GORM Pagination API Examples",
			"endpoints": examples,
		})
	})

	// 启动服务器
	port := ":8080"
	log.Printf("Server starting on port %s", port)
	log.Printf("API Documentation: http://localhost%s/api/v1/examples", port)
	log.Printf("Health Check: http://localhost%s/health", port)

	if err := router.Run(port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

// corsMiddleware CORS中间件
func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// loggingMiddleware 日志中间件
func loggingMiddleware() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format("2006/01/02 - 15:04:05"),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	})
}
