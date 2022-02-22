let i = 0;

let oof = ["oof", "ooof"];

for (i in oof) {
  console.log(i);
  i++;
}

for (i of oof) {
  console.log(i);
  i++;
}
