package main

import "fmt"

func main() {
	var a = [5]int{1, 2, 3, 4, 5}
	s := a[0:3]
	s = append(s, 4, 1)
	fmt.Println(s)
}
