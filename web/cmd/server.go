package main

import "github.com/iceyang/m-go-cookbook/web/internal/router"

func main() {
	router.Default().Run(":7900")
}
