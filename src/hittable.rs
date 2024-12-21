use crate::{
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
    fn hit(&self, r: &Ray, ray_tmin: f64, ray_tmax: f64, recc: &mut HitRecord) -> bool;
}

impl HitRecord {
    pub fn set_face_normal(self, r: &Ray, outward_normal: &Vec3) {
        //setting the hit record normal vector
        //outward normal is assumed to have unit length
        let front_face = dot(r.direction(), *outward_normal) < 0.0;
        let mut normal = Vec3::new(0.0, 0.0, 0.0);
        if front_face {
            normal = *outward_normal;
        } else {
            normal = -*outward_normal;
        }
    }
}
