// Create a Map
const fruits = new Map([
  ["apples", 500],
  ["bananas", 300],
  ["oranges", 200],
]);

let values2 = fruits.values();
let values = fruits.get("apples");
console.log(values * 2);
console.log(values2);
console.log(fruits.keys());
