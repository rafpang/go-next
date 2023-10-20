package projects

import (
	"backend/database"

	"github.com/gofiber/fiber/v2"
)

type ProjectInput struct {
	ProjectName string `json:"projectName"`
}

type AddUserToProjectsInput struct {
	Username string `json:"username"`
}

func CreateProject(c *fiber.Ctx) error {
	var newProjectInput ProjectInput
	if err := c.BodyParser(&newProjectInput); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	newProject := database.Project{
		ProjectName: newProjectInput.ProjectName,
	}

	database.AppDatabase.Create(&newProject)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"id":          newProject.ID,
		"projectName": newProjectInput.ProjectName,
	})

}

func AddUserToProject(c *fiber.Ctx) error {
	projectId := c.Params("id")

	var userInput AddUserToProjectsInput
	if err := c.BodyParser(&userInput); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	var projectRef database.Project
	results := database.AppDatabase.First(&projectRef, "id = ?", projectId)

	if results.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "project not found",
		})
	}

	if results.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": results.Error.Error(),
		})
	}

	var userRef database.User

	userResult := database.AppDatabase.First(&userRef, "username = ?", userInput.Username)
	if userResult.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "user not found",
		})
	}
	if userResult.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": results.Error.Error(),
		})
	}

	projectRef.Users = append(projectRef.Users, userRef)
	database.AppDatabase.Save(&projectRef)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User added to the project successfully",
	})
}
