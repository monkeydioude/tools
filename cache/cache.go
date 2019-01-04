package cache

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type Cache interface {
	GetExpiration() int64
	GetData() interface{}
	Write(interface{}, int64) error
	Parse() (bool, error)
}

type FileCache struct {
	path       string
	Expiration int64       `json:"expiration"`
	Data       interface{} `json:"data"`
}

func NewFileCache(path string, data interface{}) *FileCache {
	return &FileCache{
		path: path,
		Data: data,
	}
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

func (c *FileCache) Parse() (bool, error) {
	_, err := os.Stat(c.path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}

	f, err := ioutil.ReadFile(c.path)
	if err != nil {
		return false, err
	}

	if err := json.Unmarshal(f, c); err != nil {
		return false, err
	}

	return true, nil
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

func (c *FileCache) GetExpiration() int64 {
	return c.Expiration
}

func ParseFileCache(path string, data interface{}) (*FileCache, error) {
	c := &FileCache{Data: data}
	ex, err := c.Parse()

	if err != nil {
		return nil, err
	}

	if ex == false {
		return nil, errors.New("Could not parse file cache, file does not exist")
	}

	return c, nil
}
