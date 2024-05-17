// Package server provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.0 DO NOT EDIT.
package server

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gin-gonic/gin"
	"github.com/oapi-codegen/runtime"
	strictgin "github.com/oapi-codegen/runtime/strictmiddleware/gin"
)

// SystemError defines model for SystemError.
type SystemError struct {
	Error *string `json:"error,omitempty"`
}

// PostKafkaPublishJSONBody defines parameters for PostKafkaPublish.
type PostKafkaPublishJSONBody struct {
	// Message message to publish
	Message string `json:"message"`
}

// GetSortParams defines parameters for GetSort.
type GetSortParams struct {
	// Size size of the list
	Size int `form:"size" json:"size"`
}

// PostKafkaPublishJSONRequestBody defines body for PostKafkaPublish for application/json ContentType.
type PostKafkaPublishJSONRequestBody PostKafkaPublishJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// just return hello
	// (GET /hello)
	GetHello(c *gin.Context)
	// post a message to kafka
	// (POST /kafka/publish)
	PostKafkaPublish(c *gin.Context)
	// generate a list of random numbers and sort them
	// (GET /sort)
	GetSort(c *gin.Context, params GetSortParams)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// GetHello operation middleware
func (siw *ServerInterfaceWrapper) GetHello(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetHello(c)
}

// PostKafkaPublish operation middleware
func (siw *ServerInterfaceWrapper) PostKafkaPublish(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PostKafkaPublish(c)
}

// GetSort operation middleware
func (siw *ServerInterfaceWrapper) GetSort(c *gin.Context) {

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetSortParams

	// ------------- Required query parameter "size" -------------

	if paramValue := c.Query("size"); paramValue != "" {

	} else {
		siw.ErrorHandler(c, fmt.Errorf("Query argument size is required, but not found"), http.StatusBadRequest)
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "size", c.Request.URL.Query(), &params.Size)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter size: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetSort(c, params)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL      string
	Middlewares  []MiddlewareFunc
	ErrorHandler func(*gin.Context, error, int)
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router gin.IRouter, si ServerInterface) {
	RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router gin.IRouter, si ServerInterface, options GinServerOptions) {
	errorHandler := options.ErrorHandler
	if errorHandler == nil {
		errorHandler = func(c *gin.Context, err error, statusCode int) {
			c.JSON(statusCode, gin.H{"msg": err.Error()})
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandler:       errorHandler,
	}

	router.GET(options.BaseURL+"/hello", wrapper.GetHello)
	router.POST(options.BaseURL+"/kafka/publish", wrapper.PostKafkaPublish)
	router.GET(options.BaseURL+"/sort", wrapper.GetSort)
}

type SystemErrorJSONResponse struct {
	Error *string `json:"error,omitempty"`
}

type GetHelloRequestObject struct {
}

type GetHelloResponseObject interface {
	VisitGetHelloResponse(w http.ResponseWriter) error
}

type GetHello200JSONResponse struct {
	Message *string `json:"message,omitempty"`
}

func (response GetHello200JSONResponse) VisitGetHelloResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetHello500JSONResponse struct{ SystemErrorJSONResponse }

func (response GetHello500JSONResponse) VisitGetHelloResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type PostKafkaPublishRequestObject struct {
	Body *PostKafkaPublishJSONRequestBody
}

type PostKafkaPublishResponseObject interface {
	VisitPostKafkaPublishResponse(w http.ResponseWriter) error
}

type PostKafkaPublish200JSONResponse struct {
	Message *string `json:"message,omitempty"`
}

func (response PostKafkaPublish200JSONResponse) VisitPostKafkaPublishResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PostKafkaPublish500JSONResponse struct{ SystemErrorJSONResponse }

func (response PostKafkaPublish500JSONResponse) VisitPostKafkaPublishResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type GetSortRequestObject struct {
	Params GetSortParams
}

type GetSortResponseObject interface {
	VisitGetSortResponse(w http.ResponseWriter) error
}

type GetSort200JSONResponse struct {
	Message *string `json:"message,omitempty"`
}

func (response GetSort200JSONResponse) VisitGetSortResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetSort500JSONResponse struct{ SystemErrorJSONResponse }

func (response GetSort500JSONResponse) VisitGetSortResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {
	// just return hello
	// (GET /hello)
	GetHello(ctx context.Context, request GetHelloRequestObject) (GetHelloResponseObject, error)
	// post a message to kafka
	// (POST /kafka/publish)
	PostKafkaPublish(ctx context.Context, request PostKafkaPublishRequestObject) (PostKafkaPublishResponseObject, error)
	// generate a list of random numbers and sort them
	// (GET /sort)
	GetSort(ctx context.Context, request GetSortRequestObject) (GetSortResponseObject, error)
}

type StrictHandlerFunc = strictgin.StrictGinHandlerFunc
type StrictMiddlewareFunc = strictgin.StrictGinMiddlewareFunc

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
}

// GetHello operation middleware
func (sh *strictHandler) GetHello(ctx *gin.Context) {
	var request GetHelloRequestObject

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetHello(ctx, request.(GetHelloRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetHello")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(GetHelloResponseObject); ok {
		if err := validResponse.VisitGetHelloResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// PostKafkaPublish operation middleware
func (sh *strictHandler) PostKafkaPublish(ctx *gin.Context) {
	var request PostKafkaPublishRequestObject

	var body PostKafkaPublishJSONRequestBody
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.Status(http.StatusBadRequest)
		ctx.Error(err)
		return
	}
	request.Body = &body

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PostKafkaPublish(ctx, request.(PostKafkaPublishRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostKafkaPublish")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(PostKafkaPublishResponseObject); ok {
		if err := validResponse.VisitPostKafkaPublishResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// GetSort operation middleware
func (sh *strictHandler) GetSort(ctx *gin.Context, params GetSortParams) {
	var request GetSortRequestObject

	request.Params = params

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetSort(ctx, request.(GetSortRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetSort")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(GetSortResponseObject); ok {
		if err := validResponse.VisitGetSortResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/9yTQW/UMBCF/4o1cEybBcQlRyQEVQ9U2mPVgzeZzXobe9yZSUWo8t/ROMt2WyokBKfe",
	"7PHL+M33nAdoKWZKmFSgeQBGyZQEy2Y9iWL8zExs25aSYlJb+pyH0HoNlOq9ULKatDuM3laZKSNrWLrg",
	"r+/xu495QGjgIily8oMT5HtkVySO2nZkxg4q0CmbTpRD6mGejxXa7LFVmK3UobQcspk47bleei62izCk",
	"LZkBDVqul+LjzOd81tPgUw8V3CPL0ujd+ep8BXMFlDH5HKCBD6VUQfa6KzPVOxyG0rPHAsQGLjguOmjg",
	"C+rXIqieAn2/Wv0DyIgivscyyd/z+XZpM31cHLxl3EIDb+rH9Ouj0/o0d2skY4yeJ2hgP4o6Rh05uQWB",
	"nde3fnvr6zxuhiC7YpzkBSxXJHpp0quD0vDcjSj6ibrp/5B5OvThwCm5fLzzd3jmItjLa66PvW5eZPqo",
	"VB5xfm35Wm7OuxNsJdolZSHWPz35tZ3bT8I+oiILNNfPA5HwAx1tne7QDUFMH6x+NyJPUEHyEQ8yeE67",
	"OiF3wBGSYo82xc1ri6LHZHzR+QLKoLFPHUWXxrhBFudT5ywTgxnNx/wzAAD//7836NnOBQAA",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
