package math_test

import (
	"nandcep/go-gawf-example/math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiv(t *testing.T) {
	c := math.Calculation{NumA: 8, NumB: 2}
	result := c.Div()
	assert.Equal(t, float32(4), result)
}
