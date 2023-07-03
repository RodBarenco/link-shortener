package server

import (
	"fmt"

	"github.com/RodBarenco/link-shortener/model"
	"github.com/RodBarenco/link-shortener/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func getAllGolies(ctx *fiber.Ctx) error {
	golies, err := model.GetAllGolies()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error getting all goly links " + err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(golies)
}

func getGoly(c *fiber.Ctx) error {
	id := c.Params("id")

	golyID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error parsing goly ID " + err.Error(),
		})
	}

	goly, err := model.GetGoly(golyID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error could not get goly from db " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(goly)
}

func createGoly(c *fiber.Ctx) error {
	c.Accepts("application/json")
	var goly model.Goly
	err := c.BodyParser(&goly)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error parsing JSON " + err.Error(),
		})
	}

	goly.ID = uuid.New()

	if goly.Random {
		goly.Goly = utils.RandomURL(8)
	}

	err = model.CreateGoly(goly)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not creat goly in db " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(goly)
}

func updateGoly(c *fiber.Ctx) error {
	var goly model.Goly
	err := c.BodyParser(&goly)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error parsing JSON " + err.Error(),
		})
	}

	err = model.UpdateGoly(goly)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error updating goly " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(goly)
}

func deleteGoly(c *fiber.Ctx) error {
	id := c.Params("id")
	golyID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error could not parse id " + err.Error(),
		})
	}

	err = model.DeleteGoly(golyID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error deleting goly " + err.Error(),
		})
	}

	return c.SendString("Goly deleted successfully")
}

func findGolyByURL(c *fiber.Ctx) error {
	url := c.Params("url")

	goly, err := model.FindByUrl(url)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error finding goly by URL " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(goly)
}

func redirectTo(c *fiber.Ctx) error {
	url := c.Params("url")

	goly, err := model.FindByUrl(url)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error finding your voyage" + err.Error(),
		})
	}
	goly.Clicked += 1
	err = model.UpdateGoly(goly)
	if err != nil {
		fmt.Printf("there was a error: %v", err)
	}

	return c.Redirect(goly.Redirect, fiber.StatusTemporaryRedirect)
}
