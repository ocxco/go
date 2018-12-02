package tool

import (
	"math/rand"
	"time"
)

func MakeArr(num int) []int {
	rd := rand.New(rand.NewSource(time.Now().UnixNano()))
	arr := make([]int, num)
	for i := 0; i < num; i++ {
		arr[i] = rd.Intn(1000)
	}
	return arr
}