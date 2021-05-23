package main

type quickSort struct {
}

func (s quickSort) sort(nums []int) []int {

	return nums
}

// TODO 没写完
func sortNotStable(nums []int, l, r int) {
	if l >= r {
		return
	}
	p := nums[l]
	i, j := l, r
	for i < j {
		for {
			if nums[j] <= p {
				break
			}
			j--
		}
		for {
			if nums[i] >= p {
				break
			}
			i++
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	sortNotStable(nums, l, j)
	sortNotStable(nums, j+1, r)
}
