package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type employee struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var u []employee

func main() {
	r := gin.Default()
	userruotes := r.Group("/users")
	{
		userruotes.GET("/", message)
		userruotes.POST("/", create)
		userruotes.PUT("/:Id", edituser)
		userruotes.DELETE("/:Id", deletuser)
	}

	if err := r.Run(":5000"); err != nil {
		log.Println(err)
	}

}
func message(c *gin.Context) {
	c.JSON(200, u)

}
func create(c *gin.Context) {
	var usertype employee
	if err := c.ShouldBindJSON(&usertype); err != nil {
		c.JSON(422, gin.H{
			"message": "send proper request",
		})
		return
	}
	usertype.Id = uuid.New().String()
	u = append(u, usertype)
	c.JSON(200, gin.H{
		"message": "successful",
	})
}
func edituser(c *gin.Context) {
	id := c.Param("Id")
	var usertype employee
	if err := c.ShouldBindJSON(&usertype); err != nil {
		c.JSON(422, gin.H{
			"message": "send proper request",
		})
		return
	}
	for i, a := range u {
		if a.Id == id {
			u[i].Name = usertype.Name
			u[i].Age = usertype.Age

			c.JSON(200, gin.H{
				"message": "succeful",
			})
			return
		}
	}

}
func deletuser(c *gin.Context) {
	id := c.Param("Id")
	for i, a := range u {
		if a.Id == id {
			u = append(u, u[i+1:]...)

			c.JSON(200, gin.H{
				"message": "successful",
			})
			return
		}
	}

}
