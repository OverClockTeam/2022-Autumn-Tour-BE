package main

import (
	"web1/model"
	"web1/routes"
)

func main() {
	model.InitDb()

	routes.InitRouter()
}
