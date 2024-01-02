CREATE TABLE IF NOT EXISTS pets (
    id SERIAL PRIMARY KEY,
    description TEXT,
    name VARCHAR(50) NOT NULL,
    age INTEGER NOT NULL,
    is_available BOOLEAN DEFAULT TRUE,
    shelter_id INTEGER REFERENCES shelters(id),
    breed_id INTEGER REFERENCES breeds(id)
);
