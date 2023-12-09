package handlers

// Importando bibliotecas para a criação da classe e funções do usuário.
import (

	"database/sql"
	"log"
	"time"
	"github.com/dgrijalva/jwt-go"

)

// Estrutura do TOKEN.
type Claims struct {

	ID_Usuario int `json:"id_usuario"`
	jwt.StandardClaims

}

var jwtKey = []byte("my_secret_key")

// Função para verificar se o token está na tabela de tokens inválidos
func tokenEstaNaTabelaDeTokensInvalidos(db *sql.DB, tokenString string) bool {

	query := "SELECT COUNT(*) FROM tokens_invalidos WHERE token_invalido = $1"

	var count int

	err := db.QueryRow(query, tokenString).Scan(&count)

	if err != nil {

		log.Println("Erro ao consultar a tabela de tokens inválidos:", err)

		return true

	}

	return count > 0

}

// Função com finalidade de validação do token para as funções do usuário.
func ValidarOToken(db *sql.DB, tokenString string) (int, error) {

	if tokenEstaNaTabelaDeTokensInvalidos(db, tokenString) {

		return 0, nil

	}

	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {

		return jwtKey, nil

	})

	if err != nil || !tkn.Valid {

		return 0, err

	}

	return claims.ID_Usuario, nil

}

// Função com finalidade de geração do token para as funções do usuário.
func GerarOToken(usuario Usuario) (string, error) {

	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &Claims{

		ID_Usuario: usuario.ID_Usuario,

		StandardClaims: jwt.StandardClaims{

			ExpiresAt: expirationTime.Unix(),

		},

	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)

	if err != nil {

		return "Erro ao gerar token", err

	}

	return tokenString, nil
}