# bivrost-example

This repository contain example for (simple) bootstrap on KoinWork existing stack. Our stack consist of :
- [Bivrost (API) Gateway](https://github.com/koinwork/asgard-bivrost) 
- [Heimdal Utility Library](https://github.com/koinworks/asgard-heimdal)
- [Hawkeye Cron Manager](https://github.com/koinworks/asgard-hawkeye)

These libraries (at least two of them, Bivrost & Heimdal) are main dependencies if you want to boot a service. 

You also need to read [this doc](https://koinworks.atlassian.net/wiki/spaces/~433672050/pages/1128301127/Platform+-+Introducing+Bivrost) on creating service with Bivrost

## Observe The Code

Before you run this example, please observe the sample code. There are several steps you need to do to boot service using Bivrost x Heimdal

1) Add Service Config Identifier & Redis Config. Services in KoinWorks Asgard registering themselves to Redis as Service Registry. Service List on Redis will be put under `asgard-bivrost` key and used by Bivrost to find each other service
   ```golang
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
   ```
   Example of Redis Content with this example service
   ```shell
   localhost:6379> HGETALL asgard-bivrost
   1) "asgard-example-service"
   2) "{\"host\":\"fenrir\",\"port\":9999,\"key\":\"asgard-example-service\",\"name\":\"Asgard_Example_Service\",\"namespace\":\"default\",\"gateway_endpoint\":\"/ping\"}"
   ```
2) Instantiate the server
   ```golang
   server, err := libs.NewServer(registry)
	if err != nil {
		log.Fatal(err)
   }
   ```
3) Create API endpoint. On Bivrost, endpoint receive `Context` and return the message on `Result`. You can find the definition on `github.com/koinworks/asgard-bivrost/service/context.go`
```golang
...
bivrostSvc := server.AsGatewayService(
		"/ping",
	)

bivrostSvc.Get("/", pingHandler)
...
func pingHandler(ctx *bv.Context) bv.Result {

	return ctx.JSONResponse(http.StatusOK, bv.ResponseBody{
		Message: map[string]string{
			"en": "Welcome to Ping API",
			"id": "Selamat datang di Ping API",
		},
	})

}
...
```
4. Start the server
```golang
err = server.Start()
	if err != nil {
		panic(err)
}
```
## Run This Sample

1) You need to adjust `.env` file. Watch `.env.example` for the example. Most (or all of them) of our service using [godotenv](github.com/joho/godotenv) to read environment variables. If you observer, the content of `.env`, mirror the state of your `serviceConfig`
2) Adjust the Dockerfile. You might want to change the `.netrc` line, adjust it to your Github Credential (User/Access Token). Make sure you already get Github credential and added to the all respective repositories needed.
3) Adjust Docker Compose and Build Docker Image
   - checkout Bivrost repository
   - adjust docker-compose.yml
      - adjust bivrost build path according to your local path
        ```yaml
        asgard-bivrost:
          build:
            context: $HOME/Work/KoinWorks/github/asgard-bivrost
            dockerfile: $HOME/Work/KoinWorks/github/asgard-bivrost/Dockerfile
          image: asgard-bivrost
        ...  
        ```
4) Run docker-compose `docker-compose up -d` . Check also the logs via `docker-compose logs -f`
5) Test your connection 
   ```shell
   $ curl 10.5.0.3:9000/ping
   {"status":200,"message":{"en":"Welcome to Ping API","id":"Selamat datang di Ping API"}}
   ```
