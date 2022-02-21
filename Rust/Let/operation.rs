fn main(){
    let (x,y) = (1,2);
    println!("{:?}", (x, y));

    let z = {x + y};
    println!("{:?}", z);

    let z = {
        let x = 2;
        let y = 7;

        x + y
    };

    println!("{:?}", z);
}