package main

import (
	"log"
	"reflect"
)

type People struct {
	Name string
	Age  int
}
type User3 struct {
	Id int64
	People
	friends []*People
	ext     map[string]interface{}
}

func main() {

	user := User3{
		Id: 1,
		People: People{
			Name: "张三",
			Age:  18,
		},
		friends: []*People{
			{
				Name: "李四",
				Age:  18,
			},
			{
				Name: "朱五",
				Age:  18,
			},
		},
		ext: map[string]interface{}{
			"teacher": People{
				Name: "杨六",
				Age:  36,
			},
		},
	}
	Handle(user)
}
func Handle(i interface{}) {
	val := reflect.ValueOf(i)
	route(val)
}

func route(val reflect.Value) {
	if !val.IsValid() {
		return
	}
	switch val.Kind() {
	case reflect.Struct:
		handleStruct(val)
	case reflect.Slice:
		handleSlice(val)
	case reflect.Map:
		handleMap(val)
	case reflect.Ptr, reflect.Interface:
		route(val.Elem())
	default:
		log.Printf("处理能力之外的Kind，Kind:%+v", val.Kind())
	}
}
func handleStruct(obj reflect.Value) {
	objType := obj.Type()
	for i := 0; i < obj.NumField(); i++ {
		field := obj.Field(i)

		switch field.Kind() {
		case reflect.Array, reflect.Slice, reflect.Chan, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Struct:
			route(field)
		default:
			log.Printf("%s:%v", objType.Field(i).Name, field)
		}

	}
}
func handleSlice(list reflect.Value) {
	for i := 0; i < list.Len(); i++ {
		element := list.Index(i)
		route(element)
	}
}
func handleMap(m reflect.Value) {
	for iter := m.MapRange(); iter.Next(); {
		route(iter.Value())
	}
}
