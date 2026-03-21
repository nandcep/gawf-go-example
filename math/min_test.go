package math_test

import (
	"nandcep/go-gawf-example/math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMin(t *testing.T) {
	c := math.Calculation{NumA: 7, NumB: 3}
	result := c.Min()
	assert.Equal(t, float32(4), result)
}
