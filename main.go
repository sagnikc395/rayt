package main

import (
	"fmt"
	"log"

	"github.com/sagnikc395/rayt/color"
	"github.com/sagnikc395/rayt/ray"
	"github.com/sagnikc395/rayt/vec3"
)

func RayColor(r *ray.Ray) color.Color {
	center := vec3.NewVec3(0, 0, -1)
	if ray.HitSphere(center, 0.5, r) {
		// return red if ray hits the sphere
		return vec3.NewVec3(1, 0, 0)
	}

	// unitDirection := r.Direction().UnitVector()
	// t := 0.5 * (unitDirection.Y() + 1.0)

	// white := vec3.NewVec3(1.0, 1.0, 1.0)
	// blue := vec3.NewVec3(0.5, 0.7, 1.0)

	// color := vec3.VecAdd(
	// 	white.ScaleUp(1.0-t),
	// 	blue.ScaleUp(t),
	// )

	// return color
	return vec3.NewVec3(0, 0, 1)
}

func main() {
	//image

	aspect_ratio := 16.0 / 9.0
	// image_width := 256
	image_width := 400
	// image_height := 256

	//calcualte the image height, and ensure that it's at least 1
	image_height := int(float64(image_width) / aspect_ratio)
	if image_height < 1 {
		image_height = 1
	}

	//camera
	focal_length := 1.0
	viewport_height := 2.0
	viewport_width := viewport_height * (float64(image_width) / float64(image_height))
	camera_center := vec3.NewVec3(0, 0, 0)

	//cauclate the vectors across the horizontal nd down the vertical viewport edges
	viewport_u := vec3.NewVec3(viewport_width, 0, 0)
	viewport_v := vec3.NewVec3(0, -1*viewport_height, 0)

	//calculate the horizontal and vertical delta vecrors from pixel to pixel

	pixel_delta_u := viewport_u.ScaleDown(float64(image_width))
	pixel_delta_v := viewport_v.ScaleDown(float64(image_height))

	//calculate the location of the upper left pixel

	viewport_upper_left := vec3.VecSub(camera_center,
		vec3.VecAdd(vec3.NewVec3(0, 0, focal_length),
			vec3.VecAdd(viewport_v.ScaleDown(2), viewport_u.ScaleDown(2))))

	pixel00_loc := vec3.VecAdd(viewport_upper_left, vec3.VecAdd(pixel_delta_u, pixel_delta_v).ScaleUp(0.5))

	//render them
	fmt.Printf("P3\n%d %d \n255\n", image_width, image_height)

	for j := 0; j < image_height; j++ {
		log.Printf("\rScanlines remaining: %d ", (image_height - j))
		for i := 0; i < image_width; i++ {
			pixel_center := vec3.VecAdd(pixel00_loc,
				vec3.VecAdd(pixel_delta_u.ScaleUp(float64(i)), pixel_delta_v.ScaleUp(float64(j))))

			ray_direction := vec3.VecSub(pixel_center, camera_center)

			r := ray.NewRay(*camera_center, *ray_direction)

			pixel_color := RayColor(r)
			color.WriteColor(pixel_color)

		}
	}
	log.Print("\rDone.                 \n")
}
