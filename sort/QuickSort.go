package sort

import "fmt"

type QuickSort struct {
	debug bool
	SwapCount int
	CycleCount int
}

func NewQS(debug bool) QuickSort {
	return QuickSort{
		debug,
		0,
		0,
	}
}

func (this *QuickSort) partition(arr []int) int {
	head, base := 0, arr[0]
	tail := len(arr) - 1
	for head < tail {
		if this.debug {
			fmt.Printf("交换之前, %v %d %d ---", arr, head, tail)
		}
		if (arr[head + 1] > base) {
			arr[head + 1], arr[tail] = arr[tail], arr[head + 1]
			tail--
			this.SwapCount++
		} else {
			arr[head + 1], arr[head] = arr[head], arr[head + 1]
			this.SwapCount++
			head++
		}
		this.CycleCount++
		if this.debug {
			fmt.Println("交换之后,", arr, head, tail)
		}
	}
	return head
}

func (this *QuickSort) Sort(arr []int)  {
	if (len(arr) <= 1) {
		return
	}
	k := this.partition(arr);
	this.Sort(arr[:k])
	this.Sort(arr[k+1:])
}