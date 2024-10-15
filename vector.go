package main

import (
	"fmt"
	"math"
)

type Vec3 struct {
	e [3]float64
}

func NewVec3(e0, e1, e2 float64) *Vec3 {
	return &Vec3{
		e: [3]float64{e0, e1, e2},
	}
}

func (v *Vec3) x() float64 {
	return v.e[0]
}

func (v *Vec3) y() float64 {
	return v.e[1]
}

func (v *Vec3) z() float64 {
	return v.e[2]
}

func (v *Vec3) negate() Vec3 {
	return *NewVec3(-1*v.e[0], -1*v.e[1], -1*v.e[2])
}

// operator[]
func (v *Vec3) itemAt(index int) float64 {
	return v.e[index]
}

// operator+=
func (v *Vec3) plusEquals(v2 Vec3) {
	v.e[0] += v2.e[0]
	v.e[1] += v2.e[1]
	v.e[2] += v2.e[2]
}

// operator*=
func (v *Vec3) multiplyEquals(t float64) {
	v.e[0] *= t
	v.e[1] *= t
	v.e[2] *= t
}

// operator/=
func (v *Vec3) divideEqual(t float64) {
	v.multiplyEquals(1 / t)
}

func (v *Vec3) length() float64 {
	return math.Sqrt(v.lengthSquared())
}

func (v *Vec3) lengthSquared() float64 {
	return v.e[0]*v.e[0] + v.e[1]*v.e[1] + v.e[2]*v.e[2]
}

// type Point3 is a type alias for vec3, but useful for geometric clarity
type Point3 = Vec3

// vec utility functions
// & operator<<
func (v *Vec3) printVector() {
	fmt.Printf("%f %f %f\n", v.e[0], v.e[1], v.e[2])
}

// operator+
func vecAdd(v1, v2 Vec3) *Vec3 {
	return NewVec3(v1.e[0]+v2.e[0], v1.e[1]+v2.e[0], v1.e[2]+v2.e[2])
}

// operator-
func vecSub(v1, v2 Vec3) *Vec3 {
	return NewVec3(v1.e[0]-v2.e[0], v1.e[1]-v2.e[1], v1.e[2]-v2.e[2])
}

// operator*
func vecMult(v1, v2 Vec3) *Vec3 {
	return NewVec3(v1.e[0]*v2.e[0], v1.e[1]*v2.e[1], v1.e[2]*v2.e[2])
}

// operator* (scalar)
func vecMultScalar(t float64, v Vec3) *Vec3 {
	return NewVec3(t*v.e[0], t*v.e[1], t*v.e[2])
}

// operator/
func vecMultDivide(t float64, v Vec3) *Vec3 {
	return vecMultScalar(1/t, v)
}

// dot product
func dotProduct(v1, v2 Vec3) float64 {
	return v1.e[0]*v2.e[0] + v1.e[1]*v2.e[1] + v1.e[2]*v2.e[2]
}

// cross product
func crossProduct(v1, v2 Vec3) *Vec3 {
	return NewVec3(v1.e[1]*v2.e[2]-v1.e[2]*v2.e[1],
		v1.e[2]*v2.e[0]-v1.e[0]*v2.e[2],
		v1.e[0]*v2.e[1]-v1.e[1]*v2.e[0])
}

// unit vector
func unitVector(v Vec3) *Vec3 {
	return vecMultDivide(v.length(), v)
}
