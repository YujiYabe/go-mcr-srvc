package main
import (
    // "net/http"
    // "github.com/labstack/echo"
    // "os"
    // "fmt"
    // "go_api2/cmd/http"
    "go_api2/infrastructure"
)

func main() {
    // e := echo.New()
    // e.GET("/", func(c echo.Context) error {
    //     return c.String(http.StatusOK, "Hello, World!")
    // })
    // port := ":" + os.Getenv("PORT")
    // fmt.Println(port)
    // e.Logger.Fatal(e.Start(port))
    http.Run()
}
