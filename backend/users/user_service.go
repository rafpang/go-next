package users

import (
	"backend/database"
	"backend/middleware"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type UserInput struct {
	Username string
	Password string
}

type UserClaims struct {
	Username string
	jwt.StandardClaims
}

func generateJWT(username string) (string, error) {
	claims := &UserClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(15 * time.Minute).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("secret"))
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

func LoginUser(c *fiber.Ctx) error {
	var userInput UserInput

	if err := c.BodyParser(&userInput); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var user database.User
	results := database.AppDatabase.First(&user, "username = ?", userInput.Username)

	if results.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "user not found",
		})
	}

	if results.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": results.Error.Error(),
		})
	}

	err := verifyPassword(user.HashedPassword, userInput.Password)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "wrong password",
		})
	}

	token, err := generateJWT(user.Username)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate JWT",
		})
	}

	session, err := middleware.Store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get session",
		})
	}

	session.Set("token", token)

	err = session.Save()

	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": token,
	})
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func verifyPassword(hashedPassword, plainPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err
}
