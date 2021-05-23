package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {

	p := point{1, 2}
	fmt.Printf("%v\n", p)
	// 如果该值为结构，则该%+v变体将包含该结构的字段名称
	fmt.Printf("%+v\n", p)
	// 该%#v变体打印该值的Go语法表示形式，即将产生该值的源代码片段
	fmt.Printf("%#v\n", p)
	// 打印值的类型
	fmt.Printf("%T\n", p)
	// 格式化bool值
	fmt.Printf("%t\n", true)
	// 格式化整数
	fmt.Printf("%d\n", 123)
	// 打印二进制的表示形式
	fmt.Printf("%b\n", 14)
	// 打印unicode的字符
	fmt.Printf("%c\n", 33)
	// 16进制的编码
	fmt.Printf("%x\n", 456)
	// 浮点数
	fmt.Printf("%f\n", 78.9)
	// 科学计数法
	fmt.Printf("%e\n", 123400000.0)
	fmt.Printf("%E\n", 123400000.0)
	// 字符串打印
	fmt.Printf("%s\n", "\"string\"")
	// 对字符串加""
	fmt.Printf("%q\n", "\"string\"")
	//
	fmt.Printf("%x\n", "hex this")

	fmt.Printf("%p\n", &p)

	fmt.Printf("|%6d|%6d|\n", 12, 345)

	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)

	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)

	fmt.Printf("|%6s|%6s|\n", "foo", "b")

	fmt.Printf("|%-6s|%-6s|\n", "foo", "b")

	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)

	fmt.Fprintf(os.Stderr, "an %s\n", "error")
}