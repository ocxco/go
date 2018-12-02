package main

import (
	"fmt"
	"sort/sort"
	"sort/tool"
)

func cpy(arr []int) []int {
	v := make([]int, len(arr));
	copy(v, arr)
	return v
}

func doBubbleSort(arr []int, debug bool) []int  {
	fmt.Printf("冒泡排序开始,总共%d个元素\n", len(arr))
	var v = cpy(arr)
	bs := sort.NewBS(debug)
	tr.Start()
	bs.Sort(v)
	tr.End()
	fmt.Printf("冒泡排序结束，耗时%s, 循环次数%d, 交换次数%d\n", tr.MicroSec(), bs.CycleCount, bs.SwapCount)
	return v
}

func doQuickSort(arr []int, debug bool) []int  {
	tr.Start();
	fmt.Printf("快速排序开始,总共%d个元素\n", len(arr))
	qs := sort.NewQS(debug)
	var v = cpy(arr)
	tr.End()
	qs.Sort(v)
	fmt.Printf("快速排序结束，耗时%s,循环次数%d, 交换次数%d\n", tr.MicroSec(), qs.CycleCount, qs.SwapCount)
	return v
}


func doSelectionSort(arr []int, debug bool) []int {
	tr.Start();
	fmt.Printf("选择排序开始,总共%d个元素\n", len(arr))
	ss := sort.NewSS(debug)
	var v = cpy(arr)
	tr.End()
	ss.Sort(v)
	fmt.Printf("选择排序结束，耗时%s,循环次数%d, 交换次数%d\n", tr.MicroSec(), ss.CycleCount, ss.SwapCount)
	return v
}

func doSort (length int, output bool)  {
	v := tool.MakeArr(length)
	v1 := doQuickSort(v, false)
	v2 := doBubbleSort(v, false)
	v3 := doSelectionSort(v, false)
	if output {
		fmt.Println(v)
		fmt.Println(v1)
		fmt.Println(v2)
		fmt.Println(v3)
	}
}

func doTest()  {
	v := tool.MakeArr(10)
	v1, v2 := cpy(v), cpy(v)
	fmt.Println(v, v1, v2)
	v[0] = 111
	fmt.Println(v, v1, v2)
}

var tr = tool.TimeRecord{}

func main() {
	doSort(200, false)
}

