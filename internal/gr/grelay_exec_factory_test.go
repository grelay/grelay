package gr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGrelayFactoryShouldReturnGrelayExecWithGo(t *testing.T) {
	c := NewGrelayConfig()
	c = c.WithGo()

	g := getGrelayExec(c)

	assert.IsType(t, grelayExecWithGo{}, g)
}

func TestGrelayFactoryShouldReturnGrelayExecDefault(t *testing.T) {
	c := NewGrelayConfig()
	c.withGo = false

	g := getGrelayExec(c)

	assert.IsType(t, grelayExecWithGo{}, g)
}
