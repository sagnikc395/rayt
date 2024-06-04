package vec3

import "math"

type Vec3 struct {
	E [3]float64
}

func (v *Vec3) NewVec3(e0, e1, e2 float64) *Vec3 {
	return &Vec3{
		[3]float64{e0, e1, e2},
	}
}

func (v *Vec3) X() float64 {
	return v.E[0]
}

func (v *Vec3) Y() float64 {
	return v.E[1]
}

func (v *Vec3) Z() float64 {
	return v.E[2]
}

func (v *Vec3) Negate() *Vec3 {
	return &Vec3{
		[3]float64{-1 * v.E[0], -1 * v.E[1], -1 * v.E[2]},
	}
}
func (v *Vec3) ItemAt(index int) float64 {
	return v.E[index]
}

func (v *Vec3) Increment(v2 Vec3) *Vec3 {
	v.E[0] += v2.E[0]
	v.E[1] += v2.E[1]
	v.E[2] += v2.E[2]
	return v
}

func (v *Vec3) ScalarMult(t float64) *Vec3 {
	v.E[0] *= t
	v.E[1] *= t
	v.E[2] *= t
	return v
}

func (v *Vec3) ScalarDiv(t float64) *Vec3 {
	return v.ScalarMult(1 / t)
}

func (v *Vec3) length() float64 {
	return math.Sqrt(v.lengthSquared())
}

func (v *Vec3) lengthSquared() float64 {
	return v.E[0]*v.E[0] + v.E[1]*v.E[1] + v.E[2]*v.E[2]
}
