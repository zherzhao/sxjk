package mem

import (
	"sync"
	"time"
	"webconsole/global"
	"webconsole/pkg/cache/ICache"
)

func NewCache() *inMemoryCache {
	c := &inMemoryCache{
		make(map[string]value),
		sync.RWMutex{},
		ICache.Stat{},
		time.Duration(global.CacheSetting.TTL) * time.Second,
	}
	if global.CacheSetting.TTL > 0 {
		// 开启一个groutine后台处理缓存TTl
		go c.expirer()
	}
	return c
}
