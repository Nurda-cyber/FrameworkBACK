CREATE TABLE toys (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT,
    price NUMERIC,
    category_id INTEGER REFERENCES categories(id) ON DELETE SET NULL
);