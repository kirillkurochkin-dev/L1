package main

import (
	"fmt"
	"reflect"
)

func main() {
	var (
		runInt    interface{}
		runString interface{}
		runBool   interface{}
		runCh     interface{}
	)

	runInt = 10
	runString = "String"
	runBool = false
	runCh = make(chan string)

	fmt.Println(reflect.TypeOf(runInt))
	fmt.Println(reflect.TypeOf(runString))
	fmt.Println(reflect.TypeOf(runBool))
	fmt.Println(reflect.TypeOf(runCh))
}
