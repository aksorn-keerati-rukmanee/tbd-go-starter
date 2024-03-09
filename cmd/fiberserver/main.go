package main

import (
	"fmt"

	"github.com/aksorn-keerati-rukmanee/tbd-go-starter/configs/database"
	"github.com/aksorn-keerati-rukmanee/tbd-go-starter/configs/environment"
	"github.com/aksorn-keerati-rukmanee/tbd-go-starter/internal/adapters/http"
	"github.com/aksorn-keerati-rukmanee/tbd-go-starter/internal/adapters/repo"
	"github.com/aksorn-keerati-rukmanee/tbd-go-starter/internal/core/entities"
	"github.com/aksorn-keerati-rukmanee/tbd-go-starter/internal/core/usecases"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func main() {

	//env
	environment.InitViper()

	//database
	db := database.InitGorm()
	err := db.AutoMigrate(
		&entities.Order{},
		&entities.User{},
		&entities.Language{},
	)
	if err != nil {
		panic(fmt.Errorf("failed to migrate database: %v", err))
	}

	//combine interface of Order
	orderRepo := repo.NewGormOrderRepository(db)
	orderService := usecases.NewOrderService(orderRepo)
	orderHandler := http.NewHttpOrderHandler(orderService)

	//combine interface of User
	userRepo := repo.NewGormUserRepository(db)
	userService := usecases.NewUserService(userRepo)
	userHandler := http.NewHttpUserHandler(userService)

	//combine interface of Language
	languageRepo := repo.NewGormLanguageRepository(db)
	languageService := usecases.NewLanguageService(languageRepo)
	languageHandler := http.NewHttpLanguageHandler(languageService)

	//using framework of http server
	{
		// instance of framework
		app := fiber.New()

		// path route of api service
		app.Get("/", func(c *fiber.Ctx) error {
			return c.SendString("nothing here")
		})

		//order router
		app.Post("/order", orderHandler.CreateOrder)

		//user router
		app.Get("/user", userHandler.GetAllUser)
		app.Post("/user", userHandler.CreateUser)

		//language router
		app.Get("/language", languageHandler.GetAllLanguage)
		app.Post("/language", languageHandler.CreateLanguage)

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
