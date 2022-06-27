class Person {
  late String name;
  late int age;

  Person(String name, [int age = 18]) {
    this.name = name;
    this.age = age;
  }

  void showOutput() {
    print("My name is $name and i have $age years old");
  }
}

void main() {
  Person person1 = Person("João Vitor", 13);
  person1.showOutput();
}
