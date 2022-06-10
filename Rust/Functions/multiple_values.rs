fn main(){
    let employee: (u32, &str, bool, u32) = (1, "João", false, 5000);

    let emp = get_employee(employee);

    println!("{:#?}", emp);

}

fn get_employee(emp: (u32, &str, bool, u32)) -> (u32, &str, bool, u32){
    return emp;
}
