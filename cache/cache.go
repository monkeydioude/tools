package cache

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

type Cache interface {
	GetExpiration() int64
	GetData() interface{}
	Write(path string)
}

type FileCache struct {
	path       string
	Expiration int64       `json:"expiration"`
	Data       interface{} `json:"data"`
}

func (c *FileCache) IsExpired() bool {
	return time.Now().Unix() >= c.Expiration
}

func (c *FileCache) IsValid() bool {
	return !c.IsExpired()
}

func (c *FileCache) TimeRemaining() int64 {
	return time.Now().Unix() - c.Expiration
}

func ParseFileCache(path string, data interface{}) (*FileCache, error) {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	c := &FileCache{Data: data}

	if err := json.Unmarshal(f, c); err != nil {
		return nil, err
	}

	return c, nil
}

func NewFileCache(path string) *FileCache {
	return &FileCache{
		path: path,
	}
}

func (c *FileCache) Write(data interface{}, expiration int64) error {
	c.Data = data
	c.Expiration = expiration

	cache, err := json.Marshal(c)
	if err != nil {
		fmt.Printf("[ERR ] Could not marshal data. Reason: %s\n", err)
		return err
	}

	return ioutil.WriteFile(c.path, cache, 0666)
}

func (c *FileCache) GetData() interface{} {
	return c.Data
}
