package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type User struct {
	ID int `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
}

func handleQueryParameters(c *fiber.Ctx) error {
	// Get the value of the "id" query parameter
	id := c.Query("id")
	return c.SendString(fmt.Sprintf("User ID: %s", id))
}

func handleRouteParameters(c *fiber.Ctx) error {
	// Get the value of the "id" route parameter
	id := c.Params("id")
	return c.SendString(fmt.Sprintf("User ID: %s", id))
}

func handleRequestBody(c *fiber.Ctx) error {
	// Parse the JSON data from the request body into the User struct
	var user = new(User)
	err := c.BodyParser(&user)
	if err!= nil {
    return err
  }

	return c.JSON(user)
}

func setUpRoutes(app *fiber.App) {
	// Query parameters
	//testing -> testing cURL: curl --location 'localhost:3000/query-parameters?id=12'
	app.Get("/query-parameters", handleQueryParameters)

	// Route parameters
	// testing -> curl --location 'localhost:3000/route-parameters/27'
	app.Get("/route-parameters/:id", handleRouteParameters)

	// Request body
	/*
	testing:

	curl --location 'localhost:3000/request-body' \
	--header 'Content-Type: application/json' \
	--data-raw '{
			"id": 1,
			"username": "joewilson27",
			"email": "joe@wilson.com"
	}'
	*/
	app.Post("/request-body", handleRequestBody)

}



func main() {
	var port int = 3000

	app := fiber.New()

	setUpRoutes(app)

	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("started at localhost:%d", port)
	app.Listen(addr)
}