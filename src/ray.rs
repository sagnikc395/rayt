use crate::vec3::{dot, Vec3};

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
