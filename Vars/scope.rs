fn main(){
    let age: u32 = 24;

    {
        let id = 2;
        println!("age - {} | id - {}", age, id);
    }
}
