package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

// 多个返回值
func vals() (int, int) {
	return 3, 7
}

// 可变参数
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println("total =", total)
}

// 匿名函数，函数也可作为一种类型存在
func intSeq() func() (int, int) {
	i := 0
	j := 0
	return func() (int, int) {
		i++
		j++
		return i, j
	}
}

func main() {
	res := plus(1, 2)
	fmt.Println("1 + 2 =", res)

	a, b := vals()
	fmt.Println(a, b)

	_, c := vals()
	fmt.Println(c)

	sum(1, 2, 3, 4)
	// 可接收slice结构，数组不可使用
	sliceNum := []int{1, 2, 3, 4}
	sum(sliceNum...)

	// 此时nextInt参数是个函数
	// TODO 暂未理解这个i的变化
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInt := intSeq()
	fmt.Println(newInt())
}
