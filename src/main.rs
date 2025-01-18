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
use ray::{ray_color, Point3, Ray};
use vec3::Vec3;

// const IMAGE_HEIGHT: u16 = 256;
// const IMAGE_WIDTH: u16 = 256;

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
            // let r = i as f64 / (IMAGE_WIDTH - 1) as f64;
            // let g = j as f64 / (IMAGE_HEIGHT - 1) as f64;
            // let b = 0.0 as f64;

            // let ir = (255.999 * r) as i32;
            // let ig = (255.999 * g) as i32;
            // let ib = (255.999 * b) as i32;

            // print!("{} {} {}\n", ir, ig, ib);
            // let color = Color::new(
            //     i as f64 / (image_width - 1) as f64,
            //     j as f64 / (image_height - 1) as f64,
            //     0.0,
            // );
            // write_color(color);

            let pixel_center =
                pixel00_loc + (i as f64 * pixel_delta_u) + (j as f64 * pixel_delta_v);

            let ray_direction = pixel_center - camera_center;

            let ray = Ray::new(camera_center, ray_direction);

            let pixel_color: Color = ray_color(&ray);

            write_color(pixel_color);
        }
    }

    info!(target:"rayt_events","\rDone.             \n");
}
