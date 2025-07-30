package cache

import (
	"encoding/json"
	"time"

	"github.com/go-redis/redis"
)

// Cache 缓存接口
type Cache interface {
	Get(key string) ([]byte, error)
	Set(key string, value []byte, expiration time.Duration) error
	Delete(key string) error
	Exists(key string) bool
	FlushPattern(pattern string) error
}

// RedisCache Redis缓存实现
type RedisCache struct {
	client *redis.Client
}

// NewRedisCache 创建Redis缓存实例
func NewRedisCache(addr, password string, db int) *RedisCache {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	return &RedisCache{
		client: rdb,
	}
}

// Get 获取缓存值
func (r *RedisCache) Get(key string) ([]byte, error) {
	val, err := r.client.Get(key).Result()
	if err != nil {
		return nil, err
	}
	return []byte(val), nil
}

// Set 设置缓存值
func (r *RedisCache) Set(key string, value []byte, expiration time.Duration) error {
	return r.client.Set(key, value, expiration).Err()
}

// Delete 删除缓存
func (r *RedisCache) Delete(key string) error {
	return r.client.Del(key).Err()
}

// Exists 检查键是否存在
func (r *RedisCache) Exists(key string) bool {
	count, err := r.client.Exists(key).Result()
	if err != nil {
		return false
	}
	return count > 0
}

// FlushPattern 删除匹配模式的键
func (r *RedisCache) FlushPattern(pattern string) error {
	keys, err := r.client.Keys(pattern).Result()
	if err != nil {
		return err
	}

	if len(keys) > 0 {
		return r.client.Del(keys...).Err()
	}

	return nil
}

// MemoryCache 内存缓存实现（简单实现，仅用于测试）
type MemoryCache struct {
	data map[string]cacheItem
}

type cacheItem struct {
	value      []byte
	expiration time.Time
}

// NewMemoryCache 创建内存缓存实例
func NewMemoryCache() *MemoryCache {
	return &MemoryCache{
		data: make(map[string]cacheItem),
	}
}

// Get 获取缓存值
func (m *MemoryCache) Get(key string) ([]byte, error) {
	item, exists := m.data[key]
	if !exists {
		return nil, redis.Nil
	}

	if time.Now().After(item.expiration) {
		delete(m.data, key)
		return nil, redis.Nil
	}

	return item.value, nil
}

// Set 设置缓存值
func (m *MemoryCache) Set(key string, value []byte, expiration time.Duration) error {
	m.data[key] = cacheItem{
		value:      value,
		expiration: time.Now().Add(expiration),
	}
	return nil
}

// Delete 删除缓存
func (m *MemoryCache) Delete(key string) error {
	delete(m.data, key)
	return nil
}

// Exists 检查键是否存在
func (m *MemoryCache) Exists(key string) bool {
	item, exists := m.data[key]
	if !exists {
		return false
	}

	if time.Now().After(item.expiration) {
		delete(m.data, key)
		return false
	}

	return true
}

// FlushPattern 删除匹配模式的键
func (m *MemoryCache) FlushPattern(pattern string) error {
	// 简单实现，仅清空所有键
	m.data = make(map[string]cacheItem)
	return nil
}

// CacheManager 缓存管理器
type CacheManager struct {
	cache Cache
}

// NewCacheManager 创建缓存管理器
func NewCacheManager(cache Cache) *CacheManager {
	return &CacheManager{
		cache: cache,
	}
}

// GetJSON 获取JSON格式的缓存
func (cm *CacheManager) GetJSON(key string, dest interface{}) error {
	data, err := cm.cache.Get(key)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, dest)
}

// SetJSON 设置JSON格式的缓存
func (cm *CacheManager) SetJSON(key string, value interface{}, expiration time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return cm.cache.Set(key, data, expiration)
}

// GetCache 获取原始缓存接口
func (cm *CacheManager) GetCache() Cache {
	return cm.cache
}
