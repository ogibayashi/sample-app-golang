package server

import (
	"context"
	"fmt"
	"math/rand"
	"sort"

	"github.com/ogibayashi/sample-app-golang/service/kafka"
)

const randMax = 10000

type SampleAppServer struct {
	writer *kafka.Writer
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

func (s *SampleAppServer) PostKafkaPublish(ctx context.Context, request PostKafkaPublishRequestObject) (PostKafkaPublishResponseObject, error) {
	err := s.writer.Write(request.Body.Message)
	if err != nil {
		return PostKafkaPublish500JSONResponse{}, err
	}
	r := "OK"
	return PostKafkaPublish200JSONResponse{&r}, nil
}

func NewHandler() (*SampleAppServer, error) {
	w, err := kafka.NewWriter()
	if err != nil {
		return nil, fmt.Errorf("failed to create kafka writer: %w", err)
	}
	return &SampleAppServer{writer: w}, nil
}
