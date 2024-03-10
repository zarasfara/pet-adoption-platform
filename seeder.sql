-- users
INSERT INTO users (name, email, password, preferences) VALUES ('Alice Johnson', 'alice@example.com', 'password123', 'люблю собак');
INSERT INTO users (name, email, password, preferences) VALUES ('Bob Smith', 'bob@example.com', 'password456', 'люблю кошек');

-- shelters
INSERT INTO shelters (name, region) VALUES ('Happy Paws Shelter', 'Central');
INSERT INTO shelters (name, region) VALUES ('Loving Hearts Animal Shelter', 'East');

-- breeds
INSERT INTO breeds (name) VALUES ('Labrador Retriever');
INSERT INTO breeds (name) VALUES ('Siamese Cat');

-- pets
INSERT INTO pets (description, name, age, is_available, shelter_id, breed_id) VALUES ('Friendly and energetic dog', 'Max', 2, true, 1, 1);
INSERT INTO pets (description, name, age, is_available, shelter_id, breed_id) VALUES ('Playful and affectionate cat', 'Luna', 1, true, 2, 2);
