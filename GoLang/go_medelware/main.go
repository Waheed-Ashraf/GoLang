package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Welcome to your profile"})
}

func checkAge() gin.HandlerFunc {
	return func(c *gin.Context) {
		age, err := strconv.Atoi(c.Param("age"))
		if err != nil || age < 0 {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "Age must be a non-negative integer",
			})
			return
		}

		if age < 18 {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": "Sorry You are underage",
			})
			return
		}

		c.Next()
	}
}

func main() {
	router := gin.Default()
	router.GET("/profile/:age", checkAge(), home)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
