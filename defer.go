package main

import "fmt"

func deferExample() {
	defer func() {
		fmt.Println("Do something finally.")
	}()
	fmt.Println("here we go.")
}
