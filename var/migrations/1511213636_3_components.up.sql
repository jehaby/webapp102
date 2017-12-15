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
  id SERIAL PRIMARY KEY,
  manufacturer_id INT REFERENCES manufacturers(id),
  category_id SMALLINT REFERENCES categories(id),
  name VARCHAR NOT NULL
);

INSERT INTO categories(id, path, name) VALUES
  (1, '{1}', 'Complete bikes'),
  (2, '{1, 2}', 'Hardtail'),
  (3, '{1, 2, 3}', 'Tour / Cross Country')
;