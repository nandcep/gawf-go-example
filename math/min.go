package math

func (c Calculation) Min() float32 {
	return float32(c.NumA) - float32(c.NumB)
}
