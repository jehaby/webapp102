-- Stores properties for categories -- TODO: // indexes
CREATE TABLE properties(
    category_id INT REFERENCES categories(id) NOT NULL,
    name TEXT NOT NULL,
    "values" JSONB
);
