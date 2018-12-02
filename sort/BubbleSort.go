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
		for j := i + 1; j < len(arr) ; j++ {
			if arr[i] > arr[j] {
				if this.debug {
					fmt.Printf("交换之前, %v %d %d ---", arr, i, j)
				}
				arr[i], arr[j] = arr[j], arr[i]
				this.SwapCount++
				if this.debug {
					fmt.Println("交换之后, ", arr, i, j)
				}
			}
			this.CycleCount++
		}
	}
}
