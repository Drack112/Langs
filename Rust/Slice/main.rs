fn main(){
    let a: [i32; 5] = [1,2,3,4,5];
    println!("{:?}", a);

    let b: &[i32] = &a;
    println!("{:?}", b);

    let c = &a[0..2];
    println!("{:?}", c);

    let d = &a[..];
    println!("{:?}", d);

    let e = &a[..3];
    println!("{:?}", e);

    let f = &a[4..];
    println!("{:?}", f);
}