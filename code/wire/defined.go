package main

import "fmt"

type Database struct{}

func NewDatabase() *Database {
	return &Database{}
}

type UserService struct {
	Database *Database
}

func (s UserService) SayHello() {
	fmt.Println("Hello!")
}

func NewUserService(database *Database) *UserService {
	return &UserService{Database: database}
}

type UserController struct {
	UserService *UserService
}

func (c UserController) Hello() {
	c.UserService.SayHello()
}

func NewUserController(userService *UserService) *UserController {
	return &UserController{UserService: userService}
}
