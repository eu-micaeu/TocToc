package database

import (
	"log"

	"github.com/go-redis/redis/v8"
)

// Função para criar e conectar a um banco de dados Redis
func NewRedisClient() (*redis.Client, error) {

	// URL de conexão Redis
    redisURL := "rediss://red-cl22t48p2gis73fu5ja0:d2JNtmfp3LyPOp5gBXVlU4lQRoLsowBq@oregon-redis.render.com:6379"

    // Crie uma nova opção de conexão usando a URL
    options, err := redis.ParseURL(redisURL)
    if err != nil {
        log.Fatal("Erro ao analisar a URL de conexão:", err)
    }

    // Crie o cliente Redis
    client := redis.NewClient(options)

    // Verifique a conexão com o servidor
    _, err = client.Ping(client.Context()).Result()
    if err != nil {
        log.Fatal("Erro ao conectar ao banco Redis:", err)
    }

    log.Println("Conexão com o banco Redis estabelecida com sucesso!")

	return client, nil
}
