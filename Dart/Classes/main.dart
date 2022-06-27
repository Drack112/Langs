class Person {
  late String name;
  late int age;

  void showOutput() {
    print("My name is $name and i have $age years old");
  }
}

void main() {
  Person person1 = Person();

  person1.name = "João Vitor";
  person1.age = 13;
  person1.showOutput();
}
