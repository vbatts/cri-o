package debug

import (
	"runtime"

	kc "github.com/kubernetes-incubator/cri-o/libkpod/context"
	"github.com/kubernetes-incubator/cri-o/libkpod/info"
	"golang.org/x/net/context"
)

func init() {
	info.RegisterInfoGiver(debugInfo)
}

func debugInfo(ctx context.Context) (string, map[string]interface{}, error) {
	// only if the --debug flag is passed
	if ctx.Value(kc.V(kc.TypeInt, "debug")).(bool) {
		return "debug", nil, info.ErrNoInfo
	}

	i := map[string]interface{}{}
	i["compiler"] = runtime.Compiler
	i["go version"] = runtime.Version()

	return "debug", i, nil
}
