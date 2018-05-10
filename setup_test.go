package tools

import (
	"errors"
	"io"
	"testing"
)

func assertEqual(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Fatalf("%s != %s", a, b)
	}
}

func assertDiff(t *testing.T, a interface{}, b interface{}) {
	if a == b {
		t.Fatalf("%s != %s", a, b)
	}
}

type dummyReadCloser struct {
	fr     io.Reader
	isRead bool
}

func (d *dummyReadCloser) Close() error {
	if d.isRead == false {
		return errors.New("Not over")
	}
	return nil
}

func (d *dummyReadCloser) Read(p []byte) (n int, err error) {
	d.isRead = true
	return d.fr.Read(p)
}
