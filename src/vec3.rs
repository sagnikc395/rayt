use std::ops::{AddAssign, Index, MulAssign, Neg};

pub struct Vec3 {
    e: [f64; 3],
}

impl Vec3 {
    //constructors
    // default constructor
    pub fn default() -> Self {
        Self { e: [0.0, 0.0, 0.0] }
    }
    pub fn new(e0: f64, e1: f64, e2: f64) -> Self {
        Self { e: [e0, e1, e2] }
    }

    pub fn x(self) -> f64 {
        //getter for the x coordinate
        return self.e[0];
    }

    pub fn y(self) -> f64 {
        //getter for the y coordinate
        return self.e[1];
    }

    pub fn z(self) -> f64 {
        //getter for the z coordinate
        return self.e[2];
    }

    pub fn mul_equal(&mut self, t: f64) {
        self.e[0] *= t;
        self.e[1] *= t;
        self.e[2] *= t;
    }

    pub fn div_equal(&mut self, t: f64) {
        self.mul_equal(1 as f64 / t);
    }

    pub fn length(self) -> f64 {
        //https://stackoverflow.com/questions/25165627/how-do-i-use-stdnumsqrt-to-compute-the-square-root-of-a-number
        (self.length_squared()).sqrt()
    }

    fn length_squared(self) -> f64 {
        self.e[0] * self.e[0] + self.e[1] * self.e[1] + self.e[2] * self.e[2]
    }
}

//operators can be overload in rust: ref:https://web.mit.edu/rust-lang_v1.25/arch/amd64_ubuntu1404/share/doc/rust/html/book/first-edition/operators-and-overloading.html

impl Neg for Vec3 {
    type Output = Vec3;
    fn neg(self) -> Vec3 {
        Vec3::new(-self.e[0], -self.e[1], -self.e[2])
    }
}

impl Index<usize> for Vec3 {
    type Output = f64;
    fn index(&self, i: usize) -> &f64 {
        &self.e[i]
    }
}

impl AddAssign for Vec3 {
    fn add_assign(&mut self, v: Vec3) {
        self.e[0] += v.e[0];
        self.e[1] += v.e[1];
        self.e[2] += v.e[2];
    }
}
