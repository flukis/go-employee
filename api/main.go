package main

import (
	"employee/api/routes"
	"employee/pkg/department"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sqlx.Connect("postgres", "user=root dbname=employee password=wap12345 sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	defer func() {
		err := db.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}()
	// Creates a new Fiber instance.
	app := fiber.New()

	// Prepare a static middleware to serve the built React files.
	app.Static("/", "./web/build")

	// Prepare a fallback route to always serve the 'index.html', had there not be any matching routes.
	app.Static("*", "./web/build/index.html")

	deptRepo := department.NewPotgreRepository(db)
	deptService := department.NewService(deptRepo)

	api := app.Group("/api")
	routes.DeptRouter(api, deptService)

	// Listen to port 8080.
	log.Fatal(app.Listen(":8080"))
}
