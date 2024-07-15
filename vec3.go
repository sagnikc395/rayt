package main

import "math"

//vector type

type Vec3 struct {
	X float64
	Y float64
	Z float64
}

func (v Vec3) Length() float64 {
	//length of a item
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

func (v Vec3) Dot(o Vec3) float64 {
	//dot product of vectors
	return v.X*o.X + v.Y*o.Y + v.Z*o.Z
}

func (v Vec3) Normalize() Vec3 {
	l := v.Length()
	return Vec3{v.X / l, v.Y / l, v.Z / l}
}

func (v Vec3) Add(o Vec3) Vec3 {
	return Vec3{v.X + o.X, v.Y + o.Y, v.Z + o.Z}
}

func (v Vec3) Subtract(o Vec3) Vec3 {
	return Vec3{v.X - o.X, v.Y - o.Y, v.Z - o.Z}
}

func (v Vec3) AddScalar(t float64) Vec3 {
	return Vec3{v.X + t, v.Y + t, v.Z + t}
}

func (v Vec3) SubtractScalar(t float64) Vec3 {
	return Vec3{v.X - t, v.Y - t, v.Z - t}
}

func (v Vec3) MultiplyScalar(t float64) Vec3 {
	return Vec3{v.X * t, v.Y * t, v.Z * t}
}

func (v Vec3) DivideScalar(t float64) Vec3 {
	return Vec3{v.X / t, v.Y / t, v.Z / t}
}
