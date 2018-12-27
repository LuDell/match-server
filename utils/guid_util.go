package utils

import (
	"io"
	"crypto/rand"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
)

func UniqueId() string {
	by :=make([]byte,48)
	if _,err := io.ReadFull(rand.Reader,by); err != nil {
		return ""
	}
	h := md5.New()
	h.Write([]byte(base64.URLEncoding.EncodeToString(by)))
	return hex.EncodeToString(h.Sum(nil))
}
