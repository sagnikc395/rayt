package main

import (
	"fmt"
	"math"
	"math/rand/v2"
	"os"
)

func check(e error, s string) {
	if e != nil {
		fmt.Fprintf(os.Stderr, s, e)
		os.Exit(1)
	}
}

const (
	nx = 400 // x
	ny = 200 // y
	ns = 100 // number of samples for aa
	c  = 255.99
)

var (
	white = Vec3{1.0, 1.0, 1.0}
	blue  = Vec3{0.5, 0.7, 1.0}

	camera = NewCamera()

	sphere = Sphere{Vec3{0, 0, -1}, 0.5}
	floor  = Sphere{Vec3{0, -100.5, -1}, 100}

	world = World{[]Hittable{&sphere, &floor}}
)

// color method
func color(r *Ray, h Hittable) Vec3 {
	hit, record := h.Hit(r, 0.0, math.MaxFloat64)

	if hit {
		return record.Normal.AddScalar(1.0).MultiplyScalar(0.5)
	}

	// make unit vector so y is between -1.0 and 1.0
	unitDirection := r.Direction.Normalize()

	return gradient(&unitDirection)
}

func gradient(v *Vec3) Vec3 {
	// scale t to be between 0.0 and 1.0
	t := 0.5 * (v.Y + 1.0)

	// linear blend: blended_value = (1 - t) * white + t * blue
	return white.MultiplyScalar(1.0 - t).Add(blue.MultiplyScalar(t))
}

func main() {

	f, err := os.Create("dist/image.ppm")

	check(err, "Error opening file: %v\n")
	defer f.Close()

	_, err = fmt.Fprintf(f, "P3\n%d %d\n255\n", nx, ny)

	check(err, "Error opening file: %v\n")

	//render
	// main loop to write each pixel with r/g/b bvalues
	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {

			rgb := Vec3{}

			// sample rays for anti-aliasing
			for s := 0; s < ns; s++ {
				u := (float64(i) + rand.Float64()) / float64(nx)
				v := (float64(j) + rand.Float64()) / float64(ny)

				r := camera.RayAt(u, v)
				color := color(&r, &world)
				rgb = rgb.Add(color)
			}

			// average
			rgb = rgb.DivideScalar(float64(ns))

			// get intensity of colors
			ir := int(c * rgb.X)
			ig := int(c * rgb.Y)
			ib := int(c * rgb.Z)

			_, err = fmt.Fprintf(f, "%d %d %d\n", ir, ig, ib)
			check(err, "Error writing to file: %v\n")
		}
	}
}
