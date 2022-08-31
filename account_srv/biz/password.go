package biz

import (
	"crypto/md5"
	"encoding/hex"
	"io"
)

func GetMd5(s string) string {
	hash := md5.New()
	_, _ = io.WriteString(hash, s)
	return hex.EncodeToString(hash.Sum(nil))
}
