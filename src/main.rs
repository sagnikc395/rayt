use log::info;

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
            let r = (i) as f64 / (image_width - 1) as f64;
            let g = (j) as f64 / (image_height - 1) as f64;
            let b = 0.0;

            let ir = (255.999 * r) as i32;
            let ig = (255.999 * g) as i32;
            let ib = (255.999 * b) as i32;

            print!("{ir} {ig} {ib}\n");
        }
    }
    info!("\rDone.                 \n");
}
