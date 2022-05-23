package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/koinworks/asgard-bivrost/libs"
	bv "github.com/koinworks/asgard-bivrost/service"
	"github.com/koinworks/asgard-heimdal/libs/serror"
	hmodels "github.com/koinworks/asgard-heimdal/models"

	"github.com/joho/godotenv"
)

func main() {

	hostname, _ := os.Hostname()

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

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
		"/ping",
	)

	bivrostSvc.Get("/", bivrostSvc.WithMiddleware(pingHandler, exampleMiddleware))
	bivrostSvc.Get("/ping-error", pingHandlerWithError)

	err = server.Start()
	if err != nil {
		panic(err)
	}

}

func exampleMiddleware(next bv.HandlerFunc) bv.HandlerFunc {
	return func(ctx *bv.Context) bv.Result {
		log.Println("This is some middleware")
		ctx.SetHeader("X-Middleware", "Message From Middleware")
		return next(ctx)
	}
}

func pingHandler(ctx *bv.Context) bv.Result {

	return ctx.JSONResponse(http.StatusOK, bv.ResponseBody{
		Message: map[string]string{
			"en": "Welcome to Ping API",
			"id": "Selamat datang di Ping API",
		},
	})

}

func pingHandlerWithError(ctx *bv.Context) bv.Result {
	err := raiseError(1)
	if err != nil {
		ctx.CaptureSErrors(serror.NewFromErrorc(err, "[asgard-service-example][bivrost] error raised on handler"))
		return ctx.JSONResponse(http.StatusServiceUnavailable, bv.ResponseBody{
			Message: map[string]string{
				"en": "Ping API raised an error",
				"id": "Ping API mengalami kegagalan",
			},
		})
	}

	return ctx.JSONResponse(http.StatusOK, bv.ResponseBody{
		Message: map[string]string{
			"en": "Ping API successfully invoked",
			"id": "Ping API berhasil dipanggil",
		},
	})
}

func raiseError(errorCode int) error {
	return fmt.Errorf("error number: %d", errorCode)
}
