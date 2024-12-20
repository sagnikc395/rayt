use crate::vec3::Vec3;

pub struct Color {
    pub value: Vec3,
}
impl Color {
    pub fn new(r: f64, g: f64, b: f64) -> Self {
        Color {
            value: Vec3::new(r, g, b),
        }
    }

    pub fn r(&self) -> f64 {
        self.value[0]
    }
    pub fn g(&self) -> f64 {
        self.value[1]
    }
    pub fn b(&self) -> f64 {
        self.value[2]
    }
}

pub fn write_color(pixel_color: Color) {
    let r = pixel_color.r();
    let g = pixel_color.g();
    let b = pixel_color.b();

    //[0,1] component values to byte range
    let rbyte = (255.999 * r) as i64;
    let gbyte = (255.999 * g) as i64;
    let bbyte = (255.999 * b) as i64;

    //pixel color components
    print!("{} {} {}\n", rbyte, gbyte, bbyte);
}
