CREATE TABLE countries(
  id SMALLSERIAL PRIMARY KEY ,
  name VARCHAR
);

CREATE TABLE addresses(
  uuid UUID,
  --  locality_id INT REFERENCES localities.id
  country_id INTEGER REFERENCES countries(id)
);
