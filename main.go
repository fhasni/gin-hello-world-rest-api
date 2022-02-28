package main

import (
	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	c.JSON(200, gin.H{
		"data":  "get data",
		"message": "success",
	})
}

func Post(c *gin.Context) {
	c.JSON(200, gin.H{
		"data":  "post data",
		"message": "success",
	})
}

func QueryString(c *gin.Context) {
	name := c.Query("name")
	age := c.Query("age")
	c.JSON(200, gin.H{
		"data":  gin.H{
			"name":  name,
			"age":  age,
		},
		"message": "success",
	})
}

func PathParams(c *gin.Context) {
	name := c.Param("name")
	age := c.Param("age")
	c.JSON(200, gin.H{
		"data":  gin.H{
			"name":  name,
			"age":  age,
		},
		"message": "success",
	})
}

func main() {
	r := gin.Default()

	r.GET("/", Get)
	r.POST("/", Post)
	r.GET("/query", QueryString)
	r.GET("/path/:name/:age", PathParams)

	r.Run(":8080")
}