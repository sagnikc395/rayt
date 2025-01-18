use std::rc::Rc;

use hittable::{HitRecord, Hittable};
use hittable_list::HittableList;
use log::info;

//use vec3
mod vec3;
//use color
mod color;
//use Ray
mod ray;
//hittable
mod hittable;
// sphere
mod sphere;
//list of hittable objects
mod hittable_list;
//list of common constants
mod rayt_consts;

use color::{write_color, Color};
use ray::{Point3, Ray};
use rayt_consts::infinity;
use sphere::Sphere;
use vec3::{unit_vector, Vec3};

pub fn ray_color(r: &Ray, world: &dyn Hittable) -> Color {
    let mut rec = HitRecord::new();

    if world.hit(r, 0.0, infinity, &mut rec) {
        return Vec3::new(
            0.5 * (rec.normal.x() + 1.0),
            0.5 * (rec.normal.y() + 1.0),
            0.5 * (rec.normal.z() + 1.0),
        );
    }

    let unit_direction: Vec3 = unit_vector(r.direction());
    let a = 0.5 * (unit_direction.y() + 1.0);
    let white = Vec3::new(1.0, 1.0, 1.0);
    let light_blue = Vec3::new(0.5, 0.7, 1.0);
    white * (1.0 - a) + light_blue * a
}

fn main() {
    //image
    let aspect_ratio = 16.0 / 9.0;
    let image_width = 400;

    //image height now calculate from aspect ratio and maintining that it's atleast 1
    let mut image_height = image_width / aspect_ratio as i32;
    //reset if 0, then set to 1
    if image_height < 0 {
        image_height = 1
    } else {
        image_height = image_height
    }

    //world
    let mut world = HittableList::new();

    // Add spheres to the world (corrected syntax)
    world.add(Rc::new(Sphere::new(Point3::new(0.0, 0.0, -1.0), 0.5)));

    world.add(Rc::new(Sphere::new(Point3::new(0.0, -100.5, -1.0), 100.0)));
    //setup camera
    let focal_length = 1.0;
    let viewport_height = 2.0;
    let viewport_width = viewport_height * (image_width / image_height) as f64;
    let camera_center = Point3::new(0.0, 0.0, 0.0);

    //calc the vectors across the horizontal and down the vertical viewport edges
    let viewport_u = Vec3::new(viewport_width, 0.0, 0.0);
    let viewport_v = Vec3::new(0.0, -viewport_height, 0.0);

    //calc the horizontal and vertical delta vectors from pixel to pixel
    let pixel_delta_u = viewport_u / image_width as f64;
    let pixel_delta_v = viewport_v / image_height as f64;

    //calc the location of upper left pixel
    let viewport_upper_left =
        camera_center - Vec3::new(0.0, 0.0, focal_length) - viewport_u / 2.0 - viewport_v / 2.0;

    let pixel00_loc = viewport_upper_left + 0.5 * (pixel_delta_u + pixel_delta_v);
    //render

    print!("P3\n{} {}\n255\n", image_width, image_height);
    for j in 0..image_height {
        info!(target:"rayt_events","\rScanlines remaining: {}", image_height - j);
        for i in 0..image_width {
            let pixel_center =
                pixel00_loc + (i as f64 * pixel_delta_u) + (j as f64 * pixel_delta_v);
            let ray_direction = pixel_center - camera_center;
            let ray = Ray::new(camera_center, ray_direction);
            let pixel_color: Color = ray_color(&ray, &mut world);
            write_color(pixel_color);
        }
    }
    info!(target:"rayt_events","\rDone.             \n");
}
