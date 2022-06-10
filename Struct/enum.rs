#[derive(Debug)]
enum GenderCategory {
   Male,Female
}

#[derive(Debug)]
struct Person {
   name:String,
   gender:GenderCategory
}

fn main() {
   let p1 = Person {
      name:String::from("Mohtashim"),
      gender:GenderCategory::Male
   };
   let p2 = Person {
      name:String::from("Amy"),
      gender:GenderCategory::Female
   };
   println!("{:?}",p1);
   println!("{:?}",p2);
}
