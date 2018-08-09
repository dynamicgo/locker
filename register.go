package locker

import (
	"fmt"
	"sync"

	"github.com/dynamicgo/slf4go"

	config "github.com/dynamicgo/go-config"
)

type lockerRegister struct {
	sync.RWMutex
	slf4go.Logger
	drivers map[string]Driver
}

// Driver locker factory
type Driver func(config config.Config) (Locker, error)

var global *lockerRegister
var globalOnce sync.Once

func initRegister() {
	global = &lockerRegister{
		Logger:  slf4go.Get("locker-register"),
		drivers: make(map[string]Driver),
	}
}

// Register .
func Register(name string, driver Driver) {
	globalOnce.Do(initRegister)

	global.Lock()
	defer global.Unlock()

	_, ok := global.drivers[name]

	if ok {
		panic(fmt.Errorf("duplicate import drvier name %s", name))
	}

	global.drivers[name] = driver

	global.DebugF("register locker driver %s", name)
}

func getDriver(name string) Driver {
	globalOnce.Do(initRegister)

	global.RLock()
	defer global.RUnlock()

	driver, ok := global.drivers[name]

	if !ok {
		panic(fmt.Errorf("invalid locker driver %s", name))
	}

	return driver
}
