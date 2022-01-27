package handler

import (
	"database/sql"
	"fmt"

	"github.com/Komunitas-Mea/marketplace-mea-api/entity"
	"github.com/Komunitas-Mea/marketplace-mea-api/usecase"
	"github.com/gofiber/fiber/v2"
)

type userHandler struct {
	userUseCase usecase.UserUseCase
}

type UserHandler interface {
	// reader
	Find(c *fiber.Ctx) error
	FindAll(c *fiber.Ctx) error

	// writer
	Create(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

func NewUserHandler(userUseCase usecase.UserUseCase) UserHandler {
	return &userHandler{
		userUseCase: userUseCase,
	}
}

func (d userHandler) Find(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")

	data, err := d.userUseCase.Find(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusNotFound).JSON(fiber.ErrNotFound)
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.ErrBadGateway)
	}

	return c.JSON(fiber.Map{
		"code": fiber.StatusOK,
		"data": data,
	})
}

func (d userHandler) FindAll(c *fiber.Ctx) error {
	users, err := d.userUseCase.FindAll()
	if err != nil {
		fmt.Println("err")
		return c.Status(fiber.StatusBadGateway).JSON(fiber.StatusBadGateway)
	}

	return c.JSON(fiber.Map{
		"code": fiber.StatusOK,
		"data": users,
	})
}

func (d userHandler) Create(c *fiber.Ctx) error {
	user := entity.User{}

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.ErrBadRequest)
	}

	err := d.userUseCase.Create(user)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadGateway).JSON(fiber.ErrBadGateway)
	}

	return c.JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"message": "OK.",
	})
}

func (d userHandler) Update(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("user_id")
	user := entity.User{}

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.ErrBadRequest)
	}

	user.ID = id
	err := d.userUseCase.Update(user)
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.ErrBadGateway)
	}

	return c.JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"message": "OK.",
	})
}

func (d userHandler) Delete(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("user_id")
	fmt.Println("id", id)

	err := d.userUseCase.Delete(id)
	if err != nil {
		fmt.Println("err", err)
		return c.Status(fiber.StatusBadGateway).JSON(fiber.ErrBadGateway)
	}

	return c.JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"message": "OK.",
	})
}
