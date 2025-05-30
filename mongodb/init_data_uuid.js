// init-data.js
try {
  db = db.getSiblingDB("uuidDB");
  
  if (!db.getCollection("uuid").exists()) {
    db.createCollection("uuid");
  }
  
  // Clear existing data
  db.uuid.deleteMany({});
  
  // Insert beacon information for all classrooms
  db.uuid.insertMany([
    {
      uuid: "ADD8CE0A-EF05-4B57-AD8C-7651198EAB2C",
      classroom: "Building 302",
      createdAt: new Date().toISOString()
    },
    {
      uuid: "BDD8CE0A-EF05-4B57-AD8C-7651198EAB2D",
      classroom: "Building 301",
      createdAt: new Date().toISOString()
    },
    {
      uuid: "CDD8CE0A-EF05-4B57-AD8C-7651198EAB2E",
      classroom: "Building 303",
      createdAt: new Date().toISOString()
    },
    {
      uuid: "DDD8CE0A-EF05-4B57-AD8C-7651198EAB2F",
      classroom: "Building 304",
      createdAt: new Date().toISOString()
    }
  ]);
  
  print("Data initialization completed successfully");
} catch (error) {
  print("Error during initialization:", error);
  throw error;
}