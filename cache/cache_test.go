package cache

import (
	"testing"
	"time"
)

type dummyData struct {
	Test string `json:"test"`
}

func TestICanParseCacheFile(t *testing.T) {
	data := &dummyData{}
	goal := "wesh alors"

	c, err := Parse("./testdata/1-simple-cache.json", data)
	if err != nil || c == nil {
		t.Fatal(err)
	}

	trial := c.Data.(*dummyData)

	if trial.Test != goal {
		t.Fail()
	}
}

func TestIFailOnUnreachableFile(t *testing.T) {
	c, err := Parse("zbrah.json", nil)
	if err == nil || c != nil {
		t.Fatal("Should have been an error here")
	}
}

func TestIKnowIfCacheIsExpired(t *testing.T) {
	goal := int64(3600)
	c1 := &Cache{
		Creation:   time.Now().Unix(),
		Expiration: time.Now().Unix() + goal,
	}

	if c1.IsExpired() == true || c1.IsValid() == false {
		t.Fatal("Should be expired")
	}

	if c1.TimeRemaining() != -goal {
		t.Fatal("Should be equal to -goal")
	}
}
