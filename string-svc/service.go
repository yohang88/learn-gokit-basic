package main

import "strings"

func NewService() Service {
    return service{}
}

type service struct{}

type Service interface {
    HealthCheck() (string, error)
    Uppercase(s string) (string, error)
}

func (service) HealthCheck() (string, error) {
    return "ok", nil
}

func (service) Uppercase(s string) (string, error) {
    return strings.ToUpper(s), nil
}