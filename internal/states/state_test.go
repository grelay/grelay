package states

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStateClosed(t *testing.T) {
	assert.Equal(t, "CLOSED", string(Closed), "string(closed) should resturn EQUAL when CLOSE")
	assert.True(t, string(Closed) == "CLOSED", "string(closed) should resturn TRUE when CLOSED")
	assert.False(t, string(Closed) == "CLOSE", "string(closed) should resturn FALSE when CLOSE")
}

func TestStateOpen(t *testing.T) {
	assert.Equal(t, "OPEN", string(Open), "string(open) should resturn EQUAL when OPEN")
	assert.True(t, string(Open) == "OPEN", "string(open) should resturn TRUE when OPEN")
	assert.False(t, string(Open) == "OPENED", "string(open) should resturn FALSE when OPENED")
}

func TestStateHalfOpen(t *testing.T) {
	assert.Equal(t, "HALF-OPEN", string(HalfOpen), "string(halfOpen) should resturn EQUAL when HALF-OPEN")
	assert.True(t, string(HalfOpen) == "HALF-OPEN", "string(halfOpen) should resturn TRUE when HALF-OPEN")
	assert.False(t, string(HalfOpen) == "OPEN", "string(halfOpen) should resturn FALSE when OPEN")
}
