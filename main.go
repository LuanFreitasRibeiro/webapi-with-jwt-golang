package main

import (
	"github.com/LuanFreitasRibeiro/webapi-mvc-golang.git/database"
	"github.com/LuanFreitasRibeiro/webapi-mvc-golang.git/server"
)

func main() {
	database.StartDB()

	server := server.NewServer()

	server.Run()
}
