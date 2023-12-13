package handlers

import (
	"database/sql"

	"net/http"

	"time"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

type Mensagem struct {

	ID_Mensagem int `json:"id_mensagem"`

	Nickname string `json:"nickname"`

	Texto string `json:"texto"`

	DataDeEnvio string `json:"data_de_envio"`

}

func (m *Mensagem) Enviar(db *sql.DB) gin.HandlerFunc {

	return func(c *gin.Context) {

		token, err := c.Request.Cookie("token")

		if err != nil {

			c.JSON(401, gin.H{"message": "Token inválido"})

			return

		}

		tokenValue := token.Value

		_, err = ValidarOToken(db, tokenValue)

		if err != nil {

			c.JSON(401, gin.H{"message": "Token inválido"})

			return

		}

		var mensagem Mensagem

		if err := c.ShouldBindJSON(&mensagem); err != nil {

			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

			return
		}

		query := "INSERT INTO mensagens (nickname, texto, data_de_envio) VALUES ($1, $2, $3)"

		_, err = db.Exec(query, mensagem.Nickname, mensagem.Texto, time.Now())

		if err != nil {

			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert message into database"})

			return

		}

		c.JSON(http.StatusOK, gin.H{"message": "Message saved successfully"})

	}

}

func (m *Mensagem) Listar(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		query := "SELECT * FROM mensagens ORDER BY data_de_envio ASC"
		rows, err := db.Query(query)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve messages from database"})
			return
		}
		defer rows.Close()

		var mensagens []Mensagem
		for rows.Next() {
			var mensagem Mensagem
			err := rows.Scan(&mensagem.ID_Mensagem, &mensagem.Nickname, &mensagem.Texto, &mensagem.DataDeEnvio)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve messages from database"})
				return
			}
			mensagens = append(mensagens, mensagem)
		}

		c.JSON(http.StatusOK, mensagens)
	}
}
