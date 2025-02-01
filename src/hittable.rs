use crate::{
    interval::Interval,
    ray::{Point3, Ray},
    vec3::{dot, Vec3},
};

#[derive(Debug, Clone, Copy)]

pub struct HitRecord {
    pub p: Point3,
    pub normal: Vec3,
    pub t: f64,
    pub front_face: bool,
}

// virtual in cpp akin to a interface -> so we define it as a trait
pub trait Hittable {
    // fn hit(&self, r: &Ray, ray_tmin: f64, ray_tmax: f64, recc: &mut HitRecord) -> bool;
    fn hit(&self, r: &Ray, ray_t: Interval, rec: &mut HitRecord) -> bool;
}

impl HitRecord {
    pub fn new() -> Self {
        Self {
            p: Point3::new(0.0, 0.0, 0.0),
            normal: Vec3::new(0.0, 0.0, 0.0),
            t: 0.0,
            front_face: false,
        }
    }
    pub fn set_face_normal(&mut self, r: &Ray, outward_normal: &Vec3) {
        //setting the hit record normal vector
        //outward normal is assumed to have unit length
        self.front_face = dot(r.direction(), *outward_normal) < 0.0;
        self.normal = if self.front_face {
            *outward_normal
        } else {
            -*outward_normal
        };
    }
}
