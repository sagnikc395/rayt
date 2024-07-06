package main

import (
	"fmt"
	"log"

	"github.com/sagnikc395/rayt/color"
	"github.com/sagnikc395/rayt/vec3"
)

const (
	//IMAGE_WIDTH  = 256
	IMAGE_WIDTH = 400

// IMAGE_HEIGHT = 256
)

func ray_color()

func main() {
	//aspect ratio
	ASPECT_RATIO := 16.0 / 9.0
	// calc the image height ans ensure that it's atleast 1.
	var IMAGE_HEIGHT = int(IMAGE_WIDTH / ASPECT_RATIO)
	if IMAGE_HEIGHT < 1 {
		IMAGE_HEIGHT = 1
	} else {
		IMAGE_HEIGHT = IMAGE_HEIGHT
	}

	//viewport widths less than 1 are ok since they are real valued.
	VIEWPORT_HEIGHT := 2.0
	VIEWPORT_WIDTH := VIEWPORT_HEIGHT * (float64(IMAGE_WIDTH) / IMAGE_WIDTH)

	//camera
	focal_length := 1.0
	viewport_height := 2.0
	viewport_width := viewport_height * (float64(IMAGE_WIDTH) / float64(IMAGE_HEIGHT))
	camera_center := vec3.NewVec3(0, 0, 0)

	//calc the vectors across the horizontal and down the vertical viewport edges
	viewport_u := vec3.NewVec3(viewport_width, 0, 0)
	viewport_v := vec3.NewVec3(0, -1*viewport_height, 0)

	//calculate the horizontal and vertical delta vectors from pixel to pixel
	pixel_delta_u := viewport_u / IMAGE_WIDTH
	pixel_delta_v := viewport_v / IMAGE_HEIGHT

	//render
	fmt.Printf("P3\n%d %d\n255\n", IMAGE_WIDTH, IMAGE_HEIGHT)

	for j := 0; j < IMAGE_HEIGHT; j++ {
		log.Printf("\rScanlines remaining: %d ", IMAGE_HEIGHT-j)
		for i := 0; i < IMAGE_WIDTH; i++ {
			// r := float64(i) / (IMAGE_WIDTH - 1)
			// g := float64(j) / (IMAGE_HEIGHT - 1)

			// b := 0.0

			// ir := int(255.999 * r)
			// ig := int(255.999 * g)
			// ib := int(255.999 * b)

			// fmt.Printf("%d %d %d\n", ir, ig, ib)
			pixel_color := vec3.NewVec3(float64(i)/(IMAGE_WIDTH-1), float64(j)/(IMAGE_HEIGHT-1), 0)
			color.WriteColor(pixel_color)
		}
	}
	log.Printf("\rDone.                \n")
}
