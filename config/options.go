package config

import (
	"github.com/Breeze0806/go-admin-core/config/source"
)

// WithSource appends a source to list of sources
func WithSource(s source.Source) Option {
	return func(o *Options) {
		o.Source = append(o.Source, s)
	}
}
