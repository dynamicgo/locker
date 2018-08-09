package locker

import (
	"context"
	"time"

	config "github.com/dynamicgo/go-config"
)

// Locker .
type Locker interface {
	Lock(ctx context.Context) error
	Unlock(ctx context.Context) error
}

// New Create new distributed locker with driver name and config
func New(driver string, config config.Config) (Locker, error) {
	d := getDriver(driver)

	return d(config)
}

// TryLock lock .
func TryLock(locker Locker, timeout time.Duration, f func() error) (bool, error) {
	return TryLockWithContext(context.Background(), locker, timeout, f)
}

// TryLockWithContext .
func TryLockWithContext(ctx context.Context, locker Locker, timeout time.Duration, f func() error) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	if err := locker.Lock(ctx); err != nil {
		if err == context.DeadlineExceeded {
			return false, nil
		}

		return false, err
	}

	err := f()

	unlockError := locker.Unlock(context.Background())

	if unlockError != nil {
		println("warning !!!!! unlock error ", unlockError.Error())
	}

	return true, err
}

// LockWithContext .
func LockWithContext(ctx context.Context, locker Locker, f func() error) error {
	if err := locker.Lock(ctx); err != nil {

		return err
	}

	err := f()

	unlockError := locker.Unlock(context.Background())

	if unlockError != nil {
		println("warning !!!!! unlock error ", unlockError.Error())
	}

	return err
}

// Lock .
func Lock(locker Locker, f func() error) error {
	return LockWithContext(context.Background(), locker, f)
}
