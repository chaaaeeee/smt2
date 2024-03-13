package main

import (
	"github.com/gofiber/fiber/v2"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
)

var db *sql.DB
var err error

func main() {
	app := fiber.New()
	
	setupRoutes(app)
	
	fmt.Println("alr started")
	app.Listen(":9000")
}
