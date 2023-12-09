package handlers

// Importando bibliotecas para a criação da classe e funções do usuário.
import (

	"database/sql"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"

)

// Estrutura do usuário.
type Usuario struct {

	ID_Usuario 	int    `json:"id_usuario"`
	Nickname   	string `json:"nickname"`
	Senha      	string `json:"senha"`

}

func (u *Usuario) Entrar(db *sql.DB) gin.HandlerFunc {

	return func(c *gin.Context) {

		var usuario Usuario

		if err := c.BindJSON(&usuario); err != nil {

			c.JSON(400, gin.H{"message": "Erro ao fazer login"})

			return

		}

		row := db.QueryRow("SELECT id_usuario, nickname, senha FROM usuarios WHERE nickname = $1 AND senha = $2", usuario.Nickname, usuario.Senha)

		err := row.Scan(&usuario.ID_Usuario, &usuario.Nickname, &usuario.Senha)

		if err != nil {

			c.JSON(404, gin.H{"message": "Usuário ou senha incorretos"})

			return

		}

		token, _ := GerarOToken(usuario)

		http.SetCookie(c.Writer, &http.Cookie{

			Name:     "token",

			Value:    token,

			Expires:  time.Now().Add(24 * time.Hour),

			HttpOnly: true,

			Secure:   true,

			SameSite: http.SameSiteStrictMode,

		})

		c.JSON(200, gin.H{"message": "Login efetuado com sucesso!", "token": token, "usuario": usuario})

	}

}


func (u *Usuario) Cadastrar(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		var novoUsuario Usuario

		if err := c.BindJSON(&novoUsuario); err != nil {

			c.JSON(400, gin.H{"message": "Erro ao criar usuario"})

			return

		}

		_, err := db.Exec("INSERT INTO usuarios (nickname, senha) VALUES ($1, $2)", novoUsuario.Nickname, novoUsuario.Senha)

		if err != nil {

			c.JSON(500, gin.H{"message": "Erro ao criar usuário"})

			return

		}

		c.JSON(200, gin.H{"message": "Usuário criado com sucesso!"})

	}
}


