class User {
  name: string;
  age: number;

  constructor(name: string, age: number) {
    this.name = name;
    this.age = age;
    console.log(`${this.name}, ${this.age}`);
  }
}

new User("João Vitor", 12);
