CREATE TYPE CURRENCY AS ENUM ('RUB', 'USD', 'EUR');

CREATE TABLE localities (
  id INT PRIMARY KEY,
  name TEXT NOT NULL
);

CREATE TABLE ads (
  uuid         UUID PRIMARY KEY,
  name         TEXT                         NOT NULL,
  description  TEXT                         NOT NULL,
  user_uuid    UUID REFERENCES users (uuid) NOT NULL,
  component_id INT REFERENCES components (id) NOT NULL,
  category_id  INT REFERENCES categories (id) NOT NULL,
  price        INT                          NOT NULL,
  currency     CURRENCY                     NOT NULL,
  locality_id  INT REFERENCES localities (id) NOT NULL,
  created_at   TIMESTAMP WITH TIME ZONE     NOT NULL,
  updated_at   TIMESTAMP WITH TIME ZONE,
  deleted_at   TIMESTAMP WITH TIME ZONE
);
