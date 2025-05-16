// init-data.js
db = db.getSiblingDB("uuidDB") // DB 선택 (없으면 생성됨)

db.createCollection("uuid")

const id = ObjectId()

db.users.insertOne({
  _id: id,
  uuid: "ADD8CE0A-EF05-4B57-AD8C-7651198EAB2C",
  classroom: "Building 302",
  createdAt: new Date()
})


