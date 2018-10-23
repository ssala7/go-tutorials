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
}
