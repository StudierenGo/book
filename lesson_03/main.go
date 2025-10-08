package main

import "fmt"

func printType(arg interface{}) {
	if _, ok := arg.(int); ok {
		fmt.Println("Type is number")
	} else if _, ok := arg.(string); ok {
		fmt.Println("Type is string")
	} else {
		fmt.Println("Looks like some shit")
	}
}

func printTypeSwitch(value interface{}) {
	switch value.(type) {
	case int:
		fmt.Println("Type is number")
	case string:
		fmt.Println("Type is string")
	default:
		fmt.Println("Looks like some shit")
	}
}

func main() {
	printType(5)
	printType("What the fuck?")
	printType([]string{"1", "2"})

	printTypeSwitch(12)
	printTypeSwitch("Another attempt")
	printTypeSwitch([2]int{1, 2})
}
