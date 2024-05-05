package server

import (
	"context"
	"math/rand"
	"sort"
)

const randMax = 10000

type SampleAppServer struct {
}

func (s *SampleAppServer) GetHello(ctx context.Context, request GetHelloRequestObject) (GetHelloResponseObject, error) {
	r := "Hello World!"
	return GetHello200JSONResponse{&r}, nil
}

func (s *SampleAppServer) GetSort(ctx context.Context, request GetSortRequestObject) (GetSortResponseObject, error) {
	arr := make([]int, request.Params.Size)

	for i := 0; i < request.Params.Size; i++ {
		arr[i] = rand.Intn(randMax)
	}
	sort.Ints(arr)
	r := "OK"
	return GetSort200JSONResponse{&r}, nil
}

func NewHandler() *SampleAppServer {
	return &SampleAppServer{}
}
