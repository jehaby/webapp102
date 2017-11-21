CREATE TABLE ads(
  uuid UUID PRIMARY KEY,
  name TEXT NOT NULL,
  description TEXT NOT NULL,
  user_id UUID REFERENCES users(uuid),
  component_id UUID REFERENCES components(id)
);