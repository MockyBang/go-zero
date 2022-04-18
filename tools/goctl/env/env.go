package env

import (
	"fmt"

	"github.com/MockyBang/go-zero/tools/goctl/pkg/env"
	"github.com/urfave/cli"
)

func Action(c *cli.Context) error {
	write := c.StringSlice("write")
	if len(write) > 0 {
		return env.WriteEnv(write)
	}
	fmt.Println(env.Print())
	return nil
}
