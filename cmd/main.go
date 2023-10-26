package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/eu-micaeu/TocToc/middlewares"
)

func main() {

	r := gin.Default()

	r.Use(middlewares.CorsMiddleware())

	r.LoadHTMLGlob("./views/*.html")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.Static("./static", "./static")

	r.Run()
}
