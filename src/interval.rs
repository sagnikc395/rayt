use core::f64;

pub struct Interval {
    max: f64,
    min: f64,
}

impl Interval {
    //static constants
    pub const EMPTY: Self = Self {
        min: f64::INFINITY,
        max: f64::NEG_INFINITY,
    };
    pub const UNIVERSE: Self = Self {
        min: f64::NEG_INFINITY,
        max: f64::INFINITY,
    };
    // Default constructor for empty interval
    pub fn new_empty() -> Self {
        Self {
            min: f64::INFINITY,
            max: f64::NEG_INFINITY,
        }
    }

    pub fn new(min: f64, max: f64) -> Self {
        Self { min, max }
    }

    pub fn size(&self) -> f64 {
        return self.max - self.min;
    }

    pub fn contains(&self, x: f64) -> bool {
        return self.min <= x && x <= self.max;
    }
    pub fn surrounds(&self, x: f64) -> bool {
        return self.min < x && x < self.max;
    }
}

impl Default for Interval {
    fn default() -> Self {
        Self::new_empty()
    }
}
