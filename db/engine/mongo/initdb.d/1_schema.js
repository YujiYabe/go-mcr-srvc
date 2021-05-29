var user = {
    user: "user",
    pwd: "user",
    roles: [
        {
            role: "dbOwner",
            db: "app"
        }
    ]
};

db.createUser(user);
db.createCollection('staffs')

