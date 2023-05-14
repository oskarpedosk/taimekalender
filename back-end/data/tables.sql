CREATE TABLE IF NOT EXISTS "users" (
    "id" INTEGER PRIMARY KEY AUTOINCREMENT,
    "username" TEXT NOT NULL,
    "password" TEXT NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "rooms" (
    "id" INTEGER PRIMARY KEY AUTOINCREMENT,
    "name" TEXT NOT NULL,
    "index" INTEGER NOT NULL DEFAULT 0,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "plants" (
    "id" INTEGER PRIMARY KEY AUTOINCREMENT,
    "room_id" INTEGER NOT NULL,
    "name" TEXT NOT NULL,
    "description" TEXT NOT NULL DEFAULT "",
    "index" INTEGER NOT NULL DEFAULT 0,
    "watering_interval" INTEGER NOT NULL DEFAULT 7,
    "watered" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "fertilizing_interval" INTEGER NOT NULL DEFAULT 60,
    "transplanting" TEXT NOT NULL DEFAULT "",
    "fertilized" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (room_id) REFERENCES rooms (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS "sessions" (
    "id" INTEGER PRIMARY KEY AUTOINCREMENT,
    "userId" INTEGER NOT NULL,
    "uuid" TEXT NOT NULL,
    "createdAt" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (userId) REFERENCES users (id) ON DELETE CASCADE
);

-- Users

INSERT INTO users (email, password)
SELECT 'john@test.ee', '$2y$10$Vozj.mVrkY3c0i4Zf3Z6VOaQVJ0qiZq20AOlBsuKiCRTNAUdNIcPO'
WHERE NOT EXISTS
    (SELECT 1 FROM users WHERE email = 'john@test.ee');

-- Rooms

INSERT INTO rooms (name, "index")
SELECT 'elutuba', 0
WHERE NOT EXISTS
    (SELECT 1 FROM rooms WHERE name = 'elutuba');

INSERT INTO rooms (name, "index")
SELECT 'magamistuba', 1
WHERE NOT EXISTS
    (SELECT 1 FROM rooms WHERE name = 'magamistuba');

-- Plants

INSERT INTO plants (room_id, name, description, "index")
SELECT 1, 'kaktus', "tore kaktus", 0
WHERE NOT EXISTS
    (SELECT 1 FROM plants WHERE name = 'kaktus');

INSERT INTO plants (room_id, name, description, "index")
SELECT 1, 'palm', "tore palm", 1
WHERE NOT EXISTS
    (SELECT 1 FROM plants WHERE name = 'palm');

INSERT INTO plants (room_id, name, description, "index")
SELECT 2, 'muru', "väike muru", 0
WHERE NOT EXISTS
    (SELECT 1 FROM plants WHERE name = 'muru');