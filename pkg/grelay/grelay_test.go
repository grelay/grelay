package grelay

import (
	"reflect"
	"testing"

	"github.com/grelay/grelay/internal/gr"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGrelayEnqueueWithOneItemInQueueShouldReturnOneItem(t *testing.T) {
	sMock := new(gr.GrelayServiceMock)
	sMock.On("exec", mock.Anything).Return(nil, nil)

	sMock2 := new(gr.GrelayServiceMock)
	sMock2.On("exec", mock.Anything).Return(nil, nil)

	m := map[string]GrelayService{
		"test": {
			gs: sMock,
		},
		"test2": {
			gs: sMock2,
		},
	}
	g := NewGrelay(m)
	gr2 := g.CreateRequest()
	gr2 = gr2.Enqueue("test", func() (interface{}, error) { return nil, nil })

	val := reflect.ValueOf(gr2)
	queueFuncs := val.FieldByName("QueueFuncs").Interface().([]gr.GrelayRequestFunc)

	assert.Equal(t, 1, len(queueFuncs))
}

func TestGrelayEnqueueWithTwoItemsInQueueShouldReturnTwoItems(t *testing.T) {
	sMock := new(gr.GrelayServiceMock)
	sMock.On("exec", mock.Anything).Return(nil, nil)

	sMock2 := new(gr.GrelayServiceMock)
	sMock2.On("exec", mock.Anything).Return(nil, nil)

	m := map[string]GrelayService{
		"test": {
			gs: sMock,
		},
		"test2": {
			gs: sMock2,
		},
	}
	g := NewGrelay(m)

	gr2 := g.CreateRequest()
	gr2 = gr2.Enqueue("test", func() (interface{}, error) { return nil, nil })
	gr2 = gr2.Enqueue("test2", func() (interface{}, error) { return nil, nil })

	val := reflect.ValueOf(gr2)
	queueFuncs := val.FieldByName("QueueFuncs").Interface().([]gr.GrelayRequestFunc)

	assert.Equal(t, 2, len(queueFuncs))
}

func TestGrelayEnqueueTwoDifferentGrelays(t *testing.T) {
	sMock := new(gr.GrelayServiceMock)
	sMock.On("exec", mock.Anything).Return(nil, nil)

	sMock2 := new(gr.GrelayServiceMock)
	sMock2.On("exec", mock.Anything).Return(nil, nil)

	sMock3 := new(gr.GrelayServiceMock)
	sMock3.On("exec", mock.Anything).Return(nil, nil)

	m := map[string]GrelayService{
		"test": {
			gs: sMock,
		},
		"test2": {
			gs: sMock2,
		},
		"test3": {
			gs: sMock3,
		},
	}
	g := NewGrelay(m)

	gr2 := g.CreateRequest()
	gr2 = gr2.Enqueue("test", func() (interface{}, error) { return nil, nil })
	gr2 = gr2.Enqueue("test2", func() (interface{}, error) { return nil, nil })

	val := reflect.ValueOf(gr2)
	queueFuncs := val.FieldByName("QueueFuncs").Interface().([]gr.GrelayRequestFunc)
	assert.Equal(t, 2, len(queueFuncs))

	gr3 := g.CreateRequest()
	gr3 = gr3.Enqueue("test", func() (interface{}, error) { return nil, nil })
	gr3 = gr3.Enqueue("test2", func() (interface{}, error) { return nil, nil })
	gr3 = gr3.Enqueue("test3", func() (interface{}, error) { return nil, nil })

	val2 := reflect.ValueOf(gr3)
	queueFuncs2 := val2.FieldByName("QueueFuncs").Interface().([]gr.GrelayRequestFunc)

	assert.Equal(t, 3, len(queueFuncs2))
}