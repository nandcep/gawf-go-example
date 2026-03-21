package math_test

import (
	"nandcep/go-gawf-example/math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	c := math.Calculation{NumA: 5, NumB: 4}
	result := c.Add()
	assert.Equal(t, float32(9), result)
}
