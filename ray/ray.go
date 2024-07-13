package ray

import "github.com/sagnikc395/rayt/vec3"

type Point3 = vec3.Vec3
type Ray struct {
	Orig Point3
	Dir  vec3.Vec3
}

func NewRay(origin Point3, direction vec3.Vec3) *Ray {
	return &Ray{
		Orig: origin,
		Dir:  direction,
	}
}

func (r *Ray) Origin() Point3 {
	return r.Orig
}

func (r *Ray) Direction() *vec3.Vec3 {
	return &r.Dir
}

func (r *Ray) At(t float64) Point3 {
	return Point3{E: vec3.VecAdd(&r.Orig, r.Dir.ScaleUp(t)).E}
}

func (r *Ray) RayColor() vec3.Vec3 {
	if r.HitSphere(vec3.NewVec3(0, 0, -1), 0.5) {
		return *vec3.NewVec3(1, 0, 0)
	}
	unit_direction := r.Dir.UnitVector()
	a := 0.5 * (unit_direction.Y() + 1.0)

	exp1 := vec3.NewVec3(1.0, 1.0, 1.0)
	exp2 := vec3.NewVec3(0.5, 0.7, 1.0)

	return *vec3.VecAdd(exp1.ScaleUp(1.0-a), exp2.ScaleUp(a))
}

func (r *Ray) HitSphere(center *vec3.Vec3, radius float64) bool {
	oc := vec3.VecSub(center, &r.Orig)
	a := vec3.DotProduct(&r.Dir, &r.Dir)
	b := -2.0 * vec3.DotProduct(&r.Dir, oc)
	c := vec3.DotProduct(oc, oc) - (radius * radius)
	discriminant := b*b - (4 * a * c)
	return discriminant >= 0
}
