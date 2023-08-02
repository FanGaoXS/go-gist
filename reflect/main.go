package main

import (
	"fmt"
	"reflect"
)

// https://darjun.github.io/2021/05/27/godailylib/reflect/

type Student struct {
	Name  string `json:"name" diy:"s_name"`
	Score int    `json:"score" diy:"s_score"`
}

func (s Student) Print() {
	fmt.Printf("%v", s)
}

func reflectField(x interface{}) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)
	if t.Kind() != reflect.Struct {
		return
	}

	fmt.Printf("fields of struct are %d: \n", t.NumField())
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Printf("field: name: %v, type: %v, tag: %v, value: %v\n", f.Name, f.Type, f.Tag, val)
	}
	// t.FieldByName()
}

func reflectMethod(x interface{}) {
	t := reflect.TypeOf(x)
	//v := reflect.ValueOf(x)
	if t.Kind() != reflect.Struct {
		return
	}

	fmt.Printf("methods of struct are %d: \n", t.NumMethod())
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("method: name: %v\n", m.Name)

		// call the method
		//var args []reflect.Value
		//v.Method(i).Call(args)
	}

	// t.MethodByName()
}

func ReflectStruct(x interface{}) {
	reflectField(x)
	reflectMethod(x)
}

// variable f float64 = 3.5
// <type, value> -> <float64, 3.5>

func main() {
	var a float64 = 3.5
	reflectType(a)

	var b int = 1
	reflectType(b)

	var d []string
	reflectType(d)

	var e []int
	reflectType(e)

	var s = Student{
		Name:  "test",
		Score: 18,
	}
	ReflectStruct(s)
}

func reflectType(x interface{}) {
	t := reflect.TypeOf(x)
	tName := t.Name()
	tKind := t.Kind() // such as bool, int, struct, slice, func, interface ...
	fmt.Printf("type of varible is %s, name is %s\n", tKind, tName)
}
