package cache

import (
	"testing"
	"time"
)

type dummyData struct {
	uselessField string
	A            struct {
		B struct {
			NotTest string `json:"test"`
		} `json:"lvl2"`
	} `json:"lvl1"`
}

func TestICanParseComplexCacheFile(t *testing.T) {
	data := &dummyData{}
	goal := "kikoo"

	c, err := ParseFileCache("./testdata/2-simple-cache.json", data)
	if err != nil || c == nil {
		t.Fatal(err)
	}

	trial := c.Data.(*dummyData)

	if trial.A.B.NotTest != goal {
		t.Fail()
	}
}

func TestIFailOnUnreachableFile(t *testing.T) {
	c, err := ParseFileCache("zbrah.json", nil)
	if err == nil || c != nil {
		t.Fatal("Should have been an error here")
	}
}

func TestIKnowIfCacheIsExpired(t *testing.T) {
	goal := int64(3600)
	c1 := &FileCache{
		Expiration: time.Now().Unix() + goal,
	}

	if c1.IsExpired() == true || c1.IsValid() == false {
		t.Fatal("Should be expired")
	}

	if c1.TimeRemaining() != -goal {
		t.Fatal("Should be equal to -goal")
	}
}

type dummyData2 struct {
	Test string `json:"test"`
}

func TestICanWriteCache(t *testing.T) {
	p := "./testdata/test"
	f := NewFileCache(p)
	goal := dummyData2{
		Test: "wesh",
	}
	dd := &dummyData2{}

	f.Write(goal, time.Now().Add(15*time.Minute).Unix())
	cacheRead, err := ParseFileCache(p, dd)
	if err != nil {
		t.Fail()
	}

	trial := cacheRead.Data.(*dummyData2)

	if trial.Test != goal.Test {
		t.Fail()
	}
}
