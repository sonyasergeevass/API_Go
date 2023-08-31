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

func DeleteSegment(c *fiber.Ctx) error {
	body := c.Body()
	var request struct {
		SegmentName string `json:"segment_name"`
	}
	if err := json.Unmarshal(body, &request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request format",
		})
	}
	database.DB.Db.Where("segment_name = ?", request.SegmentName).Delete(&models.Segment{})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Segment and related user associations deleted successfully",
	})
}

func GetUser(c *fiber.Ctx) error {
	user := []models.User{}
	database.DB.Db.Find(&user)

	return c.Status(200).JSON(user)
}

func CreateUser(c *fiber.Ctx) error {
	user := new(models.User)
	body := c.Body()
	err := json.Unmarshal(body, user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&user)

	return c.Status(200).JSON(user)
}

func CreateUserSegment(c *fiber.Ctx) error {
	userSegment := new(models.UserSegment)
	body := c.Body()
	err := json.Unmarshal(body, userSegment)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&userSegment)

	return c.Status(200).JSON(userSegment)
}

func GetUserSegments(c *fiber.Ctx) error {
	userSegments := []models.UserSegment{}
	database.DB.Db.Find(&userSegments)

	return c.Status(200).JSON(userSegments)
}
