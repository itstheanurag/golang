package main

import (
	"fmt"

	"github.com/itstheanurag/golang/short_projects/go-fiber-crm-basic/database/database"

	"github.com/jinzhu/gorm"

	"github.com/gofiber/fiber"
)

func setupRoutes(app *fiber.App) {
	app.Get(GetLeads)
	app.Get(GetLead)
	app.Post(NewLeads)
	app.Delete(DeleteLeads)
}

func initDatabase() {
	var err error
	database.DB_CONN, err := gorm.Open("sqlite3", "leads.db")

	if err != nil {
		panic("failed to connect to the database")
	}

	fmt.Println("Connection opened to the database")
	database.DB_CONN.AutoMigrate(&lead.Lead{})
}

func main() {
	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	app.Listen(3000)
	defer database.DB_CONN.Close()
}
