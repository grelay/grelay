package grelay

import (
	"reflect"
	"sync"
	"testing"
	"time"

	"github.com/grelay/grelay/internal/errs"
	"github.com/grelay/grelay/internal/states"
	"github.com/stretchr/testify/assert"
)

func TestGrelayRequestEnqueueShouldIncludeInList(t *testing.T) {
	c := DefaultConfiguration
	c.RetryPeriod = 500 * time.Millisecond
	s := NewGrelayService(c, nil)
	m := map[string]*Service{
		"test": s,
	}
	var gr GrelayRequest = GrelayRequestImpl{
		MapServices: m,
		Mu:          &sync.RWMutex{},
	}
	gr = gr.Enqueue("test", func() (interface{}, error) { return nil, nil })

	val := reflect.ValueOf(gr)
	queueFuncs := val.FieldByName("QueueFuncs").Interface().([]GrelayRequestFunc)
	assert.Equal(t, 1, len(queueFuncs))
}

func TestGrelayRequestEnqueueShouldNotIncludeInList(t *testing.T) {
	c := DefaultConfiguration
	s := NewGrelayService(c, nil)
	m := map[string]*Service{
		"test": s,
	}
	var gr GrelayRequest = GrelayRequestImpl{
		MapServices: m,
		Mu:          &sync.RWMutex{},
	}
	gr = gr.Enqueue("test2", func() (interface{}, error) { return nil, nil })

	val := reflect.ValueOf(gr)
	queueFuncs := val.FieldByName("QueueFuncs").Interface().([]GrelayRequestFunc)
	assert.Equal(t, 0, len(queueFuncs))
}

func TestGrelayRequestExecWithEmptyQueueShouldReturnErrGrelayAllRequestsOpened(t *testing.T) {
	var gr2 GrelayRequest = GrelayRequestImpl{
		Mu: &sync.RWMutex{},
	}
	val, err := gr2.Exec()
	assert.Nil(t, val)
	assert.Error(t, err, errs.ErrGrelayAllRequestsOpened.Error())
}

func TestGrelayRequestExecWithOneItemInQueueShouldReturnNil(t *testing.T) {
	sMock := NewGrelayService(DefaultConfiguration, mockService{})

	m := map[string]*Service{
		"test": sMock,
	}
	var gr GrelayRequest = GrelayRequestImpl{
		MapServices: m,
		Mu:          &sync.RWMutex{},
	}
	gr = gr.Enqueue("test", func() (interface{}, error) { return nil, nil })

	val, err := gr.Exec()
	assert.Nil(t, val)
	assert.Nil(t, err)
}

func TestGrelayRequestExecWithOneItemOpenedInQueueShouldReturnErrGrelayAllRequestsOpened(t *testing.T) {
	sMock := NewGrelayService(DefaultConfiguration, mockService{})
	sMock.mu.Lock()
	sMock.state = states.Open
	sMock.mu.Unlock()

	m := map[string]*Service{
		"test": sMock,
	}
	var gr2 GrelayRequest = GrelayRequestImpl{
		MapServices: m,
		Mu:          &sync.RWMutex{},
	}
	gr2 = gr2.Enqueue("test", func() (interface{}, error) { return nil, nil })

	val, err := gr2.Exec()
	assert.Nil(t, val)
	assert.Error(t, err, errs.ErrGrelayAllRequestsOpened.Error())
}

func TestGrelayRequestExecWithTwoItemsWithFirstOpenedInQueueShouldReturnNil(t *testing.T) {
	sMock := NewGrelayService(DefaultConfiguration, mockService{})
	sMock.mu.Lock()
	sMock.state = states.Open
	sMock.mu.Unlock()

	sMock2 := NewGrelayService(DefaultConfiguration, mockService{})

	m := map[string]*Service{
		"test":  sMock,
		"test2": sMock2,
	}
	var gr GrelayRequest = GrelayRequestImpl{
		MapServices: m,
		Mu:          &sync.RWMutex{},
	}
	gr = gr.Enqueue("test", func() (interface{}, error) { return nil, nil })
	gr = gr.Enqueue("test2", func() (interface{}, error) { return nil, nil })

	val, err := gr.Exec()
	assert.Nil(t, val)
	assert.Nil(t, err)
}

func TestGrelayRequestExecWithOneItemInQueueReturningErrGrelayServiceTimedoutShouldReturnErrGrelayServiceTimedout(t *testing.T) {
	config := DefaultConfiguration
	config.Timeout = 10 * time.Millisecond
	sMock := NewGrelayService(config, mockService{})

	m := map[string]*Service{
		"test": sMock,
	}
	var gr2 GrelayRequest = GrelayRequestImpl{
		MapServices: m,
		Mu:          &sync.RWMutex{},
	}
	gr2 = gr2.Enqueue("test", func() (interface{}, error) {
		time.Sleep(20 * time.Millisecond)
		return nil, nil
	})

	val, err := gr2.Exec()
	assert.Nil(t, val)
	assert.Error(t, err, errs.ErrGrelayServiceTimedout.Error())
}

func TestGrelayRequestExecWithTwoItemsBothOpenedInQueueShouldReturnErrGrelayAllRequestsOpened(t *testing.T) {
	sMock := NewGrelayService(DefaultConfiguration, mockService{})
	sMock.mu.Lock()
	sMock.state = states.Open
	sMock.mu.Unlock()

	sMock2 := NewGrelayService(DefaultConfiguration, mockService{})
	sMock.mu.Lock()
	sMock2.state = states.Open
	sMock.mu.Unlock()

	m := map[string]*Service{
		"test":  sMock,
		"test2": sMock2,
	}
	var gr2 GrelayRequest = GrelayRequestImpl{
		MapServices: m,
		Mu:          &sync.RWMutex{},
	}
	gr2 = gr2.Enqueue("test", func() (interface{}, error) { return nil, nil })
	gr2 = gr2.Enqueue("test2", func() (interface{}, error) { return nil, nil })

	val, err := gr2.Exec()
	assert.Nil(t, val)
	assert.Error(t, err, errs.ErrGrelayAllRequestsOpened)
}
