package main

import (
	"fmt"
	"math"
)

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

//vector utility functions

func (v *Vector) Print() string {
	return fmt.Sprintf("%f %f %f", v.X, v.Y, v.Z)
}

// adds 2 vector and returns a new vector instance
func VectorAdd(v1, v2 Vector) Vector {
	return Vector{
		v1.X + v2.X,
		v1.Y + v2.Y,
		v1.Z + v2.Z,
	}
}

// diffs 2 vector and returns a new vector instance
func VectorDiff(v1, v2 Vector) Vector {
	return Vector{
		v1.X - v2.X,
		v1.Y - v2.Y,
		v1.Z - v2.Z,
	}
}

// multiply 2 vectors and returns a new vector instance
func VectorMultiply(v1, v2 Vector) Vector {
	return Vector{
		v1.X * v2.X,
		v1.Y * v2.Y,
		v1.Z * v2.Z,
	}
}

func VectorBumpT(t float64, v Vector) Vector {
	return Vector{
		t * v.X,
		t * v.Y,
		t * v.Z,
	}
}

func VectorScalarMultiply(v Vector, t float64) Vector {
	return Vector{
		t * v.X,
		t * v.Y,
		t * v.Z,
	}
}

func VectorScalarDivide(v Vector, t float64) Vector {
	return VectorScalarMultiply(v, 1/t)
}

func DotProduct(v1, v2 Vector) float64 {
	return v1.X*v2.X + v1.Y*v2.Y + v1.Z*v2.Z
}

func CrossProduct(v1, v2 Vector) Vector {
	return Vector{
		v1.Y*v2.Z - (v1.Z * v2.Y),
		v1.Z*v2.X - (v1.X * v2.Z),
		v1.X*v2.Y - (v1.Y * v2.X),
	}
}
