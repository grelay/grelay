package grelay

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockService struct{}

func (m mockService) Ping() error {
	return nil
}

func TestGrelayEnqueueWithOneItemInQueueShouldReturnOneItem(t *testing.T) {
	sMock := NewGrelayService(DefaultConfiguration, mockService{})
	sMock2 := NewGrelayService(DefaultConfiguration, mockService{})

	m := map[string]*Service{
		"test":  sMock,
		"test2": sMock2,
	}
	g := NewGrelay(m)
	gr2 := g.CreateRequest()
	gr2 = gr2.Enqueue("test", func() (interface{}, error) { return nil, nil })

	val := reflect.ValueOf(gr2)
	queueFuncs := val.FieldByName("QueueFuncs").Interface().([]GrelayRequestFunc)

	assert.Equal(t, 1, len(queueFuncs))
}

func TestGrelayEnqueueWithTwoItemsInQueueShouldReturnTwoItems(t *testing.T) {
	sMock := NewGrelayService(DefaultConfiguration, mockService{})
	sMock2 := NewGrelayService(DefaultConfiguration, mockService{})

	m := map[string]*Service{
		"test":  sMock,
		"test2": sMock2,
	}
	g := NewGrelay(m)

	gr2 := g.CreateRequest()
	gr2 = gr2.Enqueue("test", func() (interface{}, error) { return nil, nil })
	gr2 = gr2.Enqueue("test2", func() (interface{}, error) { return nil, nil })

	val := reflect.ValueOf(gr2)
	queueFuncs := val.FieldByName("QueueFuncs").Interface().([]GrelayRequestFunc)

	assert.Equal(t, 2, len(queueFuncs))
}

func TestGrelayEnqueueTwoDifferentGrelays(t *testing.T) {
	sMock := NewGrelayService(DefaultConfiguration, mockService{})
	sMock2 := NewGrelayService(DefaultConfiguration, mockService{})
	sMock3 := NewGrelayService(DefaultConfiguration, mockService{})

	m := map[string]*Service{
		"test":  sMock,
		"test2": sMock2,
		"test3": sMock3,
	}
	g := NewGrelay(m)

	gr2 := g.CreateRequest()
	gr2 = gr2.Enqueue("test", func() (interface{}, error) { return nil, nil })
	gr2 = gr2.Enqueue("test2", func() (interface{}, error) { return nil, nil })

	val := reflect.ValueOf(gr2)
	queueFuncs := val.FieldByName("QueueFuncs").Interface().([]GrelayRequestFunc)
	assert.Equal(t, 2, len(queueFuncs))

	gr3 := g.CreateRequest()
	gr3 = gr3.Enqueue("test", func() (interface{}, error) { return nil, nil })
	gr3 = gr3.Enqueue("test2", func() (interface{}, error) { return nil, nil })
	gr3 = gr3.Enqueue("test3", func() (interface{}, error) { return nil, nil })

	val2 := reflect.ValueOf(gr3)
	queueFuncs2 := val2.FieldByName("QueueFuncs").Interface().([]GrelayRequestFunc)

	assert.Equal(t, 3, len(queueFuncs2))
}
