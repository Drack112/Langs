fn main(){
    let x: f64 = -20.48;
    let x: i64 = x.floor() as i64;
    println!("{}", x);

    let s: &str = "Hello World";
    let s: String = s.to_uppercase();
    println!("{}", s);
}