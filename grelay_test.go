package grelay

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

/*

- Enqueue

* Enqueue uma ũnica fila, deve retornar um único item
* Enqueue dois itens, deve retornar dois itens
* Dois Enqueues diferentes um com dois itens e outro com tres itens, devem estar com tamanhos diferentes

*/

func TestGrelayEnqueueWithOneItemInQueueShouldReturnOneItem(t *testing.T) {
	sMock := new(grelayServiceMock)
	sMock.On("Exec", mock.Anything).Return(nil, nil)

	sMock2 := new(grelayServiceMock)
	sMock2.On("Exec", mock.Anything).Return(nil, nil)

	m := map[string]GrelayService{
		"test":  sMock,
		"test2": sMock2,
	}
	g := NewGrelay(m)
	gr := g.Enqueue("test", func() (interface{}, error) { return nil, nil })

	val := reflect.ValueOf(gr)
	queueFuncs := val.FieldByName("QueueFuncs").Interface().([]grelayRequestQueueStruct)

	assert.Equal(t, 1, len(queueFuncs))
}

func TestGrelayEnqueueWithTwoItemsInQueueShouldReturnTwoItems(t *testing.T) {
	sMock := new(grelayServiceMock)
	sMock.On("Exec", mock.Anything).Return(nil, nil)

	sMock2 := new(grelayServiceMock)
	sMock2.On("Exec", mock.Anything).Return(nil, nil)

	m := map[string]GrelayService{
		"test":  sMock,
		"test2": sMock2,
	}
	g := NewGrelay(m)
	gr := g.Enqueue("test", func() (interface{}, error) { return nil, nil })
	gr = gr.Enqueue("test2", func() (interface{}, error) { return nil, nil })

	val := reflect.ValueOf(gr)
	queueFuncs := val.FieldByName("QueueFuncs").Interface().([]grelayRequestQueueStruct)

	assert.Equal(t, 2, len(queueFuncs))
}

func TestGrelayEnqueueTwoDifferentGrelays(t *testing.T) {
	sMock := new(grelayServiceMock)
	sMock.On("exec", mock.Anything).Return(nil, nil)

	sMock2 := new(grelayServiceMock)
	sMock2.On("exec", mock.Anything).Return(nil, nil)

	sMock3 := new(grelayServiceMock)
	sMock3.On("exec", mock.Anything).Return(nil, nil)

	m := map[string]GrelayService{
		"test":  sMock,
		"test2": sMock2,
		"test3": sMock3,
	}
	g := NewGrelay(m)

	gr := g.Enqueue("test", func() (interface{}, error) { return nil, nil })
	gr = gr.Enqueue("test2", func() (interface{}, error) { return nil, nil })
	val := reflect.ValueOf(gr)
	queueFuncs := val.FieldByName("QueueFuncs").Interface().([]grelayRequestQueueStruct)
	assert.Equal(t, 2, len(queueFuncs))

	gr2 := g.Enqueue("test", func() (interface{}, error) { return nil, nil })
	gr2 = gr2.Enqueue("test2", func() (interface{}, error) { return nil, nil })
	gr2 = gr2.Enqueue("test3", func() (interface{}, error) { return nil, nil })

	val2 := reflect.ValueOf(gr2)
	queueFuncs2 := val2.FieldByName("QueueFuncs").Interface().([]grelayRequestQueueStruct)

	assert.Equal(t, 3, len(queueFuncs2))
}
