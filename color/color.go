package color

import (
	"fmt"

	"github.com/sagnikc395/rayt/vec3"
)

func WriteColor(v *vec3.Vec3) {
	r := v.X()
	g := v.Y()
	b := v.Z()

	//[0,1] component values to the byte range [0,255]
	rbyte := int(255.999 * r)
	gbyte := int(255.999 * g)
	bbyte := int(255.999 * b)

	//pixel color component
	fmt.Printf("%d %d %d\n", rbyte, gbyte, bbyte)
}
