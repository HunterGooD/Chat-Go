package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/packr/v2"
)

func main() {

	htmlFiles := packr.New("htmlFiles", "../../web/dist/")
	router := gin.Default()
	assetsBox := packr.New("assets", "../../web/dist/assets")
	router.StaticFS("/assets/", assetsBox)

	router.NoRoute(func(c *gin.Context) {
		data, err := htmlFiles.FindString("index.html")
		if err != nil {
			panic("html not found")
		}
		c.Data(http.StatusOK, "text/html;charset=utf-8", []byte(data))
	})

	router.Run(":9090")
}
