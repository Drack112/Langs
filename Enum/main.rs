#[derive(Debug)]

enum Status {
    ACTIVE,
    INACTIVE,
}

fn main() {
    let active = Status::ACTIVE;
    let inactive = Status::INACTIVE;
    println!("{:?}", active);
    println!("{:?}", inactive);
}
