fn main(){
    let name = "Drack".to_string();
    display_message(name, 32);
}

fn display_message(name: String, salary: u32){
    println!("{} - {}", name, salary);
}
