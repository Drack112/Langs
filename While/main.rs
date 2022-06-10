fn main() {
    let mut number = 20;
    let mut sum = 0;

    while number >= 1 {
        sum = sum + number;
        number = number - 1;
    }
    println!("The Sum is {sum}");
}
