enum WEEKEND {
    SATURDAY,
    SUNDAY,
}
fn main() {
    let weekend = WEEKEND::SUNDAY;
    match weekend {
        WEEKEND::SUNDAY => println!("SUNDAY"),
        WEEKEND::SATURDAY => println!("SATURDAY"),
        _ => println!("NOT WEEKEND"),
    }
}
