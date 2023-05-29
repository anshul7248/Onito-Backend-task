package main

import (
	adapterdatabase "testTask/adapter"
	config "testTask/config"
	"testTask/testTask/controller"
	"testTask/testTask/repository"
	"testTask/testTask/usecase"

	"sync"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var onceRest sync.Once

func main() {
	onceRest.Do(func() {
		e := echo.New()

		//Setting up the config
		config := config.GetConfig()
		//Setting up the Logger
		// logger := logging.NewLogger(config.Log.LogFile, config.Log.LogLevel)
		e.Use(middleware.Logger())
		e.Use(middleware.Recover())

		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
		}))

		db := adapterdatabase.DB(config)

		Repo := repository.NewTestTaskRepo(db)
		Uc := usecase.NewTestTaskUsecase(Repo)
		controller.NewTestTaskController(e, Uc)

		// , config.HttpConfig.HostCert, config.HttpConfig.HostKey
		//"./fullchain1.pem", "./privkey1.pem"

		// if err := e.StartTLS(config.HttpConfig.HostPort, "", ""); err != nil {
		if err := e.Start(config.HttpConfig.HostPort); err != nil {
			// 	fmt.Println("Couldn't Connect")
		}
	})
}
