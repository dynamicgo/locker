package locker

import (
	"context"
	"sync"

	config "github.com/dynamicgo/go-config"
)

type nullLocker struct {
	sync.Mutex
}

func (locker *nullLocker) Lock(ctx context.Context) error {
	locker.Mutex.Lock()
	return nil
}
func (locker *nullLocker) Unlock(ctx context.Context) error {
	locker.Mutex.Unlock()
	return nil
}

func init() {
	Register("null", func(config config.Config) (Locker, error) {
		return &nullLocker{}, nil
	})
}
