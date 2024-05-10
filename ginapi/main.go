package main

import (
	"ginapi/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, models.Users)
}

func getUsersByID(c *gin.Context) {
	id := c.Param("id")
	for _, u := range models.Users {
		if u.ID == id {
			c.IndentedJSON(http.StatusOK, u)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found"})
}

func postUsers(c *gin.Context) {
	var newUser = models.User{}
	if err := c.BindJSON(&newUser); err != nil {
		return
	}
	models.Users = append(models.Users, newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}

func main() {
	router := gin.Default()
	router.GET("/users", getUsers)
	router.GET("/users/:id", getUsersByID)
	router.POST("/users", postUsers)

	router.Run("localhost:8080")
}
