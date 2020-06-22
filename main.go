package main

import (
	"github.com/labstack/echo/middleware"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo"
)

type Data struct {
	Url string `json:"url"`
}

func getImage(c echo.Context) error {
	data := new(Data)
	err := c.Bind(data)
	if err != nil {
		return err
	}
	response, err := http.Get(data.Url)
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
	e.POST("/image", getImage)
	e.Logger.Fatal(e.Start(":8080"))
}
