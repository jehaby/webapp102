CREATE TYPE PROPERTY_TYPE AS ENUM ('RANGE', 'VALUES');

-- Stores properties for categories -- TODO: // indexes
CREATE TABLE properties(
    id SERIAL PRIMARY KEY,
    category_id INT REFERENCES categories(id) NOT NULL,
    name TEXT NOT NULL,
    "type" PROPERTY_TYPE NOT NULL,
    "values" JSONB
);
