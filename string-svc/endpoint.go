package main

import (
    "context"
    "github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
    GetHealthCheckEndpoint endpoint.Endpoint
}

func makeServerEndpoints(s Service) Endpoints {
    return Endpoints{
        GetHealthCheckEndpoint: makeHealthCheckEndpoint(s),
    }
}
func makeHealthCheckEndpoint(svc Service) endpoint.Endpoint {
    return func(_ context.Context, request interface{}) (interface{}, error) {
        status, _ := svc.HealthCheck()

        return healthCheckResponse{status}, nil
    }
}