package handlers

import (

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type Usuario struct {
	Nickname string `json:"nickname"`
	Senha    string `json:"senha"`
}

func (u *Usuario) Entrar(client *redis.Client) gin.HandlerFunc {

	return func(c *gin.Context) {

		var usuario Usuario

		if err := c.BindJSON(&usuario); err != nil {

			c.JSON(400, gin.H{"message": "Erro ao fazer login"})

			return

		}

		val, _ := client.Get(client.Context(), "usuario:" + usuario.Nickname).Result()

		if val != usuario.Senha {

			c.JSON(401, gin.H{"message": "Usuário ou senha incorretos"})

			return

		}

		c.JSON(200, gin.H{"message": "Login efetuado com sucesso!"})
	}
}
