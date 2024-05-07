use std::error;
use std::fmt;

use crate::{Direction, Orientation};

#[derive(Debug, PartialEq, Eq)]
pub enum ValidationError {
    Empty,
    NotEven,
}

pub fn validate(dirs: Vec<Direction>) -> Result<(), ValidationError> {
    if dirs.is_empty() {
        return Err(ValidationError::Empty);
    }
    if dirs.len() % 2 != 0 {
        return Err(ValidationError::NotEven);
    }

    let pos = Point::zero();
    let enter = Orientation::Up;

    Ok(())
}

impl fmt::Display for ValidationError {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        write!(f, "{}", 
        match self {
            Self::Empty => "expected more than one direction",
            Self::NotEven => "expected an even number of directions",
        })
    }
}

impl error::Error for ValidationError {}
