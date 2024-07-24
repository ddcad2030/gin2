package routes

import (
	"github.com/ddcad2030/gin-gorm-rest/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	r.GET("/hello", controllers.Hello)

	r.GET("/user", controllers.GetUser)
	r.GET("/user/:id", controllers.GetUserByID)
	r.POST("/user", controllers.CreateUser)
	r.PUT("/user/:id", controllers.UpdateUser)
	r.DELETE("/user/:id", controllers.DeleteUser)
}
