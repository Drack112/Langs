fn main(){
    let mut ctr = 0;

    while ctr <= 10 {
        ctr = ctr + 1;
        if ctr == 5 {
            break;
        }
        println!("Inside loop {}", ctr);
    }
    println!("Out of while loop");
}
