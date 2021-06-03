package main

import "encoding/json"

type student struct {
	Name string
}

// 函数参数传值
func checkParam(param student) (returnParam student) {
	println("函数参数传值地址：", &param)
	returnParam = student{
		Name: "After",
	}
	println("函数参数传值返回值地址：", &returnParam)
	return
}

// 函数参数传引用
func checkPoint(param *student) (returnPoint *student) {
	println("函数指针传值地址：", param)
	returnPoint = &student{
		Name: "After",
	}
	println("函数指针传值返回值地址：", returnPoint)
	return
}

func main() {
	// 由地址可以看出，无论是函数参数还是返回值，都会把传值的数据复制一遍
	// 由地址的分配顺序可以看出，返回值是在调用之前就已经进行了地址分配
	stu := student{
		Name: "Before",
	}
	returnParam := checkParam(stu)
	println("函数参数传值地址：", &stu)
	println("函数参数返回值地址：", &returnParam)

	// 是指针的时候，指针会复制一遍
	stuPoint := &student{
		Name: "Before",
	}
	returnPoint := checkPoint(stuPoint)
	println("函数指针传值地址：", stuPoint)
	println("函数指针返回值地址：", returnPoint)
	// 函数传指针时，指针会复制
	println(&stuPoint)
	println(&returnPoint)
	println(json.Marshal(&returnPoint))
}


