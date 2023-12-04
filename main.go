package main

import (
	"github.com/fatwahazjuang/go-rest-api/database"
	"github.com/fatwahazjuang/go-rest-api/router"
)

func main() {
	database.ConnectDatabase()
	r := router.RouteInit()
	r.Run(":8080")
}
