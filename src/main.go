package main

import (
	"context"
	homeUseCase "elentari/src/application/usecase/home"
	serviceRepository "elentari/src/infrastructure/graphql/repository/iluvatar/service"
	"elentari/src/infrastructure/http/handlers/rest/health"
	homeHandler "elentari/src/infrastructure/http/handlers/rest/v1/home"
	"elentari/src/shared/validations"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"github.com/pzentenoe/graphql-client"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	e := echo.New()
	e.Validator = validations.NewCustomValidator(validator.New())
	e.HideBanner = true
	e.Use(middleware.CORS())
	e.Use(middleware.Recover())

	health.NewHealthHandler(e)

	graphQLClient := graphql.NewClient(os.Getenv("ILUVATAR_URL") + "/graphql")

	serviceRepositoryImplementation := serviceRepository.NewServiceIluvatarRepository(graphQLClient)
	homeUseCaseImplementation := homeUseCase.NewHomeUseCase(serviceRepositoryImplementation)
	homeHandler.NewHomeHandler(e, homeUseCaseImplementation)


	quit := make(chan os.Signal, 1)
	go startServer(e, quit)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	gracefulShutdown(e)
}

func startServer(e *echo.Echo, quit chan os.Signal) {
	log.Print("Starting server")
	if err := e.Start(":" + os.Getenv("PORT")); err != nil {
		log.Error(err.Error())
		close(quit)
	}
}

func gracefulShutdown(e *echo.Echo) {
	log.Print("Shutting down server")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 10)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
