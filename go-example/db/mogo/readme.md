use admin
db.createUser({user:"root",pwd:"123456",roles:["userAdminAnyDatabase","dbAdminAnyDatabase"]})  

db.createUser(
  {
    user: "xingcuntian",
    pwd: "xingcuntian",
    roles: [ { role: "userAdminAnyDatabase", db: "admin" } ]
  }
)


use admin
db.createUser(
  {
    user: "gary",
    pwd: "gary",
    roles: [ { role: "userAdminAnyDatabase", db: "admin" } ]
  }
)

use admin
db.auth('gary', 'gary')


mongo  -u "admin" -p "123456" --authenticationDatabase "admin"



db.createUser({user:"administrator",pwd:"123456",roles:["__system"]})  


mongodb://xingcuntian:123456@localhost:27017/xingcuntian