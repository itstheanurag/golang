package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/itstheanurag/golang/short_projects/go-fiber-crm-basic/database"
	"github.com/itstheanurag/golang/short_projects/go-fiber-crm-basic/lead"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("/api/v1/lead", lead.NewLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)
}

func initDatabase() {
	var err error

	database.DB_CONN, err = gorm.Open(sqlite.Open("leads.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to the database")
	}

	fmt.Println("Connection opened to the database")
	database.DB_CONN.AutoMigrate(&lead.Lead{})
	fmt.Println("Database Migrated")
}

func main() {
	app := fiber.New()

	initDatabase()
	sqlDB, _ := database.DB_CONN.DB()
	defer sqlDB.Close()

	setupRoutes(app)

	fmt.Println("Server starting on port 3000...")
	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
