package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"task2/controller"
	"task2/database"

	"github.com/koinworks/asgard-bivrost/libs"
	bv "github.com/koinworks/asgard-bivrost/service"
	"github.com/koinworks/asgard-heimdal/libs/serror"
	hmodels "github.com/koinworks/asgard-heimdal/models"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	database.Start(os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), "5432")
}
func main() {

	hostname, _ := os.Hostname()

	portNumber, _ := strconv.Atoi(os.Getenv("APP_PORT"))

	serviceConfig := &hmodels.Service{
		Class:     "product-service",
		Key:       os.Getenv("APP_KEY"),
		Name:      os.Getenv("APP_NAME"),
		Version:   os.Getenv("APP_VERSION"),
		Host:      hostname,
		Port:      portNumber,
		Namespace: os.Getenv("K8S_NAMESPACE"),
		Metas:     make(hmodels.ServiceMetas),
	}

	registry, err := libs.InitRegistry(libs.RegistryConfig{
		Address:  os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"),
		Service:  serviceConfig,
	})

	if err != nil {
		log.Fatal(err)
	}

	server, err := libs.NewServer(registry)
	if err != nil {
		log.Fatal(err)
	}

	bivrostSvc := server.AsGatewayService(
		"/v1",
	)

	bivrostSvc.Get("/list", bivrostSvc.WithMiddleware(controller.PingHandlerList, listMiddleware))
	bivrostSvc.Post("/list", bivrostSvc.WithMiddleware(controller.PingHandlerCreateProduct, createProductMiddleware))
	bivrostSvc.Post("/createorder", bivrostSvc.WithMiddleware(controller.PingHandlerCreateOrder, createProductMiddleware))
	bivrostSvc.Get("/orders", bivrostSvc.WithMiddleware(controller.PingHandlerOrder, OrdersListMiddleware))
	bivrostSvc.Get("/ping-error", pingHandlerWithError)

	err = server.Start()
	if err != nil {
		panic(err)
	}

}

func listMiddleware(next bv.HandlerFunc) bv.HandlerFunc {
	return func(ctx *bv.Context) bv.Result {
		log.Println("This is some middleware")
		ctx.SetHeader("X-Middleware", "Message From Middleware")
		return next(ctx)
	}
}
func createProductMiddleware(next bv.HandlerFunc) bv.HandlerFunc {
	return func(ctx *bv.Context) bv.Result {
		log.Println("This is some middleware")
		ctx.SetHeader("X-Middleware", "Message From Middleware")
		return next(ctx)
	}
}
func OrdersListMiddleware(next bv.HandlerFunc) bv.HandlerFunc {
	return func(ctx *bv.Context) bv.Result {
		log.Println("This is some middleware")
		ctx.SetHeader("X-Middleware", "Message From Middleware")
		return next(ctx)
	}
}
func createOrderMiddleware(next bv.HandlerFunc) bv.HandlerFunc {
	return func(ctx *bv.Context) bv.Result {
		log.Println("This is some middleware")
		ctx.SetHeader("X-Middleware", "Message From Middleware")
		return next(ctx)
	}
}

func pingHandlerWithError(ctx *bv.Context) bv.Result {
	err := raiseError(1)
	if err != nil {
		ctx.CaptureSErrors(serror.NewFromErrorc(err, "[asgard-service-example][bivrost] error raised on handler"))
		return ctx.JSONResponse(http.StatusServiceUnavailable, bv.ResponseBody{
			Message: map[string]string{
				"en": "Ping API raised an errssor",
				"id": "Ping API mengalami kegassgalan",
			},
		})
	}

	return ctx.JSONResponse(http.StatusOK, bv.ResponseBody{
		Message: map[string]string{
			"en": "Ping API successfully invossked",
			"id": "Ping API berhasil dipanssggil",
		},
	})
}

func raiseError(errorCode int) error {
	return fmt.Errorf("error number: %d", errorCode)
}
