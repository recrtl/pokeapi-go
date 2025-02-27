package pokeapi

import (
	"encoding/gob"
	"time"
)

var CacheSettings = Settings{
	CustomExpire: 0,
	UseCache:     true,
}

var defaultCacheSettings = defaultSettings{
	MaxExpire: 10 * time.Minute,
	MinExpire: 5 * time.Minute,
}

// defaultSettings are settings not meant to be changed.
type defaultSettings struct {
	MaxExpire time.Duration
	MinExpire time.Duration
}

// CacheSettings are user settings for cache expiration.
type Settings struct {
	CustomExpire time.Duration
	UseCache     bool
}

// ClearCache clears all cached data.
func ClearCache() {
	c.Flush()
}

// setCache adds new item to local cache.
func setCache(endpoint string, obj interface{}) {
	if CacheSettings.CustomExpire != 0 {
		c.Set(endpoint, obj, CacheSettings.CustomExpire)
	} else {
		c.SetDefault(endpoint, obj)
	}
}

func init() {
	gob.Register(map[string]interface{}{})
}
