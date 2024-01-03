package sort

import (
	"fmt"
	"testing"
)

var (
	SortImpl = map[string]ISort{
		"quick_sort": quickSort{},
	}
)

func Test_AllSort(t *testing.T) {
	nums := []int{5, 24, 65, 7, 2, 46, 97, 5, 34, 6}
	for sType, impl := range SortImpl {
		ret := impl.sort(nums)
		fmt.Printf("type:%s,ret:%v\n", sType, ret)
	}
}
