package context

import (
	"time"

	"github.com/urfave/cli"
)

// FromCli provides a context.Context from a github.com/urfave/cli.Context
//
// Most of this context is a no-op except for the .Value() calls.
func FromCli(c *cli.Context) Context {
	return &cliContext{t: time.Now(), c: c}
}

type cliContext struct {
	c *cli.Context
	t time.Time
}

// Type of the value to be gotten from the cli.Context
type Type int

// the types possible to be gotten from a cli.Context
const (
	TypeBool Type = iota
	TypeString
	TypeStringSlice
	TypeInt
	TypeIntSlice
	TypeInt64
	TypeInt64Slice
	TypeUint
	TypeUint64
	TypeArgs // this is a []string of the args from the cli
)

// V is a wrapper for github.com/urfave/cli.Context value getters, to indicate the Type of the value to be gotten from the cli.Context
func V(t Type, k string) interface{} {
	return cliValue{t: t, k: k}
}

type cliValue struct {
	t Type
	k string
}

// Value returns the value from the cli.Context, and expects V() used to request the value, otherwise defaults to passing k as a string to cli.Context.Generic(k)
func (c cliContext) Value(k interface{}) interface{} {
	ck, ok := (k).(cliValue)
	if !ok {
		if s, ok := (k).(string); ok {
			return c.c.Generic(s)
		} else {
			return nil
		}
	}
	switch ck.t {
	case TypeBool:
		return c.c.Bool(ck.k)
	case TypeString:
		return c.c.String(ck.k)
	case TypeStringSlice:
		return c.c.StringSlice(ck.k)
	case TypeInt:
		return c.c.Int(ck.k)
	case TypeIntSlice:
		return c.c.IntSlice(ck.k)
	case TypeInt64:
		return c.c.Int64(ck.k)
	case TypeInt64Slice:
		return c.c.Int64Slice(ck.k)
	case TypeUint:
		return c.c.Uint(ck.k)
	case TypeUint64:
		return c.c.Uint64(ck.k)
	case TypeArgs:
		return []string(c.c.Args())
	}
	return nil
}

func (c cliContext) Deadline() (deadline time.Time, ok bool) {
	return c.t, true
}

func (c cliContext) Done() <-chan struct{} {
	return nil
}

func (c cliContext) Err() error {
	return nil
}
