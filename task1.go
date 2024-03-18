package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

func (h Human) Jump() {
	fmt.Printf("Human with name %s is jumping", h.Name)
}

type Action struct {
	// Здесь происходит встраивание
	Human
}

func main() {
	act := Action{Human{
		Name: "Alex",
		Age:  19,
	}}
	act.Jump()
}
