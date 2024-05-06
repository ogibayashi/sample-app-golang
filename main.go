package main

import (
	"log"
	"math/rand"

	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/ogibayashi/sample-app-golang/config"
	"github.com/ogibayashi/sample-app-golang/middleware"
	"github.com/ogibayashi/sample-app-golang/server"

	_ "net/http/pprof"

	_ "github.com/grafana/pyroscope-go/godeltaprof/http/pprof"
)

const serverPort = ":8080"
const pprofBindAddr = ":8081"

func main() {
	rand.Seed(time.Now().UnixNano())
	config.Init()

	go func() {
		log.Printf("running pprof server at %s", pprofBindAddr)
		err := http.ListenAndServe(pprofBindAddr, nil)
		if err != nil {
			log.Printf("failed to run pprof: %v", err)
		}
	}()

	r := gin.Default()
	r.Use(middleware.AccessLogMiddleware())

	server.RegisterHandlers(r, server.NewStrictHandler(server.NewHandler(), nil))

	err := r.Run(serverPort)
	if err != nil {
		panic(err)
	}
}
