package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(111111, i)
	}

	//v1.22 之前报错
	for i := range 10 {
		// do something
	}
}