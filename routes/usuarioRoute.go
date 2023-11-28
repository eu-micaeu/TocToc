package routes

import (

	"github.com/gin-gonic/gin"

	"github.com/eu-micaeu/TocToc/handlers"

	"github.com/go-redis/redis/v8"

)

func UsuarioRoutes(r *gin.Engine, client *redis.Client) {
	usuarioHandler := handlers.Usuario{}

	r.POST("/login", usuarioHandler.Entrar(client, "TOCTOCKEYSECRECT"))

	r.POST("/cadastrar", usuarioHandler.Cadastrar(client)) 

}
