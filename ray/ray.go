package ray

import "github.com/sagnikc395/rayt/vec3"

type Point3 = vec3.Vec3
type Ray struct {
	Orig Point3
	Dir  vec3.Vec3
}

func NewRay(origin Point3, direction vec3.Vec3) Ray {
	return Ray{
		Orig: origin,
		Dir:  direction,
	}
}

func (r *Ray) Origin() Point3 {
	return r.Orig
}

func (r *Ray) Direction() vec3.Vec3 {
	return r.Dir
}

func (r *Ray) At(t float64) Point3 {
	return Point3{E: vec3.VecAdd(&r.Orig, r.Dir.ScaleUp(t)).E}
}
