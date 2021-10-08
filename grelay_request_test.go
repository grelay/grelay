package grelay

import (
	"reflect"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGrelayRequestEnqueueShouldIncludeInList(t *testing.T) {
	c := NewGrelayConfig()
	c = c.WithRetryTimePeriod(500 * time.Millisecond)
	s := NewGrelayService(c)
	m := map[string]GrelayService{
		"test": s,
	}
	var gr GrelayRequest = grelayRequestImpl{
		mapServices: m,
		mu:          &sync.RWMutex{},
	}
	gr = gr.Enqueue("test", func() (interface{}, error) { return nil, nil })

	val := reflect.ValueOf(gr)
	queueFuncs := val.FieldByName("QueueFuncs").Interface().([]grelayRequestQueueStruct)
	assert.Equal(t, 1, len(queueFuncs))
}

func TestGrelayRequestEnqueueShouldNotIncludeInList(t *testing.T) {
	c := NewGrelayConfig()
	s := NewGrelayService(c)
	m := map[string]GrelayService{
		"test": s,
	}
	var gr GrelayRequest = grelayRequestImpl{
		mapServices: m,
		mu:          &sync.RWMutex{},
	}
	gr = gr.Enqueue("test2", func() (interface{}, error) { return nil, nil })

	val := reflect.ValueOf(gr)
	queueFuncs := val.FieldByName("QueueFuncs").Interface().([]grelayRequestQueueStruct)
	assert.Equal(t, 0, len(queueFuncs))
}

func TestGrelayRequestExecWithEmptyQueueShouldReturnErrGrelayAllRequestsOpened(t *testing.T) {
	var gr GrelayRequest = grelayRequestImpl{
		mu: &sync.RWMutex{},
	}
	val, err := gr.Exec()
	assert.Nil(t, val)
	assert.Error(t, err, ErrGrelayAllRequestsOpened.Error())
}

func TestGrelayRequestExecWithOneItemInQueueShouldReturnNil(t *testing.T) {
	sMock := new(grelayServiceMock)
	sMock.On("Exec", mock.Anything).Return(nil, nil)

	m := map[string]GrelayService{
		"test": sMock,
	}
	var gr GrelayRequest = grelayRequestImpl{
		mapServices: m,
		mu:          &sync.RWMutex{},
	}
	gr = gr.Enqueue("test", func() (interface{}, error) { return nil, nil })

	val, err := gr.Exec()
	assert.Nil(t, val)
	assert.Nil(t, err)
}

func TestGrelayRequestExecWithOneItemOpenedInQueueShouldReturnErrGrelayAllRequestsOpened(t *testing.T) {
	sMock := new(grelayServiceMock)
	sMock.On("Exec", mock.Anything).Return(nil, ErrGrelayStateOpened)

	m := map[string]GrelayService{
		"test": sMock,
	}
	var gr GrelayRequest = grelayRequestImpl{
		mapServices: m,
		mu:          &sync.RWMutex{},
	}
	gr = gr.Enqueue("test", func() (interface{}, error) { return nil, nil })

	val, err := gr.Exec()
	assert.Nil(t, val)
	assert.Error(t, err, ErrGrelayAllRequestsOpened.Error())
}

func TestGrelayRequestExecWithTwoItemsWithFirstOpenedInQueueShouldReturnNil(t *testing.T) {
	sMock := new(grelayServiceMock)
	sMock.On("Exec", mock.Anything).Return(nil, ErrGrelayStateOpened)

	sMock2 := new(grelayServiceMock)
	sMock2.On("Exec", mock.Anything).Return(nil, nil)

	m := map[string]GrelayService{
		"test":  sMock,
		"test2": sMock2,
	}
	var gr GrelayRequest = grelayRequestImpl{
		mapServices: m,
		mu:          &sync.RWMutex{},
	}
	gr = gr.Enqueue("test", func() (interface{}, error) { return nil, nil })
	gr = gr.Enqueue("test2", func() (interface{}, error) { return nil, nil })

	val, err := gr.Exec()
	assert.Nil(t, val)
	assert.Nil(t, err)
}

func TestGrelayRequestExecWithOneItemInQueueReturningErrGrelayServiceTimedoutShouldReturnErrGrelayServiceTimedout(t *testing.T) {
	sMock := new(grelayServiceMock)
	sMock.On("Exec", mock.Anything).Return(nil, ErrGrelayServiceTimedout)

	m := map[string]GrelayService{
		"test": sMock,
	}
	var gr GrelayRequest = grelayRequestImpl{
		mapServices: m,
		mu:          &sync.RWMutex{},
	}
	gr = gr.Enqueue("test", func() (interface{}, error) { return nil, nil })

	val, err := gr.Exec()
	assert.Nil(t, val)
	assert.Error(t, err, ErrGrelayServiceTimedout.Error())
}

func TestGrelayRequestExecWithTwoItemsBothOpenedInQueueShouldReturnErrGrelayAllRequestsOpened(t *testing.T) {
	sMock := new(grelayServiceMock)
	sMock.On("Exec", mock.Anything).Return(nil, ErrGrelayStateOpened)

	sMock2 := new(grelayServiceMock)
	sMock2.On("Exec", mock.Anything).Return(nil, ErrGrelayStateOpened)

	m := map[string]GrelayService{
		"test":  sMock,
		"test2": sMock2,
	}
	var gr GrelayRequest = grelayRequestImpl{
		mapServices: m,
		mu:          &sync.RWMutex{},
	}
	gr = gr.Enqueue("test", func() (interface{}, error) { return nil, nil })
	gr = gr.Enqueue("test2", func() (interface{}, error) { return nil, nil })

	val, err := gr.Exec()
	assert.Nil(t, val)
	assert.Error(t, err, ErrGrelayAllRequestsOpened)
}
