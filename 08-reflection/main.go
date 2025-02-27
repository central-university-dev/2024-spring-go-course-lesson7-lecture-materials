package main

import (
	"fmt"
	"reflect"
)

func main() {

	var e float64 = 2.7182818284

	valOfE := reflect.ValueOf(e)
	valOfE.SetFloat(3.1415) // panic: reflect: reflect.Value.SetFloat using unaddressable value
	fmt.Printf("new value of e: %f.7\n", valOfE.Interface().(float64))

	// Адресуемость - свойство переменной, которое определяет, может ли значение, которое хранится в переменной, изменяться.
	//
	// Устанавливаемость — свойство reflect.Value, и не каждое reflect.Value имеет его.
	//
	// Устанавливаемость немного напоминает адресуемость, но строже.
	// Это свойство, при котором reflection объект может изменить хранимое значение, которое было использовано при создании reflection объекта.
	// Устанавливаемость определяется тем, содержит ли reflection объект исходный элемент, или только его копию!
	//
	// В этом все дело!

	//fmt.Printf("valOfE can be set: %v\n", valOfE.CanSet())
	//
	//p := reflect.ValueOf(&e)
	//fmt.Printf("type of p: %s\n", p.Type())
	//fmt.Printf("p can be set: %v\n", p.CanSet())
	//
	//v := p.Elem() // get original Value
	//fmt.Printf("type of v: %s\n", v.Type())
	//fmt.Printf("v can be set: %v\n", v.CanSet())
	//v.SetFloat(3.1415)
	//
	//fmt.Println(v.Interface())
	//fmt.Println(e)
	//
	//// To же самое справедливо и для структур!
	//
	//type T struct {
	//	A int
	//	B string
	//}
	//
	//t := T{37, "bar"}
	//s := reflect.ValueOf(&t).Elem() // когда у нас есть !адрес! структуры мы можем изменять ее поля!
	//
	//for i := 0; i < s.NumField(); i++ {
	//	f := s.Field(i)
	//	fmt.Printf("%d: %s %s = %v\n", i, s.Type().Field(i).Name, f.Type(), f.Interface())
	//}
	//
	//s.Field(0).SetInt(42)
	//s.Field(1).SetString("foo")
	//fmt.Printf("t is now: %+v\n", t)
}
