package grelay

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGrelayFactoryShouldReturnGrelayExecWithGo(t *testing.T) {
	c := DefaultConfiguration

	g := getGrelayExec(c)

	assert.IsType(t, grelayExecWithGo{}, g)
}

func TestGrelayFactoryShouldReturnGrelayExecDefault(t *testing.T) {
	c := DefaultConfiguration
	c.WithGo = false

	g := getGrelayExec(c)

	assert.IsType(t, grelayExecWithGo{}, g)
}
