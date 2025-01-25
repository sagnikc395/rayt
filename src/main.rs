use crate::color::Color;
use color::write_color;
use log::info;

mod color;
mod vec3;

fn main() {
    //generate a PPM image format
    let image_width: i32 = 256;
    let image_height: i32 = 256;

    //render
    print!("P3\n{image_width} {image_height}\n255\n");

    for j in 0..image_height {
        let scalines_remaining = image_height - j;
        info!("\rScanlines remaining: {scalines_remaining} ");
        for i in 0..image_width {
            let pixel_color = Color::new(
                i as f64 / (image_width - 1) as f64,
                j as f64 / (image_height - 1) as f64,
                0.0,
            );
            write_color(&pixel_color);
        }
    }
    info!("\rDone.                 \n");
}
