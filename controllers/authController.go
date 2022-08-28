package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zennon-sml/go-jwt-react/database"
	"github.com/zennon-sml/go-jwt-react/models"
	"github.com/zennon-sml/go-jwt-react/services"
)

func Users(c *gin.Context) {
	var users []models.User
	database.DB.Debug().Find(&users)
	c.JSON(200, users)
}

func Register(c *gin.Context) {
	user := models.User{}
	if err := c.ShouldBind(&user); err != nil {
		log.Fatal(err)
	}
	user.Password = services.SHA256Encoder(user.Password)

	database.DB.Debug().Create(&user)
	c.JSON(200, user)
}

func Login(c *gin.Context) {
	user := models.User{}
	userAuth := models.User{}

	if err := c.ShouldBind(&user); err != nil {
		log.Fatal(err)
	}

	database.DB.Debug().Where("email = ?", user.Email).First(&userAuth)

	if userAuth.ID == 0 {
		c.JSON(404, gin.H{
			"message": "user not found",
		})
	} else if user.Password != userAuth.Password {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "wrong credentials",
		})
	} else {
		c.JSON(200, userAuth)
	}

}
