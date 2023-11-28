package routes

import (

	"github.com/gin-gonic/gin"

	"github.com/eu-micaeu/TocToc/handlers"

	"github.com/go-redis/redis/v8"

)

func MensagemRoutes(r *gin.Engine, client *redis.Client) {
	mensagemHandler := handlers.Mensagem{}

	r.POST("/enviar", mensagemHandler.EnviarMensagem(client))

	r.GET("/listar", mensagemHandler.ListarMensagens(client))

}
