package main

import "fmt"

type Color = Vec3

func writeColor(pixelColor Color) {
	r := pixelColor.x()
	g := pixelColor.y()
	b := pixelColor.z()

	//translating the [0,1] component values to the byte range [0,255].

	rbyte := int(255.999 * r)
	gbyte := int(255.999 * g)
	bbyte := int(255.999 * b)

	//write out pixel color components
	fmt.Printf("%d %d %d\n", rbyte, gbyte, bbyte)
}
