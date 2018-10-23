package main

import "fmt"
import "errors"

func spm(i, j int) (int, int, int, int, error) {

	if j == 0 {
		return i + j, i * j, i - j, 0, errors.New("you can't divide by zero")
	}
	return i + j, i * j, i - j, i / j, nil
}

func main() {
	fmt.Println(spm(100, 100))
}
