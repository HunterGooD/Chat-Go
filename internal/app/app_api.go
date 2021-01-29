package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func helloWorld(c *gin.Context) {
	type Response struct {
		Name string `json:"name"`
		Key  string `json:"key"`
	}
	obj := Response{
		Name: "Hello",
		Key:  "World",
	}
	c.JSON(http.StatusOK, obj)
}
