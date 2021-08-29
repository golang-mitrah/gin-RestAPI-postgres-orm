package main

import (
	"github.com/golang-mitrah/gin-RestAPI-postgres-orm/database"
	"github.com/golang-mitrah/gin-RestAPI-postgres-orm/models"
	"github.com/golang-mitrah/gin-RestAPI-postgres-orm/routes"
)

func main() {

	database.ConnectToDB()
	database.DB.AutoMigrate(&models.User{})

	r := routes.SetupRouter()
	//running
	r.Run()
}
