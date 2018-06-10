package tools

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
)

type SHA1Bytes []byte
type SHA1 string

func (s SHA1) Bytes() []byte {
	b := SHA1Bytes(s)
	return b.Bytes()
}

func (s SHA1) String() string {
	b := SHA1Bytes(s)
	return b.String()
}

func (s SHA1Bytes) Bytes() []byte {
	return []byte(s.String())
}

func (s SHA1Bytes) String() string {
	h := sha1.New()
	h.Write(s)
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

type MD5 string

func (m MD5) String() string {
	h := md5.New()
	h.Write([]byte(m))
	return hex.EncodeToString(h.Sum(nil))
}
