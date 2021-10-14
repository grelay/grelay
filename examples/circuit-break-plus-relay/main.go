package main

import (
	"fmt"
	"time"

	"github.com/grelay/grelay/v1"
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

	// services that grelay will manage
	services := map[string]grelay.GrelayService{
		serviceTag:  grelay.NewGrelayService(configGrelayService1()),
		service2Tag: grelay.NewGrelayService(configGrelayService2()),
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

func configGrelayService1() grelay.GrelayConfig {
	service := &myService{}
	config := grelay.NewGrelayConfig()
	config = config.WithRetryTimePeriod(5 * time.Second) // Each 5s, check if service is ok
	config = config.WithGrelayService(service)
	config = config.WithServiceTimeout(500 * time.Millisecond) // Limit timeout to 0.5s, if pass of that, increase threshould
	config = config.WithServiceThreshould(1)                   // Set the number of threshould allowed.

	return config
}

func configGrelayService2() grelay.GrelayConfig {
	service := &myService2{}
	config := grelay.NewGrelayConfig()
	config = config.WithRetryTimePeriod(5 * time.Second) // Each 5s, check if service is ok
	config = config.WithGrelayService(service)
	config = config.WithServiceTimeout(1 * time.Second) // Limit timeout to 1s, if pass of that, increase threshould
	config = config.WithServiceThreshould(5)            // Set the number of threshould allowed.

	return config
}
