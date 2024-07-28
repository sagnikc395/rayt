package main

type Vector struct {
	X float64
	Y float64
	Z float64
}

func InitVector(e0, e1, e2 float64) *Vector {
	return &Vector{
		X: e0,
		Y: e1,
		Z: e2,
	}
}

func (v *Vector) Negate()  {
	v.X * -1,
}
