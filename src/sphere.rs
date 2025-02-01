use crate::{
    hittable::{HitRecord, Hittable},
    interval::Interval,
    ray::Point3,
    vec3::dot,
};

#[derive(Debug, Clone, Copy)]
pub struct Sphere {
    center: Point3,
    radius: f64,
}

impl Sphere {
    pub fn new(center: Point3, radius: f64) -> Self {
        Self { center, radius }
    }
}

impl Hittable for Sphere {
    // fn hit(&self, r: &crate::ray::Ray, ray_tmin: f64, ray_tmax: f64, recc: &mut HitRecord) -> bool {
    //     let oc = self.center - r.origin();
    //     let a = r.direction().length_squared();
    //     let h = dot(r.direction(), oc);
    //     let c = oc.length_squared() - (self.radius * self.radius);

    //     let discriminant = h * h - a * c;
    //     if discriminant < 0.0 {
    //         return false;
    //     }

    //     let sqrtd = f64::sqrt(discriminant);

    //     //finding the nearest root that lies in acceptable range
    //     let mut root = (h - sqrtd) / a;
    //     if root <= ray_tmin || ray_tmax <= root {
    //         root = (h + sqrtd) / a;
    //         if root <= ray_tmin || ray_tmax <= root {
    //             return false;
    //         }
    //     }

    //     recc.t = root;
    //     recc.p = r.at(recc.t);
    //     let outward_normal = (recc.p - self.center) / self.radius;
    //     recc.set_face_normal(r, &outward_normal);
    //     //recc.normal = (recc.p - self.center) / self.radius;
    //     true
    // }
    fn hit(&self, r: &crate::ray::Ray, ray_t: Interval, rec: &mut HitRecord) -> bool {
        let oc = self.center - r.origin();
        let a = r.direction().length_squared();
        let h = dot(r.direction(), oc);
        let c = oc.length_squared() - (self.radius * self.radius);

        let discriminant = h * h - a * c;
        if discriminant < 0.0 {
            return false;
        }

        let sqrtd = f64::sqrt(discriminant);
        //find the nearest root that lies in acceptable range

        let mut root = (h - sqrtd) / a;
        if !ray_t.surrounds(root) {
            root = (h + sqrtd) / a;
            if !ray_t.surrounds(root) {
                return false;
            }
        }
        rec.t = root;
        rec.p = r.at(rec.t);
        let outward_normal = (rec.p - self.center) / self.radius;
        rec.set_face_normal(r, &outward_normal);
        //recc.normal = (recc.p - self.center) / self.radius;
        true
    }
}
