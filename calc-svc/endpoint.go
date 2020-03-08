package main

import (
    "context"
    "github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
    GetHealthCheckEndpoint endpoint.Endpoint
    GetAreaEndpoint endpoint.Endpoint
}

func makeServerEndpoints(svc Service) Endpoints {
    return Endpoints{
        GetHealthCheckEndpoint: makeHealthCheckEndpoint(svc),
        GetAreaEndpoint: makeAreaEndpoint(svc),
    }
}

func makeHealthCheckEndpoint(svc Service) endpoint.Endpoint {
    return func(_ context.Context, request interface{}) (interface{}, error) {
        status, _ := svc.HealthCheck()

        return healthCheckResponse{status}, nil
    }
}

func makeAreaEndpoint(svc Service) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        req := request.(areaRequest)

        result, _ := svc.Area(req.Width, req.Length)

        return areaResponse{result}, nil
    }
}