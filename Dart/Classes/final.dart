class X {
  late final name;
  static const int age = 10;

  X(this.name);
}

void main() {
  var x = X("Jack");
  print(x.name);

  print(X.age);
}
