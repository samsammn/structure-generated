package route

import (
	"github.com/Komunitas-Mea/marketplace-mea-api/handler"
	"github.com/Komunitas-Mea/marketplace-mea-api/repository"
	"github.com/Komunitas-Mea/marketplace-mea-api/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func User(db *sqlx.DB, app *fiber.App) {
	userRepo := repository.NewUserRepository(db)
	userUseCase := usecase.NewUserUseCase(userRepo)
	userHandler := handler.NewUserHandler(userUseCase)

	app.Get("/user", userHandler.FindAll)
	app.Get("/user/:id", userHandler.Find)
	app.Post("/user", userHandler.Create)
	app.Put("/user/:user_id", userHandler.Update)
	app.Delete("/user/:user_id", userHandler.Delete)
}
