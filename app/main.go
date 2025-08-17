package main

import (
	"app-demo/pkg"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()
	e.Static("/public", "public")
	e.GET("/", func(c echo.Context) error {
		return c.File("public/index.html")
		/* return c.String(http.StatusOK, "Welcome to the Movie API - by Renzo Medina") */
	})
	e.GET("/movies", func(c echo.Context) error {
		return c.JSON(200, pkg.GetMovies())
	})
	e.POST("/movies", func(c echo.Context) error {
		var movie pkg.Movies
		if err := c.Bind(&movie); err != nil {
			return c.JSON(400, map[string]string{"error": "Invalid input"})
		}
		if err := pkg.AddingMovies(movie); err != nil {
			return c.JSON(500, map[string]string{"error": "Failed to add movie"})
		}
		return c.JSON(201, movie)
	})
	e.Logger.Fatal(e.Start(":8080"))
}
