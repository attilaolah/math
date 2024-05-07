use knot::Direction;
use knot::validation::{ValidationError, validate};

#[test]
fn test_validate() {
    {
        let err = validate(Direction::from_str("").unwrap());
        assert!(err.is_err());
        assert_eq!(err.unwrap_err(), ValidationError::Empty);
    }
    {
        let err = validate(Direction::from_str("O").unwrap());
        assert!(err.is_err());
        assert_eq!(err.unwrap_err(), ValidationError::NotEven);
    }

    const unknot: &str = "LLLL";
    let in_ok = Direction::from_str(unknot).unwrap();
    assert!(validate(in_ok).is_ok(), "Validating {} failed.", unknot);
}
