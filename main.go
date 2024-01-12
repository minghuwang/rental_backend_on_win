package main

import (
	"fmt"
	"net/http"

	"rental/property/mydatabase"
	mydb "rental/property/mydatabase"

	"github.com/labstack/echo/v4"
)

func main() {

	mydb.InitDB()
	defer mydb.CloseDB()

	// Create a new Echo instance
	e := echo.New()
	e.GET("/", func(c echo.Context) error {

		return c.JSON(http.StatusOK, fmt.Sprintf("client: %v", c))
	})
	// Define a route handler
	e.GET("/properties", func(c echo.Context) error {

		pI, err := mydb.GetProperties(mydb.Db, mydb.TbName)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, pI)
		// return c.String(http.StatusOK, fmt.Sprintf("%v", pI))
	})
	e.POST("/property", func(c echo.Context) error {

		var pI = &mydatabase.PropertyInfoSchema{}
		if err := c.Bind(pI); err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		pI, err := mydb.InsertProperty(mydb.Db, pI)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, pI)
		// return c.String(http.StatusOK, fmt.Sprintf("%v", pI))
	})

	// Start the server
	e.Logger.Fatal(e.Start(":8088")) // Listen on port
}
