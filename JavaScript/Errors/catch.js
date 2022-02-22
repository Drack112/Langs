let num = 1;
try {
  num.toPrecision(500); // A number cannot have 500 significant digits
} catch (err) {
  console.log(err.name);
}
