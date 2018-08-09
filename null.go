package locker

import (
	"context"

	config "github.com/dynamicgo/go-config"
)

type nullLocker struct {
}

func (locker *nullLocker) Lock(ctx context.Context) error {
	return nil
}
func (locker *nullLocker) Unlock(ctx context.Context) error {
	return nil
}

func init() {
	Register("null", func(config config.Config) (Locker, error) {
		return &nullLocker{}, nil
	})
}
