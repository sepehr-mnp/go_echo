package main

import (
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/users/:id", func(c echo.Context) error {

		id := c.Param("id")
		return c.String(http.StatusOK, id)
	})
	e.POST("/save", save)
	e.POST("/saveIMG", saveIMG)
	//e.Logger.Fatal(e.Start(":1323"))  ///http
	e.Logger.Fatal(e.StartTLS(":8443", "cert.csr", "server.key"))

}

func save(c echo.Context) error { // running with "curl.exe -d "name=Joe Smith" -d "email=joe@labstack.com" http://localhost:1323/save"
	// Get name and email
	name := c.FormValue("name")
	email := c.FormValue("email")
	return c.String(http.StatusOK, "name:"+name+", email:"+email)
}

func saveIMG(c echo.Context) error { // running with "curl.exe -F "name=Joe Smith" -F "avatar=@/path/to/your/avatar.png" http://localhost:1323/save"
	// Get name
	name := c.FormValue("name")
	// Get avatar
	avatar, err := c.FormFile("avatar")
	if err != nil {
		return err
	}

	// Source
	src, err := avatar.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create(avatar.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.HTML(http.StatusOK, "<b>Thank you! "+name+"</b>")
}
