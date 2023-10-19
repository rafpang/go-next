package users

import (
	"backend/database"
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type UserInput struct {
	Username string
	Password string
}

func CreateUser(c *fiber.Ctx) error {
	var userInput UserInput

	if err := c.BodyParser(&userInput); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	hashedPassword, err := hashPassword(userInput.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to hash the password",
		})
	}

	newUser := database.User{
		Username:       userInput.Username,
		HashedPassword: hashedPassword,
		Base: database.Base{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	database.AppDatabase.Create(&newUser)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"user":      newUser.ID,
		"createdAt": newUser.CreatedAt,
	})
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// func verifyPassword(hashedPassword, plainPassword string) bool {
// 	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
// 	return err == nil
// }
