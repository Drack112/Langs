fn main() {
    let employee: (u32, &str, bool) = (1, "John", false);
    let employee1 = (1, "John", false);

    println!("{:?}", employee); // normal print
    println!("{:#?}", employee1); // pretty print
}
