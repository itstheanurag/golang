package lead

import (
	"github.com/gofiber/fiber/v2"
	"github.com/itstheanurag/golang/short_projects/go-fiber-crm-basic/database"
	"gorm.io/gorm"
)

type Lead struct {
	gorm.Model
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Phone   int    `json:"phone"`
}

func NewLead(c *fiber.Ctx) error {
	db := database.DB_CONN
	lead := new(Lead)

	if err := c.BodyParser(lead); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Create(lead)
	return c.JSON(lead)
}

func GetLead(c *fiber.Ctx) error {
	id := c.Params("id")

	db := database.DB_CONN
	var lead Lead

	if err := db.First(&lead, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(404).SendString("No lead found with ID")
		}
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(lead)
}

func GetLeads(c *fiber.Ctx) error {
	db := database.DB_CONN
	var leads []Lead

	db.Find(&leads)
	return c.JSON(leads)
}

func DeleteLead(c *fiber.Ctx) error {
	id := c.Params("id")

	db := database.DB_CONN
	var lead Lead

	if err := db.First(&lead, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(404).SendString("No lead found with ID")
		}
		return c.Status(500).SendString(err.Error())
	}
	
	db.Delete(&lead)
	return c.JSON(fiber.Map{"message": "Lead successfully deleted"})
}
