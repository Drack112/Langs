//O async/await é uma nova forma de tratar
// Promises dentro do nosso código, evitando a criação de cascatas de .then

async function fetchUserAndGroups() {
  const user = await api.get("/users/diego3g");
  const groups = await api.get(`/groups/${user.id}`);

  groups.map((group) => {
    const groupInfo = await api.get(`/group/${group.id}`);
    console.log(groupInfo);
  });
}
