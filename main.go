package main

import (
	"fmt"
	"github.com/fuxingZhang/zip"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"io"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	e := echo.New()

	e.POST("/submit/:id", func(c echo.Context) error {
		id := c.Param("id")
		dir := filepath.Join("uploads", id)
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}

		filename := randomString(10) + ".csv"
		body, err := io.ReadAll(c.Request().Body)
		if err != nil {
			return err
		}

		f := filepath.Join(dir, filename)
		err = os.WriteFile(f, body, 0644)
		if err != nil {
			return err
		}
		logrus.Infof("Received file %s", filename)
		return c.NoContent(http.StatusOK)
	})
	e.GET("/", func(c echo.Context) error {
		directories, err := os.ReadDir("uploads")
		if err != nil {
			return err
		}

		count := make(map[string]int)
		for _, dir := range directories {
			files, err := os.ReadDir(filepath.Join("uploads", dir.Name()))
			if err != nil {
				return err
			}
			count[dir.Name()] = len(files)
		}
		var out string
		for k, v := range count {
			out += k + ": " + fmt.Sprint(v) + "\n"
		}
		return c.String(http.StatusOK, out)
	})
	e.GET("/download", func(c echo.Context) error {
		err := zip.Dir("uploads", "dataset.zip", true)
		if err != nil {
			return err
		}
		return c.File("dataset.zip")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "4301"
	}
	logrus.Infof("Listening on port %s", port)
	err := e.Start(":" + port)
	if err != nil {
		return
	}
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz")

func randomString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
