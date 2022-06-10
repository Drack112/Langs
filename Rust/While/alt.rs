fn main() {
    let mut number = 20;
    let mut sum = 0;
    while {
     sum = sum + number;
        number = number - 1;
        if number === 0 {
            break  ;
        }
    }  {}
    println!("The Sum is {sum}");
}
