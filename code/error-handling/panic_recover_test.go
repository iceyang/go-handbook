package main

import (
	"errors"
	"log"
	"testing"
)

func Recover() {
	if err := recover(); err != nil {
		log.Printf("%s\n", err)
	}
}

func TestPanic(t *testing.T) {
	defer Recover()
	log.Println("Here we go")
	panic(errors.New("An error occurs"))
	// unreachable code because of panic
	log.Println("Here we go2")
}
