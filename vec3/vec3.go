package vec3

import (
	"fmt"
	"math"
)

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

func (v *Vec3) ScaleUp(t float64) *Vec3 {
	v.E[0] *= t
	v.E[1] *= t
	v.E[2] *= t
	return v
}

func (v *Vec3) ScaleDown(t float64) *Vec3 {
	return v.ScaleUp(1 / t)
}

func (v *Vec3) Length() float64 {
	return math.Sqrt(v.LengthSquared())
}

func (v *Vec3) LengthSquared() float64 {
	return v.E[0]*v.E[0] + v.E[1]*v.E[1] + v.E[2]*v.E[2]
}

// utility functions for vec3
func (v Vec3) String() string {
	return fmt.Sprintf("%f %f %f", v.E[0], v.E[1], v.E[2])
}

//vec3 on vec3 operators

func VecAdd(u *Vec3, v *Vec3) *Vec3 {
	return &Vec3{
		[3]float64{u.E[0] + v.E[0], u.E[1] + v.E[1], u.E[2] + v.E[2]},
	}
}

func VecSub(u *Vec3, v *Vec3) *Vec3 {
	return &Vec3{
		[3]float64{u.E[0] - v.E[0], u.E[1] - v.E[1], u.E[2] - v.E[2]},
	}
}

func VecMult(u *Vec3, v *Vec3) *Vec3 {
	return &Vec3{
		[3]float64{u.E[0] * v.E[0], u.E[1] * v.E[1], u.E[2] * v.E[2]},
	}
}

func ScalarProduct(t float64, v *Vec3) *Vec3 {
	return &Vec3{
		[3]float64{t * v.E[0], t * v.E[1], t * v.E[2]},
	}
}

func ScalarDiv(t float64, v *Vec3) *Vec3 {
	return ScalarProduct(1/t, v)
}

func DotProduct(u *Vec3, v *Vec3) float64 {
	return u.E[0]*v.E[0] + u.E[1]*v.E[1] + u.E[2]*v.E[2]
}

func CrossProduct(u *Vec3, v *Vec3) *Vec3 {
	return &Vec3{
		[3]float64{
			u.E[1]*v.E[2] - u.E[2]*v.E[1],
			u.E[2]*v.E[0] - u.E[0]*v.E[2],
			u.E[0]*v.E[1] - u.E[1]*v.E[0],
		},
	}
}

func (v *Vec3) UnitVector() *Vec3 {
	return v.ScaleDown(v.Length())
}
