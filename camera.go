package main

type Camera struct {
	lowerLeft  Vec3
	horizontal Vec3
	vertical   Vec3
	origin     Vec3
}

func NewCamera() *Camera {
	c := new(Camera)

	c.lowerLeft = Vec3{-2.0, -1.0, -1.0}
	c.horizontal = Vec3{4.0, 0., 0.0}
	c.vertical = Vec3{0.0, 2.0, 0.0}
	c.origin = Vec3{0.0, 0.0, 0.0}

	return c
}

func (c *Camera) RayAt(u float64, v float64) Ray {
	position := c.position(u, v)
	direction := c.direction(position)

	return Ray{c.origin, direction}
}

func (c *Camera) position(u float64, v float64) Vec3 {
	horizontal := c.horizontal.MultiplyScalar(u)
	vertical := c.vertical.MultiplyScalar(v)

	return horizontal.Add(vertical)
}

func (c *Camera) direction(position Vec3) Vec3 {
	return c.lowerLeft.Add(position)
}
