CREATE TABLE manufacturers (
  id SERIAL PRIMARY KEY,
  name VARCHAR NOT NULL
);

-- represents categories using materialized path
CREATE TABLE categories (
  id SMALLSERIAL PRIMARY KEY,
  path SMALLINT[] NOT NULL,
  name VARCHAR NOT NULL
);

CREATE TABLE components (
  id INT PRIMARY KEY,
  manufacturer_id INT REFERENCES manufacturers(id),
  name VARCHAR NOT NULL
);
