package main

import (
	"math/rand"

	"net/http"
	"sort"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/ogibayashi/sample-app-golang/gen"
	"github.com/ogibayashi/sample-app-golang/middleware"
)

const serverPort = ":8180"
const randMax = 10000

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
	// Ginルーターのインスタンスを作成
	r := gin.Default()
	r.Use(middleware.AccessLogMiddleware())

	gen.RegisterHandlers(r, &SampleHanlder{})

	// サーバーを起動
	r.Run(serverPort)
}
