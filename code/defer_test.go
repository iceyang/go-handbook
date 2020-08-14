package main

import (
	"testing"

	"github.com/fagongzi/log"
)

func do() {
	log.Info("Do something finally.")
}

func TestDefer(t *testing.T) {
	defer do()
	log.Info("here we go.")
}

func TestDefer2(t *testing.T) {
	defer do()
	if true {
		return
	}
	log.Info("here we go.")
}

func TestDefer3(t *testing.T) {
	if true {
		return
	}
	defer do()
	log.Info("here we go.")
}
