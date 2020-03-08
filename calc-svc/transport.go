package main

import (
    "context"
    "encoding/json"
    "github.com/go-kit/kit/log"
    "github.com/go-kit/kit/transport"
    httptransport "github.com/go-kit/kit/transport/http"
    "github.com/gorilla/mux"
    "net/http"
)

func MakeHTTPHandler(s Service, logger log.Logger) http.Handler {
    r := mux.NewRouter()
    e := makeServerEndpoints(s)

    options := []httptransport.ServerOption{
        httptransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
        httptransport.ServerErrorEncoder(encodeError),
    }

    r.Methods("GET").Path("/health").Handler(httptransport.NewServer(
        e.GetHealthCheckEndpoint,
        decodeHealthCheckRequest,
        encodeResponse,
        options...,
    ))

    r.Methods("POST").Path("/area").Handler(httptransport.NewServer(
        e.GetAreaEndpoint,
        decodeAreaRequest,
        encodeResponse,
        options...,
    ))

    return r
}

type healthCheckRequest struct{}

type healthCheckResponse struct {
    Status string `json:"status"`
}

type areaRequest struct{
    Width uint16 `json:"width"`
    Length uint16 `json:"length"`
}

type areaResponse struct {
    Result uint16 `json:"result"`
}

func decodeHealthCheckRequest(_ context.Context, r *http.Request) (interface{}, error) {
    var req healthCheckRequest
    return req, nil
}

func decodeAreaRequest(_ context.Context, r *http.Request) (interface{}, error) {
    var req areaRequest

    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        return nil, err
    }

    return req, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
    w.Header().Set("Content-Type", "application/json; charset=utf-8")

    return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
    if err == nil {
        panic("encodeError with nil error")
    }

    w.Header().Set("Content-Type", "application/json; charset=utf-8")

    json.NewEncoder(w).Encode(map[string]interface{}{
        "error": err.Error(),
    })
}
