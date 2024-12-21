use crate::{
    ray::{Point3, Ray},
    vec3::Vec3,
};

#[derive(Debug, Clone, Copy)]

pub struct hit_record {
    pub p: Point3,
    pub normal: Vec3,
    pub t: f64,
}

// virtual in cpp akin to a interface -> so we define it as a trait
pub trait hittable {
    fn hit(&self, r: &Ray, ray_tmin: f64, ray_tmax: f64, recc: &mut hit_record) -> bool;
}
