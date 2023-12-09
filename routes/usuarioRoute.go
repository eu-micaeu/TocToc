package routes

import (
	
	"database/sql"

	"github.com/gin-gonic/gin"

	"github.com/eu-micaeu/TocToc/handlers"

)

func UsuarioRoutes(r *gin.Engine, db *sql.DB) {

	usuarioHandler := handlers.Usuario{}

	r.POST("/login", usuarioHandler.Entrar(db))

	r.POST("/cadastrar", usuarioHandler.Cadastrar(db)) 

}
