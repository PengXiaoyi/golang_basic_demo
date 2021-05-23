package main

import "fmt"

type sort interface {
	sort(nums []int) []int
}

func main() {

	nums := []int{5, 24, 65, 7, 2, 46, 97, 5, 34, 6}
	fmt.Println(quickSort{}.sort(nums))
}
