use std::env;

fn main() {
    println!("Please enter your Name");

    let args: Vec<String> = env::args().collect();

    let name = &args[1];

    println!("Name is {}", name);
}
