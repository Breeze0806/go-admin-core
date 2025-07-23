package config

import (
	"github.com/Breeze0806/go-admin-core/storage"
	"github.com/Breeze0806/go-admin-core/storage/queue"
	"time"
)

type Queue struct {
	Memory *QueueMemory
}


type QueueMemory struct {
	PoolSize uint
}

var QueueConfig = new(Queue)

// Empty 空设置
func (e Queue) Empty() bool {
	return e.Memory == nil && e.Redis == nil && e.NSQ == nil
}

// Setup  memory
func (e Queue) Setup() (storage.AdapterQueue, error) {
	return queue.NewMemory(e.Memory.PoolSize), nil
}
