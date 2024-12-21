use crate::{hittable::hittable, ray::Point3};

#[derive(Debug, Clone, Copy)]
pub struct sphere {
    center: Point3,
    radius: f64,
}

impl sphere {
    fn new(self, center: Point3, radius: f64) -> Self {
        Self::new(self, center, f64::max(0.0, radius))
    }
}

impl hittable for sphere {
    fn hit(
        &self,
        r: &crate::ray::Ray,
        ray_tmin: f64,
        ray_tmax: f64,
        recc: &mut crate::hittable::hit_record,
    ) -> bool {
        false
    }
}
