package grelay

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestEmptyGrelayConfig(t *testing.T) {
	c := NewGrelayConfig()

	assert.Equal(t, c.retryTimePeriod, 10*time.Second, "RetryTimePeriod should be 10 sec with default value")
	assert.Equal(t, c.serviceTimeout, 10*time.Second, "ServiceTimeout should be 10 sec with default value")
	assert.Equal(t, c.serviceThreshould, int64(10), "ServiceThreshould should be 10 with default value")
}

func TestEmptyGrelayConfigWithRetryTimePeriod(t *testing.T) {
	c := NewGrelayConfig().WithRetryTimePeriod(20 * time.Second)

	assert.Equal(t, c.retryTimePeriod, 20*time.Second, "RetryTimePeriod should be 20 sec with default value")
	assert.Equal(t, c.serviceTimeout, 10*time.Second, "ServiceTimeout should be 10 sec with default value")
	assert.Equal(t, c.serviceThreshould, int64(10), "ServiceThreshould should be 10 with default value")
}

func TestEmptyGrelayConfigWithServiceTimeout(t *testing.T) {
	c := NewGrelayConfig().WithServiceTimeout(20 * time.Second)

	assert.Equal(t, c.retryTimePeriod, 10*time.Second, "RetryTimePeriod should be 10 sec with default value")
	assert.Equal(t, c.serviceTimeout, 20*time.Second, "ServiceTimeout should be 20 sec with default value")
	assert.Equal(t, c.serviceThreshould, int64(10), "ServiceThreshould should be 10 with default value")
}

func TestEmptyGrelayConfigWithServiceThreshould(t *testing.T) {
	c := NewGrelayConfig().WithServiceThreshould(20)

	assert.Equal(t, c.retryTimePeriod, 10*time.Second, "RetryTimePeriod should be 10 sec with default value")
	assert.Equal(t, c.serviceTimeout, 10*time.Second, "ServiceTimeout should be 10 sec with default value")
	assert.Equal(t, c.serviceThreshould, int64(20), "ServiceThreshould should be 20 with default value")
}

func TestEmptyGrelayConfigWithServiceThreshouldAndWithServiceTimeout(t *testing.T) {
	c := NewGrelayConfig()
	c = c.WithServiceThreshould(20)
	c = c.WithServiceTimeout(20 * time.Second)

	assert.Equal(t, c.retryTimePeriod, 10*time.Second, "RetryTimePeriod should be 10 sec with default value")
	assert.Equal(t, c.serviceTimeout, 20*time.Second, "ServiceTimeout should be 10 sec with default value")
	assert.Equal(t, c.serviceThreshould, int64(20), "ServiceThreshould should be 20 with default value")
}
