package main

import "math"

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

func (v *Vector) Negate() Vector {
	return Vector{
		v.X * -1,
		v.Y * -1,
		v.Z * -1,
	}
}

func (v *Vector) PlueEqual(vec2 Vector) {
	v.X += vec2.X
	v.Y += vec2.Y
	v.Z += vec2.Z
}

func (v *Vector) MultipleEqual(t float64) {
	v.X *= t
	v.Y *= t
	v.Z *= t
}

func (v *Vector) DivideEqual(t float64) {
	v.MultipleEqual(1 / t)
}

func (v *Vector) length() float64 {
	return math.Sqrt(v.lengthSquared())
}

func (v *Vector) lengthSquared() float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}
