package main

func NewService() Service {
    return service{}
}

type service struct{}

type Service interface {
    HealthCheck() (string, error)
    Area(width uint16, length uint16) (uint16, error)
}

func (service) HealthCheck() (string, error) {
    return "ok", nil
}

func (service) Area(width uint16, length uint16) (uint16, error) {
    return width * length, nil
}
