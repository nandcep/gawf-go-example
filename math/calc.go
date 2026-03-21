package math

type Calculation struct {
	NumA int
	NumB int
}

type Divider interface {
	Div() float32
}

type Adder interface {
	Add() float32
}

type Minuser interface {
	Min() float32
}

type Multipler interface {
	Mul() float32
}
