package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("Do something finally.")
	}()
	fmt.Println("here we go.")
}
