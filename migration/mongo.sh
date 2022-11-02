#!/bin/bash
mongosh <<EOF
use admin 
db.createUser(
  {
    user: "root",
    pwd: "password",
    roles: [ { role: "userAdminAnyDatabase", db: "admin" }, "readWriteAnyDatabase" ]
  }
)

use tour-of-heroes
db.createCollection("hero")
db.hero.insertMany([
  { "hero_id": 12, "name": "Dr. Nice" },
  { "hero_id": 13, "name": "Bombasto" },
  { "hero_id": 14, "name": "Celeritas" },
  { "hero_id": 15, "name": "Magneta" },
  { "hero_id": 16, "name": "RubberMan" },
  { "hero_id": 17, "name": "Dynama" },
  { "hero_id": 18, "name": "Dr. IQ" },
  { "hero_id": 19, "name": "Magma" },
  { "hero_id": 20, "name": "Tornado" },
])
EOF