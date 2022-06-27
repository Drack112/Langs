class Vehicle {
  late String model;
  late int year;

  Vehicle(this.model, this.year) {
    print(this.model);
    print(this.year);
  }

  void show() {
    print(model);
    print(year);
  }
}

class Car extends Vehicle {
  late double price;

  Car(String model, int year, this.price) : super(model, year);

  void show() {
    super.show();
    print(this.price);
  }
}

void main() {
  var car1 = Car("Nissan GTR", 2018, 313314);
  car1.show();
}
