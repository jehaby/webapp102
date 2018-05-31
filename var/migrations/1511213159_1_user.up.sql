CREATE TYPE USER_ROLE AS ENUM ('user', 'moderator', 'admin');

CREATE TABLE phones (
  uuid UUID PRIMARY KEY,
  country_code SMALLINT CONSTRAINT positive_country_code CHECK(country_code > 0),
  "number" VARCHAR(15) NOT NULL,
  UNIQUE (country_code, "number")
  -- TODO: index!
);

CREATE TABLE users (
  uuid       UUID PRIMARY KEY,
  name       VARCHAR                  NOT NULL UNIQUE,
  email      VARCHAR                  NOT NULL UNIQUE,
  password   VARCHAR                  NOT NULL,
  role       USER_ROLE                NOT NULL DEFAULT 'user',
  default_phone UUID REFERENCES phones (uuid),
  confirmed bool NOT NULL DEFAULT false,  
  confirmation_token VARCHAR,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL,
  updated_at TIMESTAMP WITH TIME ZONE,
  last_logout TIMESTAMP WITH TIME ZONE,
  banned_at TIMESTAMP WITH TIME ZONE,
  banned_info JSONB
);
