package math_test

import (
	"nandcep/go-gawf-example/math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMul(t *testing.T) {
	c := math.Calculation{NumA: 6, NumB: 3}
	result := c.Mul()
	assert.Equal(t, float32(18), result)
}
