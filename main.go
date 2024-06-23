package main

import (
	"fmt"
	"net/http"
	"package/calculator"

	"github.com/labstack/echo/v4"
)

func main() {

	numAdd := calculator.Add(1, 2)
	fmt.Println(numAdd)

	numMultiply := calculator.PublicMultiply(2, 3)
	fmt.Println(numMultiply)

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))

}
