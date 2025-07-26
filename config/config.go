// Package config is an interface for dynamic configuration.
package config

import (
	"context"

	"github.com/Breeze0806/go-admin-core/config/source"
)

// Config is an interface abstraction for dynamic configuration
type Config interface {
	// Init the config
	Init(opts ...Option) error
	// Options in the config
	Options() Options
}

// Entity 配置实体
type Entity interface {
	OnChange()
}

// Options 配置的参数
type Options struct {
	Source []source.Source

	// for alternative data
	Context context.Context

	Entity Entity
}

// Option 调用类型
type Option func(o *Options)

var (
	// DefaultConfig Default Config Manager
	DefaultConfig Config
)

// NewConfig returns new config
func NewConfig(opts ...Option) (Config, error) {
	return newConfig(opts...)
}
