package http

import (
	"io/ioutil"
	"testing"
)

func TestICanGetReaderFromInterfaceEntity(t *testing.T) {
	e := struct {
		Test string `json:"test"`
	}{Test: "wesh alors"}

	b := NewJSONEntityBody(e)
	trial, err := ioutil.ReadAll(b.GetBody())
	goal := "{\"test\":\"wesh alors\"}"

	if err != nil || string(trial) != goal {
		t.Error(err)
	}
}
