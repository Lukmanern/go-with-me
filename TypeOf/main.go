package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(TypeOf("hehehe"))
	fmt.Println(TypeOf("1.2"))
	fmt.Println(TypeOf(1.4))
	fmt.Println(TypeOf(11))
	fmt.Println(TypeOf([]int{}))
	fmt.Println(TypeOf([]int{1,2,3,4,}))
	fmt.Println(TypeOf([]string{"1", "2", "3", "4",}))
	fmt.Println(TypeOf([]interface{}{0, "2", "3", "4",}))
	fmt.Println(TypeOf([]any{0, "2", "3", "4",}))
}

// any == interface{}
func TypeOf(data any) string {
	return reflect.TypeOf(data).String()
}