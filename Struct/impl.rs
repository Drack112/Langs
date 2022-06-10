struct Retangulo {
    width: u32, height: u32
}

impl Retangulo {
    fn area(&self)-> u32 {
        self.width * self.height
    }
}

fn main() {
    let small = Retangulo {
       width:10,
       height:20
    };
    println!("width is {} height is {} area of Rectangle
    is {}",small.width,small.height,small.area());
 }
