package cache

import (
	"log"
	"webconsole/pkg/cache/ICache"
	"webconsole/pkg/cache/mem"
	"webconsole/pkg/cache/rdb"
)

func New(typ string, ttl int) ICache.Cache {
	var c ICache.Cache

	if typ == "mem" {
		c = mem.NewCache()
	} else if typ == "disk" {
		c = rdb.NewCache()
	}

	if c == nil {
		panic("未指定类型")
	}

	log.Println(typ, "服务已就位")
	return c
}
