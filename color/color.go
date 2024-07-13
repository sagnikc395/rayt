package color

import (
	"fmt"
	"math"
)

// color is a vec3 representation
type Color interface {
	X() float64
	Y() float64
	Z() float64
}

func WriteColor(pixel_color Color) {

	r := pixel_color.X()
	g := pixel_color.Y()
	b := pixel_color.Z()
	// Ensure the color values are in the correct range
	r = math.Max(0, math.Min(1, r))
	g = math.Max(0, math.Min(1, g))
	b = math.Max(0, math.Min(1, b))

	//translation of [0,1] component values to the byte range [0,255]
	rbyte := int(255.999 * r)
	gbyte := int(255.999 * g)
	bbyte := int(255.999 * b)

	//pixel color component
	fmt.Printf("%d %d %d\n", rbyte, gbyte, bbyte)
}
