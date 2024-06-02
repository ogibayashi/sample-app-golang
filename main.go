package main

import (
	"context"
	"flag"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"

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
const shutdownTimeoutSec = 5

func main() {
	configName := flag.String("config", "config", "name of config file")
	flag.Parse()

	rand.Seed(time.Now().UnixNano())
	config.Init(*configName)

	go func() {
		log.Printf("running pprof server at %s", pprofBindAddr)
		err := http.ListenAndServe(pprofBindAddr, nil)
		if err != nil {
			log.Printf("failed to run pprof: %v", err)
		}
	}()

	r := gin.Default()
	r.Use(middleware.AccessLogMiddleware())

	h, err := server.NewHandler()
	if err != nil {
		panic(err)
	}
	server.RegisterHandlers(r, server.NewStrictHandler(h, nil))

	srv := &http.Server{
		Addr:    serverPort,
		Handler: r,
	}

	go func() {
		// サービスの接続
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeoutSec*time.Second)
	defer cancel()
	if err := h.Close(); err != nil {
		log.Printf("error in closing handler: %v\n", err)
	}
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown error:%v\n", err)
	}
	log.Println("Server exiting")
}
