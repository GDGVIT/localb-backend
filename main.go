package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func dbConnect(host, port, user, dbname, password, sslmode string) (*gorm.DB, error) {
	// In the case of heroku
	if os.Getenv("DATABASE_URL") != "" {
		return gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
	}
	db, err := gorm.Open(
		postgres.Open(fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", host, port, user, dbname, password, sslmode)),
		&gorm.Config{},
	)

	return db, err
}

func GetPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
		fmt.Println("INFO: No PORT environment variable detected, defaulting to 4000")
		return "localhost:3000"
	}
	return ":" + port
}

func main() {
	if os.Getenv("onServer") != "True" {
		// Loading the .env file
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	// Setting up DB
	_, err := dbConnect(
		os.Getenv("dbHost"),
		os.Getenv("dbPort"),
		os.Getenv("dbUser"),
		os.Getenv("dbName"),
		os.Getenv("dbPass"),
		os.Getenv("sslmode"),
	)
	if err != nil {
		log.Fatalf("Error connecting to the database: %s", err.Error())
	}
	fmt.Println("Connected to DB...")

	app := fiber.New(fiber.Config{CaseSensitive: true})

	// Middlewares
	app.Use(logger.New())
	app.Use(cors.New())

	// Migrate Tables

	// Make repos, services and  handlers

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hey there looks like its working 🔥")
	})

	fmt.Println("Serving...")
	log.Fatal(app.Listen(GetPort()))
}
