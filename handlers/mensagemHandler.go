package handlers

import (
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type Mensagem struct {
	Usuario     string `json:"usuario"`
	Texto       string `json:"texto"`
	DataDeEnvio string `json:"dataDeEnvio"`
}

func (m *Mensagem) EnviarMensagem(client *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {

		var mensagem Mensagem

		err := c.BindJSON(&mensagem)

		if err != nil {
			fmt.Println("Erro ao obter dados brutos:", err)
			return
		}

		if err != nil {
			c.JSON(400, gin.H{"error": "Erro ao ler o corpo da solicitação"})
			return
		}

		dataDeEnvio := time.Now().Format(time.RFC3339)

		err = client.RPush(client.Context(), "mensagem:" + mensagem.Texto, mensagem.Usuario, dataDeEnvio, 0).Err()

		if err != nil {
			c.JSON(500, gin.H{"message": "Erro ao registrar mensagem"})
			return
		}

		c.JSON(200, gin.H{"mensagem": "Mensagem enviada com sucesso!"})
	}
}

func (m *Mensagem) ListarMensagens(client *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {

		var mensagens []Mensagem

		mensagensRedis, err := client.Keys(client.Context(), "mensagem:*").Result()

		if err != nil {
			c.JSON(500, gin.H{"message": "Erro ao listar mensagens"})
			return
		}

		for _, mensagemRedis := range mensagensRedis {

			texto := mensagemRedis[strings.Index(mensagemRedis, ":")+1:]

			mensagem := Mensagem{
				Texto:       texto,
				Usuario:    client.LIndex(client.Context(), mensagemRedis, 0).Val(),
				DataDeEnvio: client.LIndex(client.Context(), mensagemRedis, 1).Val(),
			}

			mensagens = append(mensagens, mensagem)
		}

		c.JSON(200, gin.H{"mensagens": mensagens})
	}
}
