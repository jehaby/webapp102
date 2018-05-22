CREATE TYPE PROPERTY_TYPE AS ENUM ('RANGE', 'VALUES', 'BOOL');

-- Stores properties for categories -- TODO: // indexes
CREATE TABLE properties(
    id SERIAL PRIMARY KEY,
    category_id INT REFERENCES categories(id) NOT NULL,
    name TEXT NOT NULL,
    "type" PROPERTY_TYPE NOT NULL,
    required BOOLEAN NOT NULL,
    "values" JSONB
);
