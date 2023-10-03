package main

import (
	"github.com/gofiber/fiber/v2"
	"os"
	"log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/puneethx/shopping-cart-server/db_store"
	"database/sql"
	// "github.com/joho/godotenv"
)

type Repository struct {
	DB *sql.DB
}

func status(c *fiber.Ctx) error {
	return c.SendString("Server is running! Send your request")
}



func SetUpRoutes(app *fiber.App) {
	app.Get("/", status)

	// app.Post("/customer_cart",)
	// app.Post("/customer_payment",)
	// app.Post("/mark_shelf",)

}

func main() {
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	config := &db_store.Config{
		Host: os.Getenv("DB_HOST"),
		Password: os.Getenv("DB_PASS"),
		User: os.Getenv("DB_USER"),
		DBName: os.Getenv("DB_NAME"),
		SSLMode: os.Getenv("DB_SSLMODE"),
	}
	if err != nil {
		log.Fatal("could not connect")
	}

	r := Repository{
		DB: db,
	}

	db, err := db_store.NewConnection(config)

	app := fiber.New()
	
	app.Use(cors.New(cors.Config{
        AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
        AllowOrigins:     "*",
        AllowCredentials: true,
        AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
    }))

	SetUpRoutes(app)

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}
	
	log.Fatal(app.Listen("0.0.0.0:" + port))
}
