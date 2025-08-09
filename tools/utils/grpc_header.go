package utils

import (
	"context"
	"strings"

	"github.com/google/uuid"
)

type mdIncomingKey struct{}

const (
	// RequestIDKey requestID key
	RequestIDKey = "x-request-id"
	// UsernameKey username key
	UsernameKey = "x-username"
)

// GetRequestID request id from header
func GetRequestID(ctx context.Context) string {
	id := GetHeaderFirst(ctx, RequestIDKey)
	if id == "" {
		id = NewRequestID()
	}
	return id
}

// GetUsername get username from header
func GetUsername(ctx context.Context) string {
	return GetHeaderFirst(ctx, UsernameKey)
}

// GetHeaderFirst get header first value
func GetHeaderFirst(ctx context.Context, key string) string {
	if md, ok := FromIncomingContext(ctx); ok {
		if values := md.Get(key); len(values) > 0 {
			return values[0]
		}
	}
	return ""
}

// FromIncomingContext returns the incoming metadata in ctx if it exists.  The
// returned MD should not be modified. Writing to it may cause races.
// Modification should be made to copies of the returned MD.
func FromIncomingContext(ctx context.Context) (md MD, ok bool) {
	md, ok = ctx.Value(mdIncomingKey{}).(MD)
	return
}

// NewRequestID generate a RequestId
func NewRequestID() string {
	return uuid.New().String()
}

// MD is a mapping from metadata keys to values. Users should use the following
// two convenience functions New and Pairs to generate MD.
type MD map[string][]string

// Get obtains the values for a given key.
func (md MD) Get(k string) []string {
	k = strings.ToLower(k)
	return md[k]
}
