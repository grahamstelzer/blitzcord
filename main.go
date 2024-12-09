package main

import (
    "log"

    "github.com/gofiber/fiber/v3"
    "github.com/gofiber/fiber/v3/middleware/cors"
)

func main() {
    // Initialize a new Fiber app
    app := fiber.New()

    // backend run on 8000, frontend runs on different port
    // why cors?
    app.Use(cors.New()) 


    // Define a route for the GET method on the root path '/'
    app.Get("/", func(c fiber.Ctx) error {
        // Send a string response to the client
        return c.SendString("Hello, World 👋!")
    })

    log.Fatal(app.Listen(":8000"))
}