package main

import (
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
)

func setupEcho() *echo.Echo {
	e := echo.New()
	e.Debug = true
	e.Logger.SetOutput(os.Stderr)

	return e
}

func GetOmikuji(c echo.Context) error {
	rand.Seed(time.Now().UnixNano())
	omikujis := []string{
		"大吉",
		"中吉",
		"小吉",
		"吉",
		"凶",
		"大凶",
	}
	return c.JSON(http.StatusOK, omikujis[rand.Intn(len(omikujis))])
}

func main() {
	e := setupEcho()

	e.GET("omikuji", GetOmikuji)
	e.Logger.Fatal(e.Start(":1323"))
}
