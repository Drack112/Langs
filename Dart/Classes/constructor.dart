class Person {
  late String name;
  late int age;

  Person(String name, [int age = 18]) {
    this.name = name;
    this.age = age;
  }

  Person.guest() {
    name = "Guest";
    age = 18;
  }

  void showOutput() {
    print("My name is $name and i have $age years old");
  }
}

void main() {
  var person2 = Person.guest();
  person2.showOutput();
}
