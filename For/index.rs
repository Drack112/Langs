fn main() {
    let numbers: &[i64] = &vec![1, 2, 3, 4, 5];

    for (i, item) in numbers.iter().enumerate() {
        println!("value: - {} - Index: {}", item,i);
    }
}
