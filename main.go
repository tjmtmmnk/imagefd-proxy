package main

import (
	"github.com/labstack/echo/middleware"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo"
)

func getImage(c echo.Context) error {
	imageUrl := c.QueryParam("url")
	response, err := http.Get(imageUrl)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	return c.Blob(http.StatusOK, "image/png", body)
}
func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	e.GET("/image", getImage)
	e.Logger.Fatal(e.Start(":8080"))
}
