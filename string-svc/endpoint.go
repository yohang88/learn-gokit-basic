package main

import (
    "context"
    "github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
    GetHealthCheckEndpoint endpoint.Endpoint
    GetUppercaseEndpoint endpoint.Endpoint
}

func makeServerEndpoints(s Service) Endpoints {
    return Endpoints{
        GetHealthCheckEndpoint: makeHealthCheckEndpoint(s),
        GetUppercaseEndpoint: makeUppercaseEndpoint(s),
    }
}
func makeHealthCheckEndpoint(svc Service) endpoint.Endpoint {
    return func(_ context.Context, request interface{}) (interface{}, error) {
        status, _ := svc.HealthCheck()

        return healthCheckResponse{status}, nil
    }
}
func makeUppercaseEndpoint(svc Service) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        req := request.(uppercaseRequest)

        result, _ := svc.Uppercase(req.Input)

        return uppercaseResponse{result}, nil
    }
}