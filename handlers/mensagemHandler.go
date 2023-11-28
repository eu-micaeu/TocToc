package handlers

import (
	"fmt"
	"sort"
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

		// Cria uma slice para armazenar as mensagens
		var mensagens []Mensagem

		// Obtém todas as chaves do padrão "mensagem:*" no Redis
		mensagensRedis, err := client.Keys(client.Context(), "mensagem:*").Result()

		// Verifica se houve algum erro ao obter as chaves
		if err != nil {
			c.JSON(500, gin.H{"message": "Erro ao listar mensagens"})
			return
		}

		// Itera sobre as chaves obtidas do Redis
		for _, mensagemRedis := range mensagensRedis {

			// Extrai o texto da mensagem da chave
			texto := mensagemRedis[strings.Index(mensagemRedis, ":")+1:]

			// Cria uma instância da estrutura Mensagem e popula seus campos
			mensagem := Mensagem{
				Texto:       texto,
				Usuario:     client.LIndex(client.Context(), mensagemRedis, 0).Val(),
				DataDeEnvio: client.LIndex(client.Context(), mensagemRedis, 1).Val(),
			}

			// Adiciona a mensagem à slice de mensagens
			mensagens = append(mensagens, mensagem)
		}

		// Ordena as mensagens pela data de envio (da mais antiga para a mais recente)
		sort.Slice(mensagens, func(i, j int) bool {
			return mensagens[i].DataDeEnvio < mensagens[j].DataDeEnvio
		})

		// Retorna as mensagens em formato JSON
		c.JSON(200, gin.H{"mensagens": mensagens})
	}
}


