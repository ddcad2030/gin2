package controllers

import (
	"log"
	"net/http"

	"github.com/ddcad2030/gin-gorm-rest/config"
	"github.com/ddcad2030/gin-gorm-rest/models"
	"github.com/ddcad2030/gin-gorm-rest/utils"
	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "greeting",
	})
}

func GetUser(c *gin.Context) {
	users := []models.User{}
	result := config.DB.Find(&users)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, result)
		return
	}
	c.JSON(http.StatusOK, &users)
}

func GetUserByID(c *gin.Context) {
	users := models.User{}
	result := config.DB.First(&users, c.Param("id"))

	if result.Error != nil {
		c.JSON(http.StatusNotFound, result.Error)
		return
	}
	c.JSON(http.StatusOK, &users)
}

func CreateUser(c *gin.Context) {
	users := models.User{}
	c.BindJSON(&users)
	hash, err := utils.GenerateToken(users.Email)
	log.Println(hash, err)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	users.Password = string(hash)
	result := config.DB.Create(&users)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, result.Error)
		return
	}
	c.JSON(http.StatusCreated, &users)
}

func UpdateUser(c *gin.Context) {
	users := models.User{}
	result := config.DB.Where("id = ?", c.Param("id")).First(&users)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": result.Error,
		})
		return
	}
	var body struct {
		Email    string
		Password string
	}
	c.BindJSON(&body)

	resultUpdate := config.DB.Model(&users).Updates(&body)

	if resultUpdate.Error != nil {
		c.JSON(http.StatusNotFound, resultUpdate.Error)
		return
	}
	c.JSON(http.StatusOK, &users)
}

func DeleteUser(c *gin.Context) {
	users := models.User{}
	search := config.DB.First(&users, c.Param("id"))

	if search.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": search.Error,
		})
		return
	}

	config.DB.Delete(&users)

	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})

}

func Login(c *gin.Context) {
	var body struct {
		Email    string
		Password string
	}
	users := models.User{}
	c.BindJSON(&body)

	Findresult := config.DB.Where("email = ?", body.Email).First(&users)
	if Findresult.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": Findresult.Error,
		})
		return
	}

	CheckHash := utils.CheckPasswordHash(body.Password, users.Password)

	if !CheckHash {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Password not true",
		})
		return
	}

	token, err := utils.GenerateToken(users.Email)

	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "Invalid token",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
