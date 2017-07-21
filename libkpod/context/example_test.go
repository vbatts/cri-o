package context_test

import (
	kc "github.com/kubernetes-incubator/cri-o/libkpod/context"
	"github.com/urfave/cli"
	"golang.org/x/net/context"
)

func doStuff(ctx context.Context, arg1 int) {
	if v, ok := ctx.Value(kc.V(kc.TypeInt, "flag-with-numbers")).(int); ok {
		if v == arg1 {
			println("yay. commpared a value from a cli flag")
		}
	}
}

func myCliSubCommand(c *cli.Context) {
	doStuff(kc.FromCli(c), 0)
}
