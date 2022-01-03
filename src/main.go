package main

import (
	"context"
	_ "elentari/docs"
	authUseCase "elentari/src/application/usecase/auth"
	catalogUseCase "elentari/src/application/usecase/catalog"
	categoryUseCase "elentari/src/application/usecase/category"
	homeUseCase "elentari/src/application/usecase/home"
	notificationUseCase "elentari/src/application/usecase/notification"
	postUseCase "elentari/src/application/usecase/post"
	serviceUseCase "elentari/src/application/usecase/service"
	categoryRepository "elentari/src/infrastructure/graphql/repository/iluvatar/category"
	notificationRepository "elentari/src/infrastructure/graphql/repository/iluvatar/notification"
	postRepository "elentari/src/infrastructure/graphql/repository/iluvatar/post"
	serviceRepository "elentari/src/infrastructure/graphql/repository/iluvatar/service"
	"elentari/src/infrastructure/http/handlers/rest/health"
	customMiddlewares "elentari/src/infrastructure/http/handlers/rest/middleware"
	authHandler "elentari/src/infrastructure/http/handlers/rest/v1/auth"
	catalogHandler "elentari/src/infrastructure/http/handlers/rest/v1/catalog"
	categoryHandler "elentari/src/infrastructure/http/handlers/rest/v1/category"
	homeHandler "elentari/src/infrastructure/http/handlers/rest/v1/home"
	notificationHandler "elentari/src/infrastructure/http/handlers/rest/v1/notification"
	postHandler "elentari/src/infrastructure/http/handlers/rest/v1/post"
	serviceHandler "elentari/src/infrastructure/http/handlers/rest/v1/service"
	authRepository "elentari/src/infrastructure/http/repository/iluvatar/auth"
	"elentari/src/shared/validations"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/pzentenoe/graphql-client"
	client "github.com/pzentenoe/httpclient-call-go"
	echoSwagger "github.com/swaggo/echo-swagger"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// @title Documentación Artefacto BFF
// @version 1.0
// @description En esta documentación se encuentran los detalles de los endpoints presentes en el artefacto BFF del proyecto Kümelen
// @contact.name Diego Sepúlveda
// @contact.email diego.sepulvedas@utem.cl
func main() {
	e := echo.New()
	e.Validator = validations.NewCustomValidator(validator.New())
	e.HideBanner = true
	e.Use(middleware.CORS())
	e.Use(middleware.Recover())

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	health.NewHealthHandler(e)

	graphQLClient := graphql.NewClient(os.Getenv("ILUVATAR_CMS_HOST") + "/graphql")

	serviceRepositoryImplementation := serviceRepository.NewServiceIluvatarRepository(graphQLClient)
	homeUseCaseImplementation := homeUseCase.NewHomeUseCase(serviceRepositoryImplementation)

	postRepositoryImplementation := postRepository.NewPostIluvatarRepository(graphQLClient)
	postPageUseCaseImplementation := postUseCase.NewPostUseCase(postRepositoryImplementation)

	serviceUseCaseImplementation := serviceUseCase.NewServiceUseCase(serviceRepositoryImplementation)

	categoryRepositoryImplementation := categoryRepository.NewCategoryIluvatarRepository(graphQLClient)
	catalogUseCaseImplementation := catalogUseCase.NewCatalogUseCase(serviceRepositoryImplementation, categoryRepositoryImplementation)

	categoryUseCaseImplementation := categoryUseCase.NewCategoryUseCase(categoryRepositoryImplementation)

	iluvatarHTTPClient := client.NewHTTPClientCall(os.Getenv("ILUVATAR_HOST"), &http.Client{})
	authRepositoryImplementation := authRepository.NewAuthIluvatarRepository(iluvatarHTTPClient)
	authUsecaseImplementation := authUseCase.NewAuthUseCase(authRepositoryImplementation)

	notificationRepositoryImplementation := notificationRepository.NewNotificationIluvatarRepository(graphQLClient)
	notificationUseCaseImplementation := notificationUseCase.NewNotificationUseCase(notificationRepositoryImplementation)

	jwtMiddleware := customMiddlewares.NewJWTMiddleware(authUsecaseImplementation)

	homeHandler.NewHomeHandler(e, homeUseCaseImplementation, jwtMiddleware)
	postHandler.NewPostHandler(e, postPageUseCaseImplementation, jwtMiddleware)
	serviceHandler.NewServiceHandler(e, serviceUseCaseImplementation, jwtMiddleware)
	catalogHandler.NewCatalogHandler(e, catalogUseCaseImplementation, jwtMiddleware)
	categoryHandler.NewCategoryHandler(e, categoryUseCaseImplementation, jwtMiddleware)
	notificationHandler.NewNotificationHandler(e, notificationUseCaseImplementation, jwtMiddleware)
	authHandler.NewAuthHandler(e, authUsecaseImplementation)

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
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
