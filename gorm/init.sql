-- 创建数据库
CREATE DATABASE IF NOT EXISTS gorm_example CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE gorm_example;

-- 创建用户表（GORM会自动创建，这里仅作为参考）
-- CREATE TABLE users (
--     id bigint unsigned NOT NULL AUTO_INCREMENT,
--     name varchar(100) NOT NULL,
--     email varchar(255) NOT NULL,
--     phone varchar(20) DEFAULT NULL,
--     status int NOT NULL DEFAULT '1' COMMENT '状态 1:active 2:inactive',
--     age int DEFAULT NULL,
--     city varchar(50) DEFAULT NULL,
--     created_at datetime(3) DEFAULT NULL,
--     updated_at datetime(3) DEFAULT NULL,
--     deleted_at datetime(3) DEFAULT NULL,
--     PRIMARY KEY (id),
--     UNIQUE KEY email (email),
--     KEY idx_user_search (name),
--     KEY phone (phone),
--     KEY status (status),
--     KEY age (age),
--     KEY city (city),
--     KEY created_at (created_at),
--     KEY deleted_at (deleted_at)
-- );

-- 性能优化索引建议（在GORM自动迁移后执行）
-- CREATE INDEX idx_users_status_city ON users(status, city);
-- CREATE INDEX idx_users_age_status ON users(age, status);
-- CREATE INDEX idx_users_created_status ON users(created_at, status);

-- 示例数据（可选）
-- INSERT INTO users (name, email, phone, status, age, city, created_at, updated_at) VALUES
-- ('张三', 'zhangsan@example.com', '13800138001', 1, 25, '北京', NOW(), NOW()),
-- ('李四', 'lisi@example.com', '13800138002', 1, 30, '上海', NOW(), NOW()),
-- ('王五', 'wangwu@example.com', '13800138003', 2, 28, '广州', NOW(), NOW()),
-- ('赵六', 'zhaoliu@example.com', '13800138004', 1, 35, '深圳', NOW(), NOW()),
-- ('孙七', 'sunqi@example.com', '13800138005', 1, 22, '北京', NOW(), NOW()); 