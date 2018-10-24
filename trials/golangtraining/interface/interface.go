package main

import "fmt"

type Speaker interface {
	speak()
}

type Dog struct{}

func (d Dog) speak() {
	fmt.Println("bow bow")
}

type Car struct{}

func (c Car) speak() {
	fmt.Println("honk honk")
}

type Human struct{}

func (h Human) speak() {
	fmt.Println("hello")
}

type Alien struct{}

func (a Alien) speak() {
	fmt.Println("&^&^&^&&&^&&^")
}

type T struct{}

func main() {
	speakers := []Speaker{
		Dog{},
		Car{},
		Human{},
		Alien{},
	}

	for _, sp := range speakers {
		sp.speak()
	}
	var i interface{} = 1
	s := i.(int)
	fmt.Println(s)
	f(2)
}

func f(i interface{}) {
	switch v := i.(type) {
	case string:
		println("string:", v)
	case int:
		println("int:", v)
	case int32, int64:
		println("int32 or int64:", v)
	case T:
		fmt.Println("custom type:", v)
	default:
		println("other")
	}

}
