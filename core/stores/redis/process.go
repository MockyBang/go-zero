package redis

import (
	"strings"

	"github.com/MockyBang/go-zero/core/logx"
	"github.com/MockyBang/go-zero/core/mapping"
	"github.com/MockyBang/go-zero/core/timex"
	red "github.com/go-redis/redis"
)

func checkDuration(proc func(red.Cmder) error) func(red.Cmder) error {
	return func(cmd red.Cmder) error {
		start := timex.Now()

		defer func() {
			duration := timex.Since(start)
			if duration > slowThreshold.Load() {
				var buf strings.Builder
				for i, arg := range cmd.Args() {
					if i > 0 {
						buf.WriteByte(' ')
					}
					buf.WriteString(mapping.Repr(arg))
				}
				logx.WithDuration(duration).Slowf("[REDIS] slowcall on executing: %s", buf.String())
			}
		}()

		return proc(cmd)
	}
}
