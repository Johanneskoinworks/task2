package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/koinworks/asgard-bivrost/libs"
	bv "github.com/koinworks/asgard-bivrost/service"
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

	bivrostSvc.Get("/", pingHandler)

	err = server.Start()
	if err != nil {
		panic(err)
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
