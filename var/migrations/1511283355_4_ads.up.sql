CREATE TYPE CURRENCY AS ENUM ('RUB', 'USD', 'EUR');
CREATE TYPE CONDITION AS ENUM('NEW', 'USED_LIKE_NEW', 'USED', 'MALFUNCTIONED');

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


CREATE TABLE localities (
  id INT PRIMARY KEY,
  name TEXT NOT NULL
);

CREATE TABLE ads (
  id           SERIAL UNIQUE,
  uuid          UUID PRIMARY KEY,
  name         TEXT                         NOT NULL,
  description  TEXT                         NOT NULL,
  user_uuid    UUID REFERENCES users (uuid) NOT NULL,
  condition    CONDITION                    NOT NULL,
  category_id  INT REFERENCES categories (id)  NOT NULL,
  brand_id INT REFERENCES brands (id),
  weight       INT CONSTRAINT positive_weight CHECK(weight > 0),
  price        INT NOT NULL CONSTRAINT positive_price CHECK(price >= 0),
  phone_uuid   UUID REFERENCES phones (uuid) ON DELETE SET NULL,
  properties   JSONB,
  currency     CURRENCY                        NOT NULL,
  locality_id  INT REFERENCES localities (id)  NOT NULL,
  created_at   TIMESTAMP WITH TIME ZONE        NOT NULL DEFAULT NOW(),
  updated_at   TIMESTAMP WITH TIME ZONE,
  deleted_at   TIMESTAMP WITH TIME ZONE
);
