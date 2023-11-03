package main

import (
	"todo-app/model"
	"todo-app/routes"
)

func main() {
	model.Setup()
	routes.SetupAndRun()
}
