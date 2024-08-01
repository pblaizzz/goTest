package main

import "fmt"

func main() {
	var a *string
	fmt.Println(a)
	fmt.Println("p的值是%s/n", a)
	if a != nil {
		fmt.Println("非空")
	} else {
		fmt.Println("空")
	}
}
