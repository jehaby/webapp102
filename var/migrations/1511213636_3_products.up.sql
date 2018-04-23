CREATE TABLE brands (
  id SERIAL PRIMARY KEY,
  name VARCHAR NOT NULL
);

-- represents categories using materialized path
CREATE TABLE categories (
  id SMALLSERIAL PRIMARY KEY,
  path SMALLINT[] NOT NULL,
  name VARCHAR NOT NULL
);

CREATE TABLE products (
  id SERIAL PRIMARY KEY,
  brand_id INT NOT NULL REFERENCES brands(id),
  category_id SMALLINT NOT NULL REFERENCES categories(id),
  name VARCHAR NOT NULL
);
