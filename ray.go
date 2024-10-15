package main

type Ray struct {
	orig Point3
	dir  Vec3
}

func NewRay(orig Point3, dir Vec3) *Ray {
	return &Ray{
		orig: orig,
		dir:  dir,
	}
}

func (r *Ray) getOrigin() Point3 {
	return r.orig
}

func (r *Ray) getDirection() Vec3 {
	return r.dir
}

func (r *Ray) at(t float64) Point3 {
	scaledDir := vecMultScalar(t, r.dir)
	orig := r.orig.Vec3
	result := vecAdd(orig, *scaledDir)
	return Point3{*result}
}
