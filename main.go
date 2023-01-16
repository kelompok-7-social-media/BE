package main

import (
	"api/config"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	cfg := config.InitConfig()
	db := config.InitDB(*cfg)
	config.Migrate(db)

	// userData := data.New(db)
	// userSrv := services.New(userData)
	// userHdl := handler.New(userSrv)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, error=${error}\n",
	}))

	// e.POST("/register", userHdl.Register())
	// e.POST("/login", userHdl.Login())
	// e.GET("/users", userHdl.Profile(), middleware.JWT([]byte(config.JWT_KEY)))
	// e.PUT("/users", userHdl.Update(), middleware.JWT([]byte(config.JWT_KEY)))
	// e.DELETE("/users", userHdl.Deactive(), middleware.JWT([]byte(config.JWT_KEY)))

	if err := e.Start(":8000"); err != nil {
		log.Println(err.Error())
	}
}