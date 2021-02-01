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

func (a *App) signUp(c *gin.Context) {
	user := &SignUpUser{}
	if err := c.Bind(user); err != nil {
		errorJSON := &ErrorRequest{
			Error:   err.Error(),
			Message: "не удачная обработка данных",
			Errors:  nil,
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, errorJSON)
	}
	c.JSON(http.StatusCreated, "")
}
