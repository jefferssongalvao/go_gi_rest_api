package main

import (
	"go-gin-rest-api/database"
	"go-gin-rest-api/routes"
)

func main() {
	database.Connect()
	routes.HandleRequests()
}
