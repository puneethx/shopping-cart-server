package main

import (
	"github.com/gofiber/fiber/v2"
	"os"
	"log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/puneethx/shopping-cart-server/storage"
	"database/sql"
	// "github.com/joho/godotenv"
)

type Repository struct {
	DB *sql.DB
}

func status(c *fiber.Ctx) error {
	return c.SendString("Server is running! Send your request")
}
func(r *Repository) CustomerCart(context *fiber.Ctx) error{

	return context.JSON(fiber.Map{"message": "customer cart details"})
}

func(r *Repository) CustomerPayment(context *fiber.Ctx) error{

	return context.JSON(fiber.Map{"message": "customer payment details"})
}
func(r *Repository) MarkShelf(context *fiber.Ctx) error{

	return context.JSON(fiber.Map{"message": "market shelf"})
}

func(r *Repository) SetUpRoutes(app *fiber.App){

	app.Get("/", status)

	app.Get("/customer_cart", r.CustomerCart)
	app.Get("/customer_payment", r.CustomerPayment)
	app.Get("/mark_shelf", r.MarkShelf)
	
}

func main() {
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	config := &storage.Config{
		Host: os.Getenv("DB_HOST"),
		Password: os.Getenv("DB_PASS"),
		User: os.Getenv("DB_USER"),
		DBName: os.Getenv("DB_NAME"),
		SSLMode: os.Getenv("DB_SSLMODE"),
	}

	db, err := storage.NewConnection(config)

	if err != nil {
		log.Fatal("could not connect")
	}

	r := Repository{
		DB: db,
	}

	app := fiber.New()
	app.Use(cors.New(cors.Config{
        AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
        AllowOrigins:     "*",
        AllowCredentials: true,
        AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
    }))
	
	r.SetUpRoutes(app)



	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}
	
	log.Fatal(app.Listen("0.0.0.0:" + port))
}
