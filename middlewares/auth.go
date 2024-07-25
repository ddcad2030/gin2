package middlewares

import (
	"log"
	"net/http"

	"github.com/ddcad2030/gin-gorm-rest/utils"
	"github.com/gin-gonic/gin"
)

func Authentication(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	log.Println(token)
	if len(token) < 3 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Not authorization",
		})
	}

	err := utils.VerifyToken(token)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		return
	}

	c.Next()

}
