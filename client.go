package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/patrickmn/go-cache"
)

const apiurl = "https://pokeapi.co/api/v2/"

var c *cache.Cache

func init() {
	c = cache.New(defaultCacheSettings.MinExpire, defaultCacheSettings.MaxExpire)
}

func do(endpoint string, obj interface{}) error {
	cached, found := c.Get(endpoint)
	if found && CacheSettings.UseCache {
		v := reflect.ValueOf(obj)
		if v.Kind() != reflect.Ptr {
			return errors.New("obj must be a pointer")
		}

		v.Elem().Set(reflect.ValueOf(cached))
		return nil
	}

	req, err := http.NewRequest(http.MethodGet, apiurl+endpoint, nil)
	if err != nil {
		return err
	}
	client := &http.Client{Timeout: 10 * time.Second}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, obj)
	if err != nil {
		return err
	}

	// trust it's a pointer, otherwise json.Unmarshal would fail
	objValue := reflect.ValueOf(obj).Elem().Interface()
	setCache(endpoint, objValue)
	return nil
}

// URL returns a result from a url. Useful to follow resources.
func URL[T any](url string) (result T, err error) {
	if !strings.HasPrefix(url, apiurl) {
		return result, fmt.Errorf("url %s is invalid", url)
	}

	url = strings.TrimPrefix(url, apiurl)
	err = do(url, &result)

	return result, err
}
