fn main() {
    let str = String::from("welcome");

    match str.as_str() {
        "welcome" => {
            println!("Matched");
        }
        _ => {
            println!("Not matched");
        }
    }
}
