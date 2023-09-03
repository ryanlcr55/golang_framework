// Package genapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package genapi

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
)

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (POST /post)
	PostCreate(c *gin.Context)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// PostCreate operation middleware
func (siw *ServerInterfaceWrapper) PostCreate(c *gin.Context) {

	c.Set(BearerAuthScopes, []string{""})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.PostCreate(c)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL      string
	Middlewares  []MiddlewareFunc
	ErrorHandler func(*gin.Context, error, int)
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router *gin.Engine, si ServerInterface) *gin.Engine {
	return RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router *gin.Engine, si ServerInterface, options GinServerOptions) *gin.Engine {

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

	router.POST(options.BaseURL+"/post", wrapper.PostCreate)

	return router
}

type PostCreateRequestObject struct {
	Body *PostCreateJSONRequestBody
}

type PostCreateResponseObject interface {
	VisitPostCreateResponse(w http.ResponseWriter) error
}

type PostCreate200Response struct {
}

func (response PostCreate200Response) VisitPostCreateResponse(w http.ResponseWriter) error {
	w.WriteHeader(200)
	return nil
}

type PostCreatedefaultJSONResponse struct {
	Body       Error
	StatusCode int
}

func (response PostCreatedefaultJSONResponse) VisitPostCreateResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)

	return json.NewEncoder(w).Encode(response.Body)
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {

	// (POST /post)
	PostCreate(ctx context.Context, request PostCreateRequestObject) (PostCreateResponseObject, error)
}

type StrictHandlerFunc func(ctx *gin.Context, args interface{}) (interface{}, error)

type StrictMiddlewareFunc func(f StrictHandlerFunc, operationID string) StrictHandlerFunc

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
}

// PostCreate operation middleware
func (sh *strictHandler) PostCreate(ctx *gin.Context) {
	var request PostCreateRequestObject

	var body PostCreateJSONRequestBody
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.Status(http.StatusBadRequest)
		ctx.Error(err)
		return
	}
	request.Body = &body

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PostCreate(ctx, request.(PostCreateRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostCreate")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
	} else if validResponse, ok := response.(PostCreateResponseObject); ok {
		if err := validResponse.VisitPostCreateResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("Unexpected response type: %T", response))
	}
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/6xTzY7bPAx8lYDfd9SuHedS6Lb9OfTSLdDegqBQZCbWwpZUiko3CPLuBWUn2U1y7MkE",
	"xeFwhvQBbBhi8Og5gT5Ash0OpoRfiAJJEClEJHZY0gOmZLYoIe8jgobE5PwWjkcFhL+zI2xBL8+FK3Uq",
	"DOsXtAxHBd9D4k+EhvGWwAbP6FlCfDVD7AUaQ+LZ6UVdMytIXfjzdIVp6qZ+qBcPTQMKNoEGw6ChFdJ7",
	"HdhwLgOgzwPo5Vw1l9GdZ9wiSSE77vHOdGP+TueckL6F9wjJ/Zo3zeLDLeLKxwl+IlZwsWGa+Sz/1mpp",
	"5vymsLeYLLnILnjQ8PP58/O5qS4bmf1A2jkrHDukNNbNH+vHWmSEiN5EBxoWJaUgGu6KY5UYUDY5fWWf",
	"Roi+tlPvadujNEz8MbT7q22bGHtnC6x6SUJ+ukeJ/ifcgIb/qsvBVtO1Vm8IiuL3Sjm0Ad56ypSxmJxi",
	"8Gm8uqaubz0qyNJvY3LP/2ze8de6M2r2+BrRMrYzPNVIVUKbyfEe9PIAazSE9JS5A71cHVfyTLKw8pqp",
	"Bw0dc0y6qvpgTd+FxJWsTsHOkDPrftQseW8GHIVPEuEMEfLV8W8AAAD//wKvvRYiBAAA",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
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
	var res = make(map[string]func() ([]byte, error))
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
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
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
