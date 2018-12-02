package sort

import "fmt"

// BubbleSort 冒泡排序
// 依次比较数组中元素与第i个的值，符合条件的交换位置
// 每次遍历可能交换多次
// 时间复杂度为O(n^2)
type BubbleSort struct {
	debug bool
	SwapCount int
	CycleCount int
}

func NewBS(debug bool) BubbleSort  {
	return BubbleSort{
		debug,
		0,
		0,
	}
}

func (this *BubbleSort) Sort(arr []int)  {
	for i := 0; i < len(arr) ; i++ {
		for j := 0; j < len(arr) - 1; j++ {
			if arr[j] > arr[j+1] {
				if this.debug {
					fmt.Printf("%v - %d %d - ", arr, i, j)
				}
				arr[j], arr[j+1] = arr[j+1], arr[j]
				this.SwapCount++
				if this.debug {
					fmt.Println(arr)
				}
			}
			this.CycleCount++
		}
	}
}
