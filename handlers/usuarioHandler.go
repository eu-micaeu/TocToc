package handlers

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type Usuario struct {
	Nickname string `json:"nickname"`
	Senha    string `json:"senha"`
}

// Estrutura do TOKEN.
type Claims struct {
	Nickname string `json:"nickname"`
	jwt.StandardClaims
}

func GerarToken(nickname string, secretKey string) (string, error) {
	claims := Claims{
		Nickname: nickname,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // Token expira em 24 horas
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (u *Usuario) Entrar(client *redis.Client, secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var usuario Usuario

		if err := c.BindJSON(&usuario); err != nil {
			c.JSON(400, gin.H{"message": "Erro ao fazer login"})
			return
		}

		val, _ := client.Get(client.Context(), "usuario:"+usuario.Nickname).Result()

		if val != usuario.Senha {
			c.JSON(401, gin.H{"message": "Usuário ou senha incorretos"})
			return
		}

		// Gere o token
		token, err := GerarToken(usuario.Nickname, secretKey)
		if err != nil {
			c.JSON(500, gin.H{"message": "Erro ao gerar token"})
			return
		}
		
		http.SetCookie(c.Writer, &http.Cookie{

			Name:     "token",

			Value:    token,

			HttpOnly: true,

			Secure:   true,

			SameSite: http.SameSiteStrictMode,

		})

		c.JSON(200, gin.H{"message": "Login efetuado com sucesso!"})
	}
}


func (u *Usuario) Cadastrar(client *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		var usuario Usuario

		if err := c.BindJSON(&usuario); err != nil {
			c.JSON(400, gin.H{"message": "Erro ao registrar usuário"})
			return
		}

		// Verifique se o usuário já existe no Redis
		existe, _ := client.Exists(client.Context(), "usuario:" + usuario.Nickname).Result()
		
		if existe != 0 {
			c.JSON(400, gin.H{"message": "Este nome de usuário já está em uso"})
			return
		}

		err := client.Set(client.Context(), "usuario:"+usuario.Nickname, usuario.Senha, 0).Err()
		if err != nil {
			c.JSON(500, gin.H{"message": "Erro ao registrar usuário"})
			return
		}

		c.JSON(200, gin.H{"message": "Registro bem-sucedido"})
	}
}


