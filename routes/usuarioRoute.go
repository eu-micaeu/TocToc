package routes

import (

	"github.com/gin-gonic/gin"

	"github.com/eu-micaeu/TocToc/handlers"

	"github.com/go-redis/redis/v8"

)

func UsuarioRoutes(r *gin.Engine, client *redis.Client) {
	userHandler := handlers.Usuario{}

	r.POST("/login", userHandler.Entrar(client, "TOCTOCKEYSECRECT"))

	r.POST("/cadastrar", userHandler.Cadastrar(client)) 

}
