package states_test

import (
	"testing"

	"github.com/grelay/grelay/internal/states"
	"github.com/stretchr/testify/assert"
)

func TestStateClosed(t *testing.T) {
	assert.Equal(t, "CLOSED", string(states.Closed), "string(closed) should resturn EQUAL when CLOSE")
	assert.True(t, string(states.Closed) == "CLOSED", "string(closed) should resturn TRUE when CLOSED")
	assert.False(t, string(states.Closed) == "CLOSE", "string(closed) should resturn FALSE when CLOSE")
}

func TestStateOpen(t *testing.T) {
	assert.Equal(t, "OPEN", string(states.Open), "string(open) should resturn EQUAL when OPEN")
	assert.True(t, string(states.Open) == "OPEN", "string(open) should resturn TRUE when OPEN")
	assert.False(t, string(states.Open) == "OPENED", "string(open) should resturn FALSE when OPENED")
}

func TestStateHalfOpen(t *testing.T) {
	assert.Equal(t, "HALF-OPEN", string(states.HalfOpen), "string(halfOpen) should resturn EQUAL when HALF-OPEN")
	assert.True(t, string(states.HalfOpen) == "HALF-OPEN", "string(halfOpen) should resturn TRUE when HALF-OPEN")
	assert.False(t, string(states.HalfOpen) == "OPEN", "string(halfOpen) should resturn FALSE when OPEN")
}
