"use strict";
function Greet(name: string, greeting: string = "Hello") {
  return console.log(greeting + " " + name + "!");
}
Greet("Steve"); //OK, returns "Hello Steve!"
Greet("Steve", "Hi"); // OK, returns "Hi Steve!".
Greet("Bill"); //OK, returns "Hello Bill!"
