package main

import (
	"log"
	"math/rand"

	"net/http"
	"sort"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/ogibayashi/sample-app-golang/gen"
	"github.com/ogibayashi/sample-app-golang/middleware"

	_ "net/http/pprof"

	_ "github.com/grafana/pyroscope-go/godeltaprof/http/pprof"
)

const serverPort = ":8080"
const randMax = 10000
const pprofBindAddr = ":8081"

type SampleHanlder struct {
}

func (h *SampleHanlder) GetHello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, World!",
	})
}

func (h *SampleHanlder) GetSort(c *gin.Context, params gen.GetSortParams) {
	arr := make([]int, params.Size)

	for i := 0; i < params.Size; i++ {
		arr[i] = rand.Intn(randMax)
	}
	sort.Ints(arr)

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

func main() {
	rand.Seed(time.Now().UnixNano())

	go func() {
		log.Printf("running pprof server at %s", pprofBindAddr)
		err := http.ListenAndServe(pprofBindAddr, nil)
		if err != nil {
			log.Printf("failed to run pprof: %v", err)
		}
	}()

	r := gin.Default()
	r.Use(middleware.AccessLogMiddleware())

	gen.RegisterHandlers(r, &SampleHanlder{})

	err := r.Run(serverPort)
	if err != nil {
		panic(err)
	}
}
