use crate::{
    color::Color,
    vec3::{unit_vector, Vec3},
};

pub type Point3 = Vec3;

#[derive(Debug, Clone, Copy)]
pub struct Ray {
    orig: Point3,
    dir: Vec3,
}

impl Ray {
    //constructor with parameters
    pub fn new(origin: Point3, direction: Vec3) -> Self {
        Self {
            orig: origin,
            dir: direction,
        }
    }

    //getter for origin and direction

    pub fn origin(&self) -> &Point3 {
        &self.orig
    }

    pub fn direction(&self) -> Vec3 {
        self.dir
    }

    //calc point at param t
    pub fn at(self, t: f64) -> Point3 {
        self.orig + self.dir * t
    }
}

pub fn ray_color(r: &Ray) -> Color {
    let unit_direction = unit_vector(r.direction());
    let a = 0.5 * (unit_direction.y() + 1.0);
    //return Color::new(0.0, 0.0, 0.0);
    return (1.0 - a) * Color::new(1.0, 1.0, 1.0) + a * Color::new(0.5, 0.7, 1.0);
}
