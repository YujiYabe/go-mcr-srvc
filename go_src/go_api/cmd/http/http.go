package http

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"go_api/cmd/http/controller"
	"go_api/internal/service2"
    "go_api/pkg/mysql"

    "github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

    // "github.com/labstack/echo"
	// "github.com/labstack/echo/v4/middleware"
)

var e = createMux()

func Run() {
	http.Handle("/", e)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func init() {
    // Mysql
    c := mysql.Config{
        Host:    os.Getenv("DB_HOST"),
        Port:    os.Getenv("DB_PORT"),
        User:    os.Getenv("DB_USER"),
        DBName:  os.Getenv("DB_NAME"),
        Passwd:  os.Getenv("DB_PASSWORD"),

        // Host:    "localhost",
        // Port:    "3306",
        // DBName:  "app",
        // User:    "root",
        // Passwd:  "password",

        AllowNativePasswords: true,
    }
    db, err := mysql.Connect(c)
    if err != nil {
        log.Fatalln(err)
    }

	ctrl := &controller.Controller{}

    // DI
    gateway2 := service2.NewGateway(db)
    provider2 := service2.NewProvider(gateway2)
    ctrl2 := controller.NewController2(provider2)

	e.GET("/:message", ctrl.HandleMessage)
    e.GET("/people/:personID", ctrl2.HandlePersonGet)
    e.POST("/people", ctrl2.HandlePersonRegister) 

}  

func createMux() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip()) 

	return e
}
