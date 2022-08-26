package gr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStateClosed(t *testing.T) {
	assert.Equal(t, "CLOSED", string(closed), "string(closed) should resturn EQUAL when CLOSE")
	assert.True(t, string(closed) == "CLOSED", "string(closed) should resturn TRUE when CLOSED")
	assert.False(t, string(closed) == "CLOSE", "string(closed) should resturn FALSE when CLOSE")
}

func TestStateOpen(t *testing.T) {
	assert.Equal(t, "OPEN", string(open), "string(open) should resturn EQUAL when OPEN")
	assert.True(t, string(open) == "OPEN", "string(open) should resturn TRUE when OPEN")
	assert.False(t, string(open) == "OPENED", "string(open) should resturn FALSE when OPENED")
}

func TestStateHalfOpen(t *testing.T) {
	assert.Equal(t, "HALF-OPEN", string(halfOpen), "string(halfOpen) should resturn EQUAL when HALF-OPEN")
	assert.True(t, string(halfOpen) == "HALF-OPEN", "string(halfOpen) should resturn TRUE when HALF-OPEN")
	assert.False(t, string(halfOpen) == "OPEN", "string(halfOpen) should resturn FALSE when OPEN")
}
