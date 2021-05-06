package mem

import (
	"sync"
	"time"
	"webconsole/global"
	"webconsole/pkg/cache/ICache"
)

type value struct {
	c       []byte
	created time.Time
}

type inMemoryCache struct {
	c           map[string]value //缓存键值对
	mutex       sync.RWMutex     //读写一致性控制
	ICache.Stat                  //缓存当前状态
	ttl         time.Duration    //缓存生存时间
}

func (c *inMemoryCache) expirer() {
	for {
		time.Sleep(c.ttl)
		c.mutex.Lock()

		for k, v := range c.c {
			c.mutex.Unlock()
			if v.created.Add(c.ttl).Before(time.Now()) {
				c.Del(k)
			}
			c.mutex.Lock()
		}

		c.mutex.Unlock()
	}
}

func (c *inMemoryCache) Set(k string, v []byte) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	tmp, exist := c.c[k]
	if exist {
		c.Remove(k, tmp.c)
	}
	c.c[k] = value{v, time.Now()}
	c.Add(k, v)
	return nil
}

func (c *inMemoryCache) Get(k string) ([]byte, error) {
	if global.CacheSetting.TTL != 0 {
		c.mutex.Lock()
		val := c.c[k].c
		c.c[k] = value{val, time.Now()}

		defer c.mutex.Unlock()
		return val, nil
	} else {
		c.mutex.RLock()
		defer c.mutex.RUnlock()
		return c.c[k].c, nil
	}
}

func (c *inMemoryCache) Del(k string) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	v, exist := c.c[k]
	if exist {
		delete(c.c, k)
		c.Remove(k, v.c)
	}
	return nil
}

func (c *inMemoryCache) GetStat() ICache.Stat {
	return c.Stat
}
