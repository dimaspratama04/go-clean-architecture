package config

import (
	"golang-redis/internal/delivery/http/route"
	"golang-redis/internal/repository"
	"golang-redis/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/rabbitmq/amqp091-go"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	Db       *gorm.DB
	App      *fiber.App
	Rdb      *redis.Client
	RabbitMQ *amqp091.Connection
}

func Bootstrap(config *BootstrapConfig) {
	productRepository := repository.NewProductRepository(config.Db, config.Rdb)
	usersRepository := repository.NewUsersRepository(config.Db)

	// setup usecase
	guestUC := usecase.NewGuestUsecase()
	usersUC := usecase.NewUsersUseCase(usersRepository)
	productUC := usecase.NewProductUseCase(productRepository)

	// router config
	rc := route.RouteConfig{
		App:       config.App,
		UsersUC:   usersUC,
		GuestUC:   guestUC,
		ProductUC: productUC}

	rc.SetupGuestRoute()
	rc.SetupAuthRoute()
}
