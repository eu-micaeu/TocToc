package main

import (
	
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/eu-micaeu/TocToc/middlewares"

	"github.com/eu-micaeu/TocToc/database"

	"github.com/eu-micaeu/TocToc/routes"

)

func main() {

	r := gin.Default()

	r.Use(middlewares.CorsMiddleware())

	db, err := database.NewRedisClient()

	if err != nil {

		panic(err)

	}

	routes.UsuarioRoutes(r, db)

	routes.MensagemRoutes(r, db)

	r.LoadHTMLGlob("./views/*.html")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", nil)
	})

	r.GET("/cadastro", func(c *gin.Context) {
		c.HTML(http.StatusOK, "cadastro.html", nil)
	})

	r.Static("./static", "./static")

	r.Run()
}
