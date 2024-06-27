package main

import (
	"fmt"
	"log"

	"github.com/sagnikc395/rayt/color"
	"github.com/sagnikc395/rayt/vec3"
)

const (
	IMAGE_WIDTH  = 256
	IMAGE_HEIGHT = 256
)

func main() {
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
