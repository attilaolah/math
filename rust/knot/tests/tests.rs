use knot::{Direction, Orientation};

#[test]
fn test_direction_from_str() {
    let in_ok = "LROU";
    let exp_ok = vec![
        Direction::Left,
        Direction::Right,
        Direction::Over,
        Direction::Under,
    ];

    let in_bad = "LRX";

    let res_ok = Direction::from_str(in_ok);
    assert!(res_ok.is_ok(), "Parsing '{}' failed.", in_ok);
    assert_eq!(res_ok.unwrap(), exp_ok);

    let res_err = Direction::from_str(in_bad);
    assert!(res_err.is_err(), "Parsing '{}' should fail.", in_bad);
    assert_eq!(
        res_err.unwrap_err().to_string(),
        format!("'X' is not a valid direction")
    );
}

#[test]
fn test_orientation_from_str() {
    let in_ok = "LRUD";
    let exp_ok = vec![
        Orientation::Left,
        Orientation::Right,
        Orientation::Up,
        Orientation::Down,
    ];

    let in_bad = "LRX";

    let res_ok = Orientation::from_str(in_ok);
    assert!(res_ok.is_ok(), "Parsing '{}' failed.", in_ok);
    assert_eq!(res_ok.unwrap(), exp_ok);

    let res_err = Orientation::from_str(in_bad);
    assert!(res_err.is_err(), "Parsing '{}' should fail.", in_bad);
    assert_eq!(
        res_err.unwrap_err().to_string(),
        format!("'X' is not a valid orientation")
    );
}
