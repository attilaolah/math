use std::error;
use std::fmt;

pub mod validation;

#[derive(Debug, Copy, Clone)]
struct Point {
    x: i32,
    y: i32,
}

#[derive(Debug, PartialEq)]
pub enum Direction {
    Over,
    Under,
    Left,
    Right,
}

#[derive(Debug)]
pub struct BadDirection(char);

#[derive(Debug, PartialEq)]
pub enum Orientation {
    Up,
    Down,
    Left,
    Right,
}

#[derive(Debug)]
pub struct BadOrientation(char);

impl Point {
    fn new(x: i32, y: i32) -> Self {
        Self { x, y }
    }
    fn zero() -> Self {
        Self { x: 0, y: 0 }
    }
}

impl Direction {
    pub fn from_str(repr: &str) -> Result<Vec<Self>, BadDirection> {
        let orientations: Result<Vec<Self>, BadDirection> =
            repr.chars().map(|c| Self::try_from(c)).collect();
        orientations
    }
}

impl TryFrom<char> for Direction {
    type Error = BadDirection;

    fn try_from(ch: char) -> Result<Self, Self::Error> {
        match ch {
            'O' => Ok(Direction::Over),
            'U' => Ok(Direction::Under),
            'L' => Ok(Direction::Left),
            'R' => Ok(Direction::Right),
            _ => Err(BadDirection(ch)),
        }
    }
}

impl fmt::Display for BadDirection {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        write!(f, "'{}' is not a valid direction", self.0)
    }
}

impl Orientation {
    /// Returns the new orientation after turning towards a direction.
    pub fn turn(self, toward: Direction) -> Self {
        match toward {
            Direction::Left => match self {
                Self::Up => Self::Left,
                Self::Down => Self::Right,
                Self::Left => Self::Down,
                Self::Right => Self::Up,
            },
            Direction::Right => match self {
                Self::Up => Self::Right,
                Self::Down => Self::Left,
                Self::Left => Self::Up,
                Self::Right => Self::Down,
            },
            // Turning "over" or "under" doesn't change the orientation.
            _ => self,
        }
    }

    pub fn from_str(repr: &str) -> Result<Vec<Self>, BadOrientation> {
        let orientations: Result<Vec<Self>, BadOrientation> =
            repr.chars().map(|c| Self::try_from(c)).collect();
        orientations
    }
}

impl TryFrom<char> for Orientation {
    type Error = BadOrientation;

    fn try_from(ch: char) -> Result<Self, Self::Error> {
        match ch {
            'U' => Ok(Orientation::Up),
            'D' => Ok(Orientation::Down),
            'L' => Ok(Orientation::Left),
            'R' => Ok(Orientation::Right),
            _ => Err(BadOrientation(ch)),
        }
    }
}

impl fmt::Display for BadOrientation {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        write!(f, "'{}' is not a valid orientation", self.0)
    }
}

impl error::Error for BadOrientation {}
