package main

import (
	"github.com/gofiber/fiber/v2"
	"os"
	"log"
	"fmt"
	"net/http"
	// "github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/puneethx/shopping-cart-server/storage"
	"database/sql"
	// "github.com/joho/godotenv"
)

type Repository struct {
	DB *sql.DB
}

type Shelf struct {
	CameraID	int	`json:"cameraid"`
	RightShelf	[][]float64 `json:"rightshelf"`
	LeftShelf	[][]float64	`json:"leftshelf"`
}

type Catelog struct{
	Name string `json:"name"`
	Id int 		`json:"id"`
	Cost int 	`json:"cost"`
}

type Cart struct{
	CartID int			`json:"cartid"`
	Bottle int			`json:"bottle"`
	CellPhone int		`json:"cellphone"`
	Mouse int			`json:"mouse"`
	TotalCost int		`json:"totalcost"`
	Paid bool			`json:"paid"`
}

type Pay struct{
	CartID int 			`json:"cartid"`
}

func status(c *fiber.Ctx) error {
	return c.SendString("Server is running! Send your request")
}

func(r *Repository) CustomerCart(context *fiber.Ctx) error{
	customCart := Cart{}

	if err := context.BodyParser(&customCart); err != nil {
		return err
	}
	query := fmt.Sprintf("INSERT INTO cart (cartid,Bottle,cellphone,mouse,TotalCost,Paid) VALUES (%d,%d,%d,%d,%d,%t);", customCart.CartID, customCart.Bottle, customCart.CellPhone, customCart.Mouse, customCart.TotalCost, customCart.Paid)
	var output string
	err := r.DB.QueryRow(query).Scan(&output)
	if err != nil {
		return context.JSON(fiber.Map{"message": "Account already exits", "output": err})
	}

	return context.JSON(fiber.Map{"message": output})
}

func(r *Repository) CustomerPayment(context *fiber.Ctx) error{
	customPay := Pay{}

	if err := context.BodyParser(&customPay); err != nil {
		return err
	}
	query := fmt.Sprintf("UPDATE cart SET paid = true WHERE cartId = %d;", customPay.CartID)
	var output string
	err := r.DB.QueryRow(query).Scan(&output)
	if err != nil {
		return context.JSON(fiber.Map{"message": "Account already exits", "output": err})
	}
	return context.JSON(fiber.Map{"message": output})
}

func(r *Repository) MarkShelf(context *fiber.Ctx) error{

	camShelf := Shelf{}

	if err := context.BodyParser(&camShelf); err != nil {
		return err
	}

	// query

	return context.JSON(fiber.Map{"message": "shelf created"})
}

func(r *Repository) GetCatelog(context *fiber.Ctx) error{
	productModels := []Catelog{} 

	rows, err := r.DB.Query("select * from catelog")
	if err != nil{
		log.Fatal(err)
	}
	for rows.Next(){
		var t Catelog
		err := rows.Scan(&t.Name, &t.Id, &t.Cost)
		if err != nil{
			log.Fatal(err)
		}
		productModels = append(productModels, t)
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "products fetched successfully",
		"data": productModels,
	})

	return nil
}
func(r *Repository) SetUpRoutes(app *fiber.App){

	app.Get("/", status)

	app.Post("/customer_cart", r.CustomerCart)
	app.Get("/customer_payment", r.CustomerPayment)
	app.Get("/mark_shelf", r.MarkShelf)
	app.Get("/get_catelog", r.GetCatelog)
	
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
	// app.Use(cors.New(cors.Config{
    //     AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
    //     AllowOrigins:     "*",
    //     AllowCredentials: true,
    //     AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
    // }))

	r.SetUpRoutes(app)



	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}
	
	log.Fatal(app.Listen("0.0.0.0:" + port))
}
