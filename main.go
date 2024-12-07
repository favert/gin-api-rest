package main

import (
	"github.com/favert/api-go-gin/database"
	"github.com/favert/api-go-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
