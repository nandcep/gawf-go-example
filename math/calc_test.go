package math_test

import (
	"nandcep/go-gawf-example/math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculationStruct(t *testing.T) {
	c := math.Calculation{NumA: 1, NumB: 2}
	assert.Equal(t, 1, c.NumA)
	assert.Equal(t, 2, c.NumB)
}
