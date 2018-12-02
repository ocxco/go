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
	for i := 0; i < len(arr) ; i++ {
		for j := 0; j < i; j++ {
			if arr[i] < arr[j] {
				if this.debug {
					fmt.Printf("%v - %d %d - ", arr, i, j)
				}
				arr[i], arr[j] = arr[j], arr[i]
				this.SwapCount++
				if this.debug {
					fmt.Println(arr)
				}
			}
			this.CycleCount++
		}
	}
}
