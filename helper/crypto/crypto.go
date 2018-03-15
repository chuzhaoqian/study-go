package crypto

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"time"
)

func Md5(str string) string {
	result := md5.Sum([]byte(str))
	return fmt.Sprintf("%x", result)
}

func Rand() *rand.Rand {
	source := rand.NewSource(time.Now().UnixNano())
	return rand.New(source)
}

func RandInt64() int64 {
	r := Rand()
	return r.Int63()
}

func RandInt64Str() string {
	return fmt.Sprintf("%d", RandInt64())
}
