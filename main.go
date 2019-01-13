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

func doBubbleSort(arr []int, debug bool, output bool) {
	fmt.Printf("冒泡排序开始,总共%d个元素\n", len(arr))
	var v = cpy(arr)
	bs := sort.NewBS(debug)
	tr.Start()
	bs.Sort(v)
	tr.End()
	fmt.Printf("冒泡排序结束，耗时%s, 循环次数%d, 交换次数%d\n", tr.MicroSec(), bs.CycleCount, bs.SwapCount)
	fmt.Println(v)
}

func doQuickSort(arr []int, debug bool, output bool) {
	tr.Start();
	fmt.Printf("快速排序开始,总共%d个元素\n", len(arr))
	qs := sort.NewQS(debug)
	var v = cpy(arr)
	tr.End()
	qs.Sort(v)
	fmt.Printf("快速排序结束，耗时%s,循环次数%d, 交换次数%d\n", tr.MicroSec(), qs.CycleCount, qs.SwapCount)
	fmt.Println(v)
}

func doSelectionSort(arr []int, debug bool, output bool) {
	tr.Start();
	fmt.Printf("选择排序开始,总共%d个元素\n", len(arr))
	ss := sort.NewSS(debug)
	var v = cpy(arr)
	tr.End()
	ss.Sort(v)
	fmt.Printf("选择排序结束，耗时%s,循环次数%d, 交换次数%d\n", tr.MicroSec(), ss.CycleCount, ss.SwapCount)
	fmt.Println(v)
}

func doInsertionSort(arr []int, debug bool, output bool) {
	tr.Start();
	fmt.Printf("插入排序开始,总共%d个元素\n", len(arr))
	is := sort.NewIS(debug)
	var v = cpy(arr)
	tr.End()
	is.Sort(v)
	fmt.Printf("插入排序结束，耗时%s,循环次数%d, 交换次数%d\n", tr.MicroSec(), is.CycleCount, is.SwapCount)
	fmt.Println(v)
}

func doSort (length int, output bool)  {
	v := tool.MakeArr(length)
	if output {
		fmt.Println(v)
	}
	doQuickSort(v, false, output)
	doBubbleSort(v, false, output)
	doSelectionSort(v, false, output)
	doInsertionSort(v, false, output)
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
	doSort(10, true)
}

