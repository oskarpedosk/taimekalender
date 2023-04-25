CREATE TABLE IF NOT EXISTS "users" (
    "id" INTEGER PRIMARY KEY AUTOINCREMENT,
    "email" TEXT NOT NULL,
    "password" TEXT NOT NULL,
    "firstName" TEXT NOT NULL,
    "lastName" TEXT NOT NULL,
    "dateOfBirth" TIMESTAMP NOT NULL,
    "img" TEXT NOT NULL DEFAULT "default_profile.png",
    "nickname" TEXT NOT NULL DEFAULT "",
    "about" TEXT NOT NULL DEFAULT "",
    "profilePublic" INTEGER NOT NULL DEFAULT 1,
    "createdAt" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Users

INSERT INTO users (email, password, firstName, lastName, dateOfBirth, nickname, about, img)
SELECT 'john@test.ee', '$2y$10$Vozj.mVrkY3c0i4Zf3Z6VOaQVJ0qiZq20AOlBsuKiCRTNAUdNIcPO', 'John', 'Smith', '1990-01-01', 'johnsmith', 'I love programming', 'john_smith.png'
WHERE NOT EXISTS
    (SELECT 1 FROM users WHERE email = 'john@test.ee');