package main

import "fmt"

func main() {
	currency := map[string]string{"India": "inr", "USA": "dollar", "Europe": "pound"}
	c, ok := currency["India"]   // define to var it returns two values
	currency["Internet"] = "btc" // append to map
	fmt.Println(c, ok)
	fmt.Println(currency)
	for k := range currency {
		if val, ok := currency[k]; ok { //check if map exists
			fmt.Println(val)
		}
	}
}
