drop table users;


CREATE TABLE users (
  uuid UUID PRIMARY KEY ,
  name VARCHAR NOT NULL UNIQUE ,
  email VARCHAR NOT NULL UNIQUE,
  password VARCHAR NOT NULL,
  last_login TIMESTAMP WITH TIME ZONE,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL,
  updated_at TIMESTAMP WITH TIME ZONE
);


insert into
  users
  (uuid, name, email, password, created_at)
VALUES
  (
    'e12087ab-23b9-4d97-8b61-e7016e4e956b',
    'urf',
    'u@j.com',
    '$2a$10$R2iIpKeBPb12wcF3cZnzDuzlWKbM4fyFQo01S2d5eiNEXMO.8t7cS',
    now()
  )
;