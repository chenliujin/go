package main

import (
       "github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/") 
	{
		v1.GET("/", getMessage)
	}

	router.Run(":80")
}


func getMessage(c *gin.Context) {
    c.JSON(200, gin.H{"message": "hello world!"})
}
