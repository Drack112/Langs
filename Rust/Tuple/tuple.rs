fn main(){
    let a = (1, 1.5, true, 'a');
    let a: (i32, f64, bool, char) = (1, 1.5, true, 'a');

    let mut b = (1, 1.5);
    b.0 = 2;
    b.1 = 3.0;

    println!("{:?}", b); // (2, 3.0)
    println!("{:#?}", b);
}