package routes

import (

	"github.com/gin-gonic/gin"

	"github.com/eu-micaeu/TocToc/handlers"

	"database/sql"
)

func MensagemRoutes(r *gin.Engine, db *sql.DB) {
	mensagemHandler := handlers.Mensagem{}

	r.POST("/enviar", mensagemHandler.Enviar(db))

}
