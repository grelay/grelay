package main

import (
	"fmt"
	"time"

	"github.com/grelay/grelay/pkg/grelay"
)

const serviceTag = "myService"
const service2Tag = "myService2"

type myService struct{}
type myService2 struct{}
type myResponse struct {
	name string
}

func (s *myService) Ping() error {
	// make request to check if service is ok, if return anything != nil, means that server has some problem.
	return nil
}

func (s *myService2) Ping() error {
	// make request to check if service is ok, if return anything != nil, means that server has some problem.
	return nil
}

func main() {
	service1 := &myService{}
	config1 := grelay.NewGrelayConfig()
	config1 = config1.WithRetryTimePeriod(5 * time.Second) // Each 5s, check if service is ok
	config1 = config1.WithGrelayService(service1)
	config1 = config1.WithServiceTimeout(500 * time.Millisecond) // Limit timeout to 0.5s, if pass of that, increase threshould
	config1 = config1.WithServiceThreshould(1)                   // Set the number of threshould allowed.

	service2 := &myService2{}
	config2 := grelay.NewGrelayConfig()
	config2 = config2.WithRetryTimePeriod(5 * time.Second) // Each 5s, check if service is ok
	config2 = config2.WithGrelayService(service2)
	config2 = config2.WithServiceTimeout(1 * time.Second) // Limit timeout to 1s, if pass of that, increase threshould
	config2 = config2.WithServiceThreshould(5)

	// services that grelay will manage
	services := map[string]grelay.GrelayService{
		serviceTag:  grelay.NewGrelayService(config1),
		service2Tag: grelay.NewGrelayService(config2),
	}

	g := grelay.NewGrelay(services)

	for i := 0; i < 2; i++ {
		gr := g.CreateRequest()
		gr = gr.Enqueue(serviceTag, func() (interface{}, error) {
			// make request that you want
			time.Sleep(1 * time.Second)
			return myResponse{
				name: "gRelay",
			}, nil
		})
		gr = gr.Enqueue(service2Tag, func() (interface{}, error) {
			// make request that you want
			return myResponse{
				name: "gRelay2",
			}, nil
		})

		val, err := gr.Exec()
		if err != nil {
			fmt.Println(fmt.Sprintf("error: %s", err.Error()))
			continue
		}
		fmt.Println(val.(myResponse).name)
	}
}
