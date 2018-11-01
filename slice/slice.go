package slice

import (
	"math/rand"
	"time"
)

//打乱切片顺序
func Shuffle(sliceData []interface{}) []interface{} {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	sliceDataLen := len(sliceData)
	result := make([]interface{}, sliceDataLen)
	for k, v := range r.Perm(sliceDataLen) {
		result[k] = sliceData[v]
	}
	return result
}
