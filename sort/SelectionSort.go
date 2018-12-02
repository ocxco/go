package sort

import "fmt"

// SelectionSort 选择排序
// 每次循环找到数组中最小的值，与i交换
// 每次遍历只需要交换1次
// 时间复杂度为O(n^2)
type SelectionSort struct {
	debug bool
	SwapCount int
	CycleCount int
}

func NewSS(debug bool) SelectionSort  {
	return SelectionSort{
		debug,
		0,
		0,
	}
}

func (this *SelectionSort) Sort(arr []int)  {
	for i := 0; i < len(arr) ; i++ {
		minIndex := i
		for j := i + 1; j < len(arr) ; j++ {
			if arr[minIndex] > arr[j] {
				minIndex = j
			}
			this.CycleCount++
		}
		if i != minIndex {
			if this.debug {
				fmt.Printf("交换之前, %v %d %d ---", arr, i, minIndex)
			}
			arr[minIndex], arr[i] = arr[i], arr[minIndex]
			this.SwapCount++
			if this.debug {
				fmt.Println("交换之后,", arr, i, minIndex)
			}
		}
	}
}
