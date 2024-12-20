//underlying storage

use std::fmt;

#[derive(Debug, Clone, Copy)]
pub struct Vec3 {
    e: [f64; 3],
}
//methods for vec3
impl Vec3 {
    //constructor ... according to conventions
    pub fn new(e0: f64, e1: f64, e2: f64) -> Self {
        //ref: https://stackoverflow.com/questions/32304595/whats-the-difference-between-self-and-self
        Self { e: [e0, e1, e2] }
    }

    //for zero vector
    pub fn zero() -> Self {
        Self { e: [0.0, 0.0, 0.0] }
    }

    pub fn x(&self) -> f64 {
        self.e[0]
    }

    pub fn y(&self) -> f64 {
        self.e[1]
    }

    pub fn z(&self) -> f64 {
        self.e[2]
    }

    //get length
    pub fn length(&self) -> f64 {
        self.length_squared().sqrt()
    }

    pub fn length_squared(&self) -> f64 {
        self.e[0] * self.e[0] + self.e[1] * self.e[1] + self.e[2] * self.e[2]
    }
}

//implement the display
impl fmt::Display for Vec3 {
    //implement the fmt
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "{} {} {}", self.e[0], self.e[1], self.e[2])
    }
}
