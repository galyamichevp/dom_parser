package main

import "go-dom-parser/api/routes"

func main() {

	router := routes.SetupRouter()

	router.Run()
}
