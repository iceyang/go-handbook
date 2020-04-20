package main

import "github.com/google/wire"

func InitializeUserController() *UserController {
	wire.Build(NewUserController, NewUserService, NewDatabase)
	return nil
}

func main() {
	c := InitializeUserController()
	c.Hello()
}
