package cache

import (
	"encoding/json"
	"io/ioutil"
	"time"
)

type Cache struct {
	Creation   int64       `json:"creation"`
	Expiration int64       `json:"expiration"`
	Data       interface{} `json:"data"`
	Rotting    int64       `json:"rotting"`
}

func (c *Cache) IsExpired() bool {
	return time.Now().Unix() >= c.Expiration
}

func (c *Cache) IsValid() bool {
	return !c.IsExpired()
}

func (c *Cache) TimeRemaining() int64 {
	return time.Now().Unix() - c.Expiration
}

func Parse(path string, data interface{}) (*Cache, error) {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	c := &Cache{Data: data}

	if err := json.Unmarshal(f, c); err != nil {
		return nil, err
	}

	return c, nil
}
