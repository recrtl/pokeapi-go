package pokeapi

import (
	"testing"
	"time"

	"github.com/recrtl/pokeapi-go/structs"
	"github.com/stretchr/testify/assert"
)

var endpoint = "machine?offset=0&limit=1"
var mockResource structs.Resource

func TestSetCache(t *testing.T) {
	_, found1 := c.Get(endpoint)
	assert.Equal(t, false, found1,
		"Expect to not have cached data before first call.")

	var resource1 structs.Resource
	_ = do(endpoint, &resource1)
	_, expires2, found2 := c.GetWithExpiration(endpoint)
	assert.Equal(t, true, found2,
		"Expect to have cached data after first call.")

	var resource2 structs.Resource
	_ = do(endpoint, &resource2)
	_, expires3, _ := c.GetWithExpiration(endpoint)
	assert.Equal(t, resource1, resource2,
		"Expect data to match cached data.")
	assert.Equal(t, expires2, expires3,
		"Expect expiration times to match.")
}

func TestClearCache(t *testing.T) {
	_, found := c.Get(endpoint)
	if !found {
		_ = do(endpoint, &mockResource)
		_, found = c.Get(endpoint)
	}
	assert.NotEqual(t, false, found,
		"Expect to start with cached data.")

	ClearCache()
	nocache, nofound := c.Get(endpoint)
	assert.Equal(t, false, nofound,
		"Expect no data found after flushing cache.")
	assert.Equal(t, nil, nocache,
		"Expect no data after flushing cache.")
}

func TestCustomExpiration(t *testing.T) {
	ClearCache()
	defaultExpire := time.Now().Add(defaultCacheSettings.MinExpire).Minute()
	_ = do(endpoint, &mockResource)
	_, expires1, _ := c.GetWithExpiration(endpoint)
	assert.Equal(t, defaultExpire, expires1.Minute(),
		"Expect expiration time to match default setting.")

	ClearCache()
	CacheSettings.CustomExpire = 20 * time.Minute
	customExpire := time.Now().Add(CacheSettings.CustomExpire).Minute()
	_ = do(endpoint, &mockResource)
	_, expires2, _ := c.GetWithExpiration(endpoint)
	assert.Equal(t, customExpire, expires2.Minute(),
		"Expect expiration time to match custom setting.")
}

func TestNoCache(t *testing.T) {
	ClearCache()
	_ = do(endpoint, &mockResource)
	_, expires1, found1 := c.GetWithExpiration(endpoint)
	assert.Equal(t, true, found1,
		"Expect to have cached data after first call.")

	CacheSettings.UseCache = false
	_ = do(endpoint, &mockResource)
	_, expires2, _ := c.GetWithExpiration(endpoint)
	assert.NotEqual(t, expires1, expires2,
		"Expect cache expiration not to match first call.")
}
