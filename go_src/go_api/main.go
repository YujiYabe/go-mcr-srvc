package main
import (
    "net/http"
    "github.com/labstack/echo"
    "os"
    "fmt"

)

func main() {
    e := echo.New()
    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello, World!")
    })
    port := ":" + os.Getenv("PORTS")
    fmt.Println(port)
    e.Logger.Fatal(e.Start(port))
}
