db.createUser({
  user: "root",
  pwd: "root",
  roles: [
    {
      role: "readWrite",
      db: "game-shop",
    },
  ],
});

db.createCollection("users");
