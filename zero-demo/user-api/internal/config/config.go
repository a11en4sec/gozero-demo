package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	DB struct {
		DataSource string
	}
	// Cache struct {
	// 	Host string
	// 	Pass string
	// }
	Cache cache.CacheConf
}
