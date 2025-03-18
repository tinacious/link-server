package utils

import (
	"math/rand"
	"time"
)

func GetRandomPort() int {
	var min int = 1024
	var max int = 66535

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	return r.Intn(max-min+1) + min
}
