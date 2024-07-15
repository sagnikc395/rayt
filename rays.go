package main

// model the ray struct
type Ray struct {
	Origin, Direction Vec3
}

// point
func (r *Ray) Point(t float64) Vec3 {
	b := r.Direction.MultiplyScalar(t)
	a := r.Origin
	return a.Add(b)
}

// func (r Ray) HitSphere(s Sphere) bool {
// 	oc := r.Origin.Subtract(s.Center)
// 	a := r.Direction.Dot(r.Direction)
// 	b := -2.0 * oc.Dot(r.Direction)
// 	c := oc.Dot(oc) - s.Radius*s.Radius
// 	discriminant := b*b - 4*a*c
// 	return discriminant >= 0
// }

// func (r Ray) Color() Vec3 {
// 	sphere := Sphere{Center: Vec3{0, 0, -1}, Radius: 0.5}

// 	if r.HitSphere(sphere) {
// 		return Vec3{1.0, 0.0, 0.0} // red
// 	}

// 	unitDir := r.Direction.Normalize()

// 	//scale t to 0.0 and 1.0
// 	t := 0.5 * (unitDir.Y + 1.0)
// 	// linear blend
// 	// blended_value = (1 - t) * white + t * blue
// 	white := Vec3{1.0, 1.0, 1.0}
// 	blue := Vec3{0.5, 0.7, 1.0}

// 	return white.MultiplyScalar(1.0 - t).Add(blue.MultiplyScalar(t))
// }
