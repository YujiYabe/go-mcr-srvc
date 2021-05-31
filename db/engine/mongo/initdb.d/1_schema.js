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
db.createCollection('bans')
// db.createCollection('patties')
// db.createCollection('vegetables')
// db.createCollection('ingredients')




// var user = {
//     user: "user",
//     pwd: "user",
//     roles: [
//         {
//             role: "userAdminAnyDatabase",
//             db: "app"
//         },
//         "readWriteAnyDatabase"
//     ]
// };

// db.createUser(user);
// db.createCollection('bans')
// db.createCollection('patties')
// db.createCollection('vegetables')
// db.createCollection('ingredients')

