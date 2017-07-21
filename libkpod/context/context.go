package context

import "time"

// Context satisfies the "golang.org/x/net/context.Context" (or "context.Context" in go1.8+)
type Context interface {
	Deadline() (deadline time.Time, ok bool)
	Done() <-chan struct{}
	Err() error
	Value(key interface{}) interface{}
}
