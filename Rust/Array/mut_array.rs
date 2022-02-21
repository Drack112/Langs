fn main(){
    let a = [1,2,3,4,5];
    println!("{:?}", a);

    let b: &[i32] = &a;
    println!("{:?}", b);

    let c = &a[0..2];
    println!("{:?}", c);

    let d = &a[2..];
    println!("{:?}", d);

    let e = &a[..3];
    println!("{:?}", e);

    let f = &a[..];
    println!("{:?}", f);
}