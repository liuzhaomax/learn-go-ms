package biz

import (
	"crypto/md5"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
	"testing"
)

func TestGetMd5(t *testing.T) {
	s1 := "happy"
	fmt.Println(GetMd5(s1))
	options := password.Options{
		SaltLen:      16,
		Iterations:   100,
		KeyLen:       32,
		HashFunction: md5.New,
	}
	salt, encodedPwd := password.Encode("happy", &options)
	fmt.Println(salt)
	fmt.Println(encodedPwd)
	check := password.Verify("happy", salt, encodedPwd, &options)
	fmt.Println(check)
}
