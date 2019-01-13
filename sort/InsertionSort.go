package sort

import "fmt"

type InsertionSort struct {
	debug bool
	CycleCount int
	SwapCount int
}


func NewIS(debug bool) InsertionSort  {
	return InsertionSort{
		debug,
		0,
		0,
	}
}

func (this *InsertionSort) Sort(arr []int)  {
	for i := 1; i < len(arr) ; i++ {
		preIndex := i - 1
		current := arr[i]
		for  ;preIndex >= 0 && arr[preIndex] > current; preIndex-- {
			if this.debug {
				fmt.Printf("%v - %d %d - ", arr, i, preIndex)
			}
			this.CycleCount++
			this.SwapCount++
			arr[preIndex + 1] = arr[preIndex]
			if this.debug {
				fmt.Println(arr)
			}
		}
		arr[preIndex + 1] = current
	}
}
