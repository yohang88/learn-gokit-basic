package main

func NewService() Service {
    return service{}
}

type service struct{}

type Service interface {
    HealthCheck() (string, error)
}

func (service) HealthCheck() (string, error) {
    return "ok", nil
}