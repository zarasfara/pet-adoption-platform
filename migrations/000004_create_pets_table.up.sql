CREATE TABLE IF NOT EXISTS pets (
    id SERIAL PRIMARY KEY,
    description TEXT,
    name VARCHAR(50),
    age INTEGER,
    is_available BOOLEAN DEFAULT TRUE,
    shelter_id INTEGER REFERENCES shelters(id),
    breed_id INTEGER REFERENCES breeds(id)
);
