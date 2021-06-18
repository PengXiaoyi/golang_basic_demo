package main

import (
	"encoding/json"
	"fmt"
)

type Person2 struct {  //Person2结构体，包含年龄，名称，车
	age  int
	name string
	car  Car
}

type Car struct {  //Person2名下的车
	name string  //车的名字
}

var Person2Map map[string]Person2   //一个存放Person2的map

func setName(Person2 Person2, name string) { //给参数Person2设置名字
	Person2.name = name
}

func (Person2 Person2) setName(name string) {  //设置名字
	Person2.name = name
}
func printName(Person2 Person2){  //打印Person2的名字
	fmt.Println(Person2.name)
}
func (Person2 Person2)printName(){  //结构体Person2自己支持打印名字
	fmt.Println(Person2.name)
}

func main() {
	Person2 := Person2{}
	fmt.Println(Person2)  //{0  {}}
	Person2.age = 12
	Person2.name = "小明"
	Person2.car = Car{"宝马"}
	fmt.Println(Person2)  //{12 小明 {宝马}}，正常赋值给Person2变量，因为这是在方法里面的变量
	setName(Person2, "小红")
	fmt.Println(Person2) //{12 小明 {宝马}}，小红赋值失败，传递给setName方法的Person2没有赋值成功
	Person2.setName("小红")
	fmt.Println(Person2) //{12 小明 {宝马}}，Person2自己setName，还是失败
	//Person2Map = make(map[string]Person2)
	//Person2Map["test"] = Person2
	//Person2 = Person2Map["test"]
	//Person2.name = "小红"
	//fmt.Println(Person2) //{12 小红 {宝马}},从map中取出Person2，给小红赋值成功
	//for _, value := range Person2Map { //遍历map
	//	fmt.Println(value)//{12 小明 {宝马}}，打印的还是小明，而不是小红，说明上面Person2Map["test"]对象赋值失败
	//}
	whiteList := EncryptedNumberWhiteList{TenantIdList: []int64{1, 2, 3, 4}}
	marshal, _ := json.Marshal(whiteList)
	fmt.Println(string(marshal))
}

type EncryptedNumberWhiteList struct {
	TenantIdList []int64 `json:"tenant_id_list"`
}