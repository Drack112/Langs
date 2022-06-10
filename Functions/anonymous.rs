fn main(){
    let display = |name: String | -> String {
        println!("{}", name);
    };

    let name = "poof".to_string();

    display(name);
}
