package main

import (
	"net/http"

	"fmt"

	"github.com/labstack/echo"
)

// Resource the web resource
type Resource struct {
	// ID resource unique identifier.
	ID int
	/// FirstName
	FirstName string
	// LastName
	LastName string
	// Email
	Email string
	// Password
	Password string
}

// ToString return string representation of the resource.
func (r Resource) ToString() string {
	return fmt.Sprintf("first=%s last=%s", r.FirstName, r.LastName)
}

var resources map[string]Resource

func init() {
	resources = make(map[string]Resource)

	resources["1"] = Resource{
		ID:        1,
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@example.com",
		Password:  "very_difficult",
	}
}

// getResource handles GET /users/:id
func getResource(c echo.Context) error {
	fmt.Println("getResource()")

	// User ID from path `users/:id`
	id := c.Param("id")

	val, ok := resources[id]
	if ok {
		fmt.Printf("getResource(%s)\n", id)
		return c.String(http.StatusOK, val.ToString())
	}

	fmt.Println("getResource(): NOT FOUND")
	return c.String(http.StatusNotFound, "")
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Alive!")
	})

	// e.POST("/resources", saveResource)
	e.GET("/resources/:id", getResource)
	// e.PUT("/resources/:id", updateResource)
	// e.DELETE("/resources/:id", deleteResource)

	e.Logger.Fatal(e.Start(":1323"))
}
