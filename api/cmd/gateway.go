package main

import (
	"os"

	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
	"github.com/guilherme-or/vivo-synq/api/internal/database"
	"github.com/guilherme-or/vivo-synq/api/internal/handler"
	"github.com/guilherme-or/vivo-synq/api/internal/repository"
	"github.com/joho/godotenv"
)

func main() {
	// Carregar as variáveis de ambiente
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}

	// Carregar as conexões com o banco de dados
	mongoDBConn, redisConn := databaseConnection()
	defer mongoDBConn.Close()
	defer redisConn.Close()

	// Instanciar os repositórios de persistência
	productRepository := repository.NewMixedProductRepository(mongoDBConn, redisConn)

	// Instanciar os handlers
	productHandler := handler.NewProductHandler(productRepository)

	// Criação da aplicação Fiber com configurações customizadas
	app := fiber.New(fiber.Config{
		Prefork:       false,
		CaseSensitive: true,
		StrictRouting: false,
		ServerHeader:  "VivoSynq",
		AppName:       "VivoSynq Main API",
		GETOnly:       true,
		
	})

	// Definição da rota principal
	app.Get("/users/:user_id/products", productHandler.FindUserProducts)

	// Inicialização do servidor
	app.Listen(os.Getenv("SERVER_ADDR"))
}

func databaseConnection() (*database.MongoDBConn, *database.RedisConn) {
	mongoDBConn := database.NewMongoDBConn(os.Getenv("MONGO_URI"))
	if err := mongoDBConn.Open(); err != nil {
		panic(err)
	}

	opt, err := redis.ParseURL(os.Getenv("REDIS_URI"))
	if err != nil {
		panic(err)
	}

	redisConn := database.NewRedisConn(opt)
	if err := redisConn.Open(); err != nil {
		panic(err)
	}

	return mongoDBConn.(*database.MongoDBConn), redisConn.(*database.RedisConn)
}
