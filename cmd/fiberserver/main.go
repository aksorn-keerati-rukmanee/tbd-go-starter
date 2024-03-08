package main

import (
	"fmt"

	"github.com/aksorn-keerati-rukmanee/tbd-go-starter/internal/adapters"
	"github.com/aksorn-keerati-rukmanee/tbd-go-starter/internal/app/entities"
	"github.com/aksorn-keerati-rukmanee/tbd-go-starter/internal/app/usecases"
	"github.com/aksorn-keerati-rukmanee/tbd-go-starter/pkg/database"
	"github.com/aksorn-keerati-rukmanee/tbd-go-starter/pkg/env"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func main() {

	//env
	env.InitViper()

	//database
	db := database.InitGorm()
	if err := db.AutoMigrate(&entities.Order{}); err != nil {
		panic(fmt.Errorf("failed to migrate database: %v", err))
	}

	//combine interface of Order
	orderRepo := adapters.NewGormOrderRepository(db)
	orderService := usecases.NewOrderService(orderRepo)
	orderHandler := adapters.NewHttpOrderHandler(orderService)

	//using framework of http server
	{
		// instance of framework
		app := fiber.New()

		// path route of api service
		app.Get("/", func(c *fiber.Ctx) error {
			return c.SendString("nothing here")
		})
		app.Post("/order", orderHandler.CreateOrder)

		//start http server
		appMode := viper.Get("app.env")
		appPort := viper.Get("app.port")
		if appMode == "prod" || appMode == "production" {
			app.Listen(fmt.Sprintf(":%v", appPort))
		} else {
			app.Listen(fmt.Sprintf("localhost:%v", appPort))
		}
	}
}
