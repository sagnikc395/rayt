//some usful constants available globally
// akin to the rtweekend.h common header file

//constants
pub const infinity: f64 = f64::INFINITY;
pub const pi: f64 = 3.1415926535897932385;

//utility function
pub fn degrees_to_radians(degrees: f64) -> f64 {
    degrees * (pi / 180.0)
}
