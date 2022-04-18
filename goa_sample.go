package goasampleapi

import (
	"fmt"
	"context"
	goasample "goa-sample/gen/goa_sample"
	"log"
)

// goa-sample service example implementation.
// The example methods log the requests and return zero values.
type goaSamplesrvc struct {
	logger *log.Logger
}

// NewGoaSample returns the goa-sample service implementation.
func NewGoaSample(logger *log.Logger) goasample.Service {
	return &goaSamplesrvc{logger}
}

// Echo implements echo.
func (s *goaSamplesrvc) Echo(ctx context.Context, p *goasample.EchoPayload) (res string, err error) {
	s.logger.Printf("Request name:%s", p.Name)
	s.logger.Printf("Request age:%d", p.Age)
	return fmt.Sprintf("Your name: %s, Your age: %d", p.Name, p.Age), nil
}
