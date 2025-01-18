use crate::hittable::HitRecord;
use crate::hittable::Hittable;
use crate::ray::Ray;
use std::rc::Rc;
use std::vec::Vec;

pub struct HittableList {
    //rc is useuful for shared references
    objects: Vec<Rc<dyn Hittable>>,
}

impl HittableList {
    //create a new empty list
    pub fn new() -> Self {
        Self {
            objects: Vec::new(),
        }
    }

    //create a new list with 1 object
    pub fn from_object(object: Rc<dyn Hittable>) -> Self {
        let mut list = Self::new();
        list.add(object);
        list
    }

    //clear all the objects
    pub fn clear(&mut self) {
        self.objects.clear();
    }

    //add a object to the list
    pub fn add(&mut self, object: Rc<dyn Hittable>) {
        self.objects.push(object);
    }
}

// implementation of the Hittable trait for Hittable list
impl Hittable for HittableList {
    fn hit(&self, r: &Ray, ray_tmin: f64, ray_tmax: f64, rec: &mut HitRecord) -> bool {
        let mut temp_rec = HitRecord::new();
        let mut hit_anything = false;
        let mut closest_so_far = ray_tmax;

        for object in &self.objects {
            if object.hit(r, ray_tmin, closest_so_far, &mut temp_rec) {
                hit_anything = true;
                closest_so_far = temp_rec.t;
                *rec = temp_rec;
            }
        }

        hit_anything
    }
}
