
//Authenticate first
var m = new Mongo()
var authDB = m.getDB("test")
authDB.auth('admin','adminpassword');


//Check if dbName was passed as argument
//Otherwise use "ogreeDevelop"
try {
  dbName;
} catch(e) {
  dbName = "ogreeDevelop"
}
var db = m.getDB(dbName)

db.createUser({ user: dbName+"Admin", pwd: "adminpassword",
                roles: [{role: "readWrite", db: dbName}]
                })