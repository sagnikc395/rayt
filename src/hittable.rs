use crate::{
    ray::{Point3, Ray},
    vec3::Vec3,
};

#[derive(Debug, Clone, Copy)]

pub struct hit_record {
    p: Point3,
    normal: Vec3,
    t: f64,
}

// virtual in cpp akin to a interface -> so we define it as a trait
pub trait hittable {
    fn hit(&self, r: &Ray, ray_tmin: f64, ray_tmax: f64, recc: &mut hit_record) -> bool;
}

