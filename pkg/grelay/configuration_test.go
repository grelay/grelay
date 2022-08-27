package grelay_test

import (
	"testing"
	"time"

	"github.com/grelay/grelay/pkg/grelay"
	"github.com/stretchr/testify/assert"
)

func TestConfiguration(t *testing.T) {
	type output struct {
		WithGo      bool
		Threshould  int64
		RetryPeriod time.Duration
		Timeout     time.Duration
		Service     grelay.Pingable
	}
	type testCase struct {
		name string
		in   grelay.Configuration
		out  output
	}
	tests := []testCase{
		{
			name: "should return default configuration when uses the DefaultConfiguration",
			in:   grelay.DefaultConfiguration,
			out: output{
				RetryPeriod: 10 * time.Second,
				Timeout:     10 * time.Second,
				Threshould:  10,
				Service:     nil,
				WithGo:      true,
			},
		},
		{
			name: "should return custom configuration when changes the DefaultConfiguration",
			in: createCustomConfigurationFromDefaultConfiguration(grelay.Configuration{
				RetryPeriod: 5 * time.Second,
				Timeout:     7 * time.Second,
				Service:     nil,
				Threshould:  5,
				WithGo:      false,
			}),
			out: output{
				RetryPeriod: 5 * time.Second,
				Timeout:     7 * time.Second,
				Service:     nil,
				Threshould:  5,
				WithGo:      false,
			},
		},
		{
			name: "should return custom configuration when creates a new Configuration",
			in: grelay.Configuration{
				RetryPeriod: 5 * time.Second,
				Timeout:     7 * time.Second,
				Service:     nil,
				Threshould:  5,
				WithGo:      false,
			},
			out: output{
				RetryPeriod: 5 * time.Second,
				Timeout:     7 * time.Second,
				Service:     nil,
				Threshould:  5,
				WithGo:      false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.in.RetryPeriod, tt.out.RetryPeriod)
			assert.Equal(t, tt.in.Timeout, tt.out.Timeout)
			assert.Equal(t, tt.in.Threshould, tt.out.Threshould)
			assert.Equal(t, tt.in.Service, tt.out.Service)
			assert.Equal(t, tt.in.WithGo, tt.out.WithGo)
		})
	}
}

func createCustomConfigurationFromDefaultConfiguration(custom grelay.Configuration) grelay.Configuration {
	config := grelay.DefaultConfiguration

	config.Service = custom.Service
	config.RetryPeriod = custom.RetryPeriod
	config.Threshould = custom.Threshould
	config.Timeout = custom.Timeout
	config.WithGo = custom.WithGo

	return config
}
