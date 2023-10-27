package handlers

// Importando bibliotecas para a criação da classe e funções do usuário.
import (

	"database/sql"

	"github.com/gin-gonic/gin"

)

// Estrutura do usuário.
type Usuario struct {

	ID_Usuario 	int    `json:"id_usuario"`

	Email   	string `json:"email"`

	Senha      	string `json:"senha"`

}

// Função com finalidade de login do usuário.
func (u *Usuario) Entrar(db *sql.DB) gin.HandlerFunc {

	return func(c *gin.Context) {

		var usuario Usuario

		if err := c.BindJSON(&usuario); err != nil {

			c.JSON(400, gin.H{"message": "Erro ao fazer login"})

			return

		}

		row := db.QueryRow("SELECT id_usuario, email, senha FROM usuarios WHERE email = $1 AND senha = $2", usuario.Email, usuario.Senha)

		err := row.Scan(&usuario.ID_Usuario, &usuario.Email, &usuario.Senha)

		if err != nil {

			c.JSON(404, gin.H{"message": "Usuário ou senha incorretos"})

			return

		}

		c.JSON(200, gin.H{"message": "Login efetuado com sucesso!"})

	}

}