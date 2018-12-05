package utils

import (
	"math/rand"
	"strings"
	"sync"
	"time"
)

var (
	sd  int64
	mtx sync.Mutex
)

// RandString 随机字串
// 根据内置字典随机生成字符串
// length 用来指定生成的随机串的长度
func RandString(length ...int) string {
	dict := strings.Split("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789", "")
	dtLen := len(dict)

	limit := 32
	if len(length) > 0 && length[0] > 0 {
		limit = length[0]
	}

	r := rand.New(rand.NewSource(randSeed()))
	res := ""
	for i := 0; i < limit; i++ {
		res = res + dict[r.Intn(dtLen)]
	}

	return res
}

func randSeed() int64 {
	mtx.Lock()
	defer mtx.Unlock()

	if sd >= 100000000 {
		sd = 1
	}

	sd++
	return time.Now().UnixNano() + sd
}
