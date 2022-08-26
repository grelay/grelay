package main

import (
	"fmt"
	"time"

	"github.com/grelay/grelay/pkg/grelay"
)

const serviceTag = "myService"

type myService struct{}
type myResponse struct {
	name string
}

func (s *myService) Ping() error {
	// make request to check if service is ok, if return anything != nil, means that server has some problem.
	return nil
}

func main() {
	service := &myService{}
	config := grelay.NewGrelayConfig()
	config = config.WithRetryTimePeriod(5 * time.Second) // Each 5s, check if service is ok
	config = config.WithGrelayService(service)
	config = config.WithServiceTimeout(1 * time.Second) // Limit timeout to 1s, if pass of that, increase threshould
	config = config.WithServiceThreshould(5)            // Set the number of threshould allowed.

	// services that grelay will manage
	services := map[string]grelay.GrelayService{
		serviceTag: grelay.NewGrelayService(config),
	}

	g := grelay.NewGrelay(services)
	gr := g.CreateRequest()
	gr = gr.Enqueue(serviceTag, func() (interface{}, error) {
		// make request that you want
		return myResponse{
			name: "gRelay",
		}, nil
	})

	val, err := gr.Exec()
	if err != nil {
		panic(err)
	}
	fmt.Println(val.(myResponse).name)
}
