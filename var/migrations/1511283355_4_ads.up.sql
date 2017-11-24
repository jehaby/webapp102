CREATE TABLE ads(
  uuid UUID PRIMARY KEY,
  name TEXT NOT NULL,
  description TEXT NOT NULL,
  user_uuid UUID REFERENCES users(uuid) NOT NULL,
  component_id INT REFERENCES components(id),
  created_at TIMESTAMP WITH TIME ZONE NOT NULL,
  updated_at TIMESTAMP WITH TIME ZONE
);