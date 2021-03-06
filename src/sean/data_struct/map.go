package main

import "fmt"

func main() {

	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	v1_, p1_ := m["k1"]
	fmt.Println("v1_:", v1_)
	fmt.Println("p1_", p1_)

	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}

// 错误示例
type info struct {
	result int
}

func work() (int, error) {
	return 3, nil
}
