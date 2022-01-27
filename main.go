package main

import (
	"fmt"

	"github.com/Komunitas-Mea/marketplace-mea-api/database"
	"github.com/Komunitas-Mea/marketplace-mea-api/route"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println(".env err :", err)
	}

	app := fiber.New(fiber.Config{
		AppName: "Mea Marketplace",
	})

	db := database.NewMysqlDB()

	route.User(db, app)

	app.Listen(":8080")
}
