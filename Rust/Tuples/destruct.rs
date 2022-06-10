fn main() {
    let employee: (u32, &str, bool) = (1, "John", false);

    let (id, name, active) = employee;
    println!("{}", id);
    println!("{}", name);
    println!("{}", active);
}
