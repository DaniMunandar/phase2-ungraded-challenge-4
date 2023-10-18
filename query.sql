-- Membuat tabel Heroes
CREATE TABLE Heroes (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    universe VARCHAR(255) NOT NULL,
    skill VARCHAR(255) NOT NULL,
    imageURL VARCHAR(255) NOT NULL
);

-- Membuat tabel Villain
CREATE TABLE Villain (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    universe VARCHAR(255) NOT NULL,
    imageURL VARCHAR(255) NOT NULL
);

-- Membuat tabel CriminalReports 
CREATE TABLE CriminalReports (
    id INT AUTO_INCREMENT PRIMARY KEY,
    hero_id INT NOT NULL,
    villain_id INT NOT NULL,
    description TEXT NOT NULL,
    incident DATETIME, 
    FOREIGN KEY (hero_id) REFERENCES Heroes(id),
    FOREIGN KEY (villain_id) REFERENCES Villain(id)
);

-- Memasukkan contoh data ke dalam tabel Heroes
INSERT INTO Heroes (name, universe, skill, imageURL) VALUES
    ('Iron Man', 'Marvel', 'Genius inventor and suit of armor', 'ironman.jpg'),
    ('Captain America', 'Marvel', 'Super Soldier', 'captainamerica.jpg'),
    ('Thor', 'Marvel', 'God of Thunder', 'thor.jpg');

-- Memasukkan contoh data ke dalam tabel Villain
INSERT INTO Villain (name, universe, imageURL) VALUES
    ('Loki', 'Marvel', 'loki.jpg'),
    ('Red Skull', 'Marvel', 'redskull.jpg'),
    ('Thanos', 'Marvel', 'thanos.jpg');

-- Memasukkan contoh data ke dalam tabel CriminalReports
INSERT INTO CriminalReports (hero_id, villain_id, description, incident) VALUES
    (1, 3, 'Iron Man vs. Thanos battle in New York', '2023-10-15'),
    (2, 1, 'Captain America foils Loki''s mischief in Asgard', '2023-10-16'),
    (3, 2, 'Thor battles Red Skull in the ancient ruins', '2023-10-14');
