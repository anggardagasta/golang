package main

import "fmt"

func main() {
	for i := 4; i >= 0; i-- {
		fmt.Print(i)
		fmt.Print(i)
		for j := i; j >= 0; j-- {
			fmt.Print(i + 2)
		}
		fmt.Print("\n")
	}
}
