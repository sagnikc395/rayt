//underlying storage

use std::{
    fmt,
    ops::{Div, Index, IndexMut, Mul, Neg},
};

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

//cross product
pub fn cross(u: Vec3, v: Vec3) -> Vec3 {
    //return a new instance of the cross product
    Vec3::new(
        u.e[1] * v.e[2] - u.e[2] * v.e[1],
        u.e[2] * v.e[0] - u.e[0] * v.e[2],
        u.e[0] * v.e[1] - u.e[1] * v.e[0],
    )
}
//dot product
pub fn dot(u: Vec3, v: Vec3) -> f64 {
    u.e[0] * v.e[0] + u.e[1] * v.e[1] + u.e[2] * v.e[2]
}

//get the unit vector
pub fn unit_vector(v: Vec3) -> Vec3 {
    v / v.length()
}

//implement the traits for various operators
impl Neg for Vec3 {
    //filled from error description to satisfy
    type Output = Self;
    fn neg(self) -> Self {
        Vec3::new(-1 * self.e[0], -1 * self.e[1], -1 * self.e[2])
    }
}
impl Div<f64> for Vec3 {
    type Output = Self;

    fn div(self, t: f64) -> Self {
        self * (1.0 / t)
    }
}
impl Mul<f64> for Vec3 {
    type Output = Self;

    fn mul(self, t: f64) -> Self {
        Vec3::new(self.e[0] * t, self.e[1] * t, self.e[2] * t)
    }
}

//implement indexing operator for Vec3
impl Index<usize> for Vec3 {
    type Output = Self;

    fn index(&self, i: usize) -> &Self::Output {
        &self.e[i]
    }
}

impl IndexMut<usize> for Vec3 {
    fn index_mut(&mut self, i: usize) -> &mut Self::Output {
        &mut self.e[i]
    }
}
