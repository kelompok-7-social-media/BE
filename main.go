package main

import (
	"log"
	"project/config"
	cd "project/features/komentar/data"
	chl "project/features/komentar/handler"
	csrv "project/features/komentar/service"

	"project/features/user/data"
	"project/features/user/handler"
	"project/features/user/services"

	pd "project/features/posting/data"
	phl "project/features/posting/handler"
	psrv "project/features/posting/services"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	cfg := config.InitConfig()
	db := config.InitDB(*cfg)
	config.Migrate(db)

	userData := data.New(db)
	userSrv := services.New(userData)
	userHdl := handler.New(userSrv)

	postingData := pd.New(db)
	postingSrv := psrv.New(postingData)
	postingHdl := phl.New(postingSrv)

	commentData := cd.New(db)
	commentSrv := csrv.New(commentData)
	commentHdl := chl.New(commentSrv)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, error=${error}\n",
	}))

	e.POST("users/register", userHdl.Register())
	e.POST("users/login", userHdl.Login())
	e.GET("/users", userHdl.AllUser())
	e.GET("/users/profile", userHdl.Profile(), middleware.JWT([]byte(config.JWT_KEY)))
	e.PUT("/users", userHdl.Update(), middleware.JWT([]byte(config.JWT_KEY)))

	e.DELETE("/users", userHdl.Delete(), middleware.JWT([]byte(config.JWT_KEY)))

	e.POST("/posting", postingHdl.Add(), middleware.JWT([]byte(config.JWT_KEY)))
	e.POST("/comment", commentHdl.Add(), middleware.JWT([]byte(config.JWT_KEY)))
	if err := e.Start(":8000"); err != nil {
		log.Println(err.Error())
	}
}
