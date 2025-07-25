package config

import (
	"github.com/Breeze0806/go-admin-core/storage"
	"github.com/Breeze0806/go-admin-core/storage/cache"
)

type Cache struct {
	Memory interface{}
}

// CacheConfig cache配置
var CacheConfig = new(Cache)

// Setup 构造cache 顺序 redis > 其他 > memory
func (e Cache) Setup() (storage.AdapterCache, error) {
	return cache.NewMemory(), nil
}
