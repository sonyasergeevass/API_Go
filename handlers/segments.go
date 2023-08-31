package handlers

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/sonyasergeevass/API_Go/database"
	"github.com/sonyasergeevass/API_Go/models"
)

func GetSegm(c *fiber.Ctx) error {
	segm := []models.Segment{}
	database.DB.Db.Find(&segm)

	return c.Status(200).JSON(segm)
}

func CreateSegment(c *fiber.Ctx) error {
	segm := new(models.Segment)
	body := c.Body()
	err := json.Unmarshal(body, segm)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&segm)

	return c.Status(200).JSON(segm)
}
