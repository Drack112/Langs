fn main() {
    let employee: (u32, &str, bool) = (1, "John", false);

    println!("{}", employee.0);
    println!("{}", employee.1);
    println!("{}", employee.2);
}
