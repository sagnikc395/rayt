package main

import (
	"fmt"
	"log"

	"github.com/sagnikc395/rayt/color"
	"github.com/sagnikc395/rayt/ray"
	"github.com/sagnikc395/rayt/vec3"
)

func main() {
	//aspect ratio
	aspect_ratio := 16.0 / 9.0
	image_width := 400
	// calc the image height ans ensure that it's atleast 1.
	var image_height = int(float64(image_width) / aspect_ratio)
	if image_height < 1 {
		image_height = 1
	}

	//viewport widths less than 1 are ok since they are real valued.
	//camera
	focal_length := 1.0
	viewport_height := 2.0
	viewport_width := viewport_height * (float64(image_width) / float64(image_height))
	camera_center := vec3.NewVec3(0, 0, 0)

	//calc the vectors across the horizontal and down the vertical viewport edges
	viewport_u := vec3.NewVec3(viewport_width, 0, 0)
	viewport_v := vec3.NewVec3(0, -1*viewport_height, 0)

	//calculate the horizontal and vertical delta vectors from pixel to pixel
	pixel_delta_u := viewport_u.ScaleDown(float64(image_width))
	pixel_delta_v := viewport_v.ScaleDown(float64(image_height))

	//calculate the location of the upper left pixel

	expressionRight := vec3.VecAdd(vec3.VecAdd(vec3.NewVec3(0, 0, focal_length), viewport_u.ScaleDown(2)), viewport_v.ScaleDown(2))
	viewport_upper_left := vec3.VecSub(camera_center, expressionRight)

	expressionRight2 := vec3.VecAdd(pixel_delta_u, pixel_delta_v).ScaleUp(0.5)
	pixel00_loc := vec3.VecAdd(viewport_upper_left, expressionRight2)

	//render
	fmt.Printf("P3\n%d %d\n255\n", image_width, image_height)
	for j := 0; j < image_height; j++ {
		log.Printf("\rScanlines remaining: %d ", image_height-j)
		for i := 0; i < image_width; i++ {
			pixel_center := vec3.VecAdd(pixel00_loc, vec3.VecAdd(pixel_delta_u.ScaleUp(float64(i)), pixel_delta_v.ScaleUp(float64(j))))
			ray_direction := vec3.VecSub(pixel_center, camera_center)
			r := ray.NewRay(*camera_center, *ray_direction)
			pixel_color := r.RayColor()
			color.WriteColor(&pixel_color)
		}
	}
	log.Printf("\rDone.                \n")
}
