use crate::{
    color::Color,
    vec3::{dot, unit_vector, Vec3},
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

    pub fn origin(&self) -> Point3 {
        self.orig
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
    // if check_hit_sphere(Point3::new(0.0, 0.0, -1.0), 0.5, r) {
    //     return Color::new(1.0, 0.0, 0.0);
    // }

    let t = hit_sphere(Point3::new(0.0, 0.0, -1.0), 0.5, r);
    if t > 0.0 {
        let n = unit_vector(r.at(t) - Vec3::new(0.0, 0.0, -1.0));
        return 0.5 * Color::new(n.x() + 1.0, n.y() + 1.0, n.z() + 1.0);
    }

    let unit_direction = unit_vector(r.direction());
    let a = 0.5 * (unit_direction.y() + 1.0);
    //return Color::new(0.0, 0.0, 0.0);
    (1.0 - a) * Color::new(1.0, 1.0, 1.0) + a * Color::new(0.5, 0.7, 1.0)
}

pub fn check_hit_sphere(center: Point3, radius: f64, r: &Ray) -> bool {
    // check if it hits the sphere or not
    let oc = center - r.origin();
    let a = dot(r.direction(), r.direction());
    let b = -2.0 * dot(r.direction(), oc);
    let c = dot(oc, oc) - radius * radius;
    let discriminant = b * b - (4.0 * a * c);
    discriminant >= 0.0
}

//although in the text book , they have modified the existing implementation hit_sphere to return closest hit point
// here i have abstracted into a different function

pub fn hit_sphere(center: Point3, radius: f64, r: &Ray) -> f64 {
    let oc = center - r.origin();
    let a = dot(r.direction(), r.direction());
    let b = -2.0 * dot(r.direction(), oc);
    let c = dot(oc, oc) - radius * radius;
    let discriminant = b * b - (4.0 * a * c);

    if discriminant < 0.0 {
        return -1.0;
    } else {
        return (-b - discriminant.sqrt()) / 2.0 * a;
    }
}


