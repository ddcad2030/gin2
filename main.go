package main

import (
	"github.com/ddcad2030/gin-gorm-rest/config"
	"github.com/ddcad2030/gin-gorm-rest/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config.ConnectDB()
	routes.UserRoutes(r)
	r.Run()
}
