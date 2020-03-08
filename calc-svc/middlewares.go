package main

import (
    "github.com/go-kit/kit/log"
    "time"
)

type Middleware func(Service) Service

func LoggingMiddleware(logger log.Logger) Middleware {
    return func(next Service) Service {
        return &loggingMiddleware{next: next, logger: logger}
    }
}

type loggingMiddleware struct {
    next   Service
    logger log.Logger
}

func (mw loggingMiddleware) Area(width uint16, length uint16) (result uint16, err error) {
    defer func(begin time.Time) {
        mw.logger.Log("method", "GetArea", "time", time.Since(begin), "err", err)
    }(time.Now())

    return mw.next.Area(width, length)
}

func (mw loggingMiddleware) HealthCheck() (s string, err error) {
    defer func(begin time.Time) {
        mw.logger.Log("method", "GetHealthCheck", "time", time.Since(begin), "err", err)
    }(time.Now())

    return mw.next.HealthCheck()
}
