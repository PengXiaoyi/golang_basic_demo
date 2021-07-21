package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("run reflect demo fun 1")
	reflectdemo1()
	fmt.Println("run reflect demo fun 2")
	reflectdemo2()
	fmt.Println("run reflect demo fun 3")
	reflectdemo3()
	fmt.Println("run reflect demo fun 4")
	reflectdemo4()
	fmt.Println("run reflect demo fun 5")
	reflectdemo5()
}

func reflectdemo1() {
	author := "draven"
	fmt.Println("TypeOf author:", reflect.TypeOf(author))
	fmt.Println("ValueOf author:", reflect.ValueOf(author))

	v := reflect.ValueOf(1)
	// 从反射对象强转回普通对象
	i := v.Interface().(int)
	fmt.Println("ValueOf Number", i)
}

func reflectdemo2() {
	// 这段会报错，因为反射不允许修改无法被更改的值
	/*	i := 1
		v := reflect.ValueOf(i)
		v.SetInt(10)emptyInterface
		fmt.Println(i)*/

	// 可以用指针的方式操作
	i := 1
	v := reflect.ValueOf(&i)
	v.Elem().SetInt(10)
	fmt.Println(i)
}

//demo3 判断类型是否实现了某个接口

type CustomError struct{}

// CustomError类型没有实现借口，*CustomError指针类型实现了接口。（指针和结构体不属于同一种类型）
func (*CustomError) Error() string {
	return ""
}

func reflectdemo3() {
	typeOfError := reflect.TypeOf((*error)(nil)).Elem()
	customErrorPtr := reflect.TypeOf(&CustomError{})
	customError := reflect.TypeOf(CustomError{})

	fmt.Println(customErrorPtr.Implements(typeOfError)) // #=> true
	fmt.Println(customError.Implements(typeOfError))    // #=> false
}

// demo4 通过反射执行Add方法
//通过 reflect.ValueOf 获取函数 Add 对应的反射对象；
//调用 reflect.rtype.NumIn 获取函数的入参个数；
//多次调用 reflect.ValueOf 函数逐一设置 argv 数组中的各个参数；
//调用反射对象 Add 的 reflect.Value.Call 方法并传入参数列表；
//获取返回值数组、验证数组的长度以及类型并打印其中的数据；

func Add(a, b int) int { return a + b }

func reflectdemo4() {
	v := reflect.ValueOf(Add)
	if v.Kind() != reflect.Func {
		return
	}
	t := v.Type()
	argv := make([]reflect.Value, t.NumIn())
	for i := range argv {
		if t.In(i).Kind() != reflect.Int {
			return
		}
		argv[i] = reflect.ValueOf(i)
	}
	result := v.Call(argv)
	if len(result) != 1 || result[0].Kind() != reflect.Int {
		return
	}
	fmt.Println(result[0].Int()) // #=> 1
}

type Point struct {
	X int
	Y int
	Z int
}
type User struct {
	Name string
	Age  int
	p    Point
}

func reflectdemo5() {
	user := User{
		Name: "张三",
		Age:  18,
	}
	userType := reflect.TypeOf(user)
	index := []int{2} //代表下标为2的字段是一个内嵌结构体，取出这个内嵌结构体中下标为1的字段
	field := userType.FieldByIndex(index)
	fmt.Printf("userType.FieldByIndex():%+v\n", field)

	i := 1
	v1 := reflect.ValueOf(&i)
	v2 := v1.Elem()
	fmt.Printf("Type:%+v\n", v1.Type())

	//可用Value.CanAddr()判断是不是可以被寻址。
	fmt.Println(v1.CanAddr()) // false
	fmt.Println(v2.CanAddr()) // true
}
