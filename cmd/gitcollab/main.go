// Test main.go file to create docker image
package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/GitCollabCode/GitCollab/internal/db"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "Hello, Docker! <3")
	})

	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
	})

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	authDB, err := db.ConnectPostgres(e.Logger)
	authDB.Connection.Close(context.Background())
	if err != nil {
		fmt.Println("we ball")
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}
