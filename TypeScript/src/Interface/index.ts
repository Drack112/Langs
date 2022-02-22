interface User {
  name: string;
  last: string;
}

function getName(oof: User) {
  return console.log(oof.name, oof.last);
}

let Draco: User = {
  name: "Drack",
  last: "oof",
};

getName(Draco);
