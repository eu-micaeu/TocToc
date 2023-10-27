package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	
	"github.com/eu-micaeu/TocToc/handlers"
)

func UsuarioRoutes(r *gin.Engine, db *sql.DB) {

	userHandler := handlers.Usuario{}

	r.POST("/login", userHandler.Entrar(db))

}
