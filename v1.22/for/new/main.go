package main

import "fmt"

func main() {
	//v1.22 之前报错
	for i := range 10 {
		fmt.Println(111111, i)
	}
}