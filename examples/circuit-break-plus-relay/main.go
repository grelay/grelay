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
	config1 := grelay.DefaultConfiguration
	config1.RetryPeriod = 5 * time.Second // Each 5s, check if service is ok
	config1.Service = service1
	config1.Timeout = 500 * time.Millisecond // Limit timeout to 0.5s, if 0.5s hits, increase threshould
	config1.Threshould = 1

	service2 := &myService2{}
	config2 := grelay.DefaultConfiguration
	config2.RetryPeriod = 5 * time.Second // Each 5s, check if service is ok
	config2.Service = service2
	config2.Timeout = 1 * time.Second // Limit timeout to 1s, if 1s hits, increase threshould
	config2.Threshould = 5

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
