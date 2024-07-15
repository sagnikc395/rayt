package main

import (
	"fmt"
	"os"
)

func check(e error, s string) {
	if e != nil {
		fmt.Fprintf(os.Stderr, s, e)
		os.Exit(1)
	}
}

func main() {
	//size of image x and y
	nx := 400
	ny := 200

	const color = 255.99

	f, err := os.Create("dist/image.ppm")

	check(err, "Error opening file: %v\n")
	defer f.Close()

	_, err = fmt.Fprintf(f, "P3\n%d %d\n255\n", nx, ny)

	check(err, "Error opening file: %v\n")

	lowerLeft := Vec3{-2.0, -1.0, -1.0}
	horizontal := Vec3{4.0, 0.0, 0.0}
	vertical := Vec3{0.0, 2.0, 0.0}
	origin := Vec3{0.0, 0.0, 0.0}

	//render
	// main loop to write each pixel with r/g/b bvalues
	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {

			u := float64(i) / float64(nx)
			v := float64(j) / float64(ny)

			position := horizontal.MultiplyScalar(u).Add(vertical.MultiplyScalar(v))

			//direction = lowerleft + (u *horizontal) + (v * vertical)

			direction := lowerLeft.Add(position)

			rgb := Ray{origin, direction}.Color()

			//get intensity of colors
			ir := int(color * rgb.X)
			ig := int(color * rgb.Y)
			ib := int(color * rgb.Z)

			_, err = fmt.Fprintf(f, "%d %d %d\n", ir, ig, ib)
			check(err, "Error writting to file: %v\n")
		}
	}
}
