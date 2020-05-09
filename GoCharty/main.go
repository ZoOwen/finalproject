package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
  "github.com/zoowen/model"
  "github.com/go-sql-driver/mysql"
	"github.com/jumadimuhammad/http-request/model"
	"github.com/labstack/echo"
)

func app(e *echo.Echo, store model.UserStore) {
  e.POST("/user", func(c echo.Context) error {
		// Given
		title := c.FormValue("title")
		body := c.FormValue("body")

		// Create instance
		user, _ := model.CreateUser(title, body)

		// Persist
		store.Save(user)

		// Response
		return c.JSON(http.StatusOK, user)
	})
}

func main() {
	// Echo instance
	e := echo.New()

	var store model.UserStore
	store = model.NewUserMySQL()

	// // Middleware
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	// // Routes
	// e.GET("/", hello)
  // e.POST("/user", store)
  
  app(e, store)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
