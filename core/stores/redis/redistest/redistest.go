package redistest

import (
	"time"

	"github.com/MockyBang/go-zero/core/lang"
	"github.com/MockyBang/go-zero/core/stores/redis"
	"github.com/alicebob/miniredis/v2"
)

// CreateRedis returns a in process redis.Redis.
func CreateRedis() (r *redis.Redis, clean func(), err error) {
	mr, err := miniredis.Run()
	if err != nil {
		return nil, nil, err
	}

	return redis.New(mr.Addr()), func() {
		ch := make(chan lang.PlaceholderType)
		go func() {
			mr.Close()
			close(ch)
		}()
		select {
		case <-ch:
		case <-time.After(time.Second):
		}
	}, nil
}
