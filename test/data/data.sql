INSERT INTO users (uuid, name, email, password, created_at)
VALUES
  (
    'e12087ab-23b9-4d97-8b61-e7016e4e956b',
    'urf',
    'u@j.com',
    '$2a$10$R2iIpKeBPb12wcF3cZnzDuzlWKbM4fyFQo01S2d5eiNEXMO.8t7cS',
    now()
  );

INSERT INTO brands (id, name) VALUES
    (1, 'Shimano'),
    (2, 'SRAM'),
    (3, 'Gary Fisher'),
    (4, 'Merida'),
    (5, 'Stels'),
    (6, 'Scott'),
    (7, 'Ritchey')
;

INSERT INTO categories(id, path, name) VALUES
  (1, '{1}', 'Complete bikes'),
  (2, '{1, 2}', 'Hardtail'),
  (3, '{1, 2, 3}', 'Tour / Cross Country'),

  (10, '{10}', 'Components'),
  (11, '{10, 11}', 'Drive/circuit'),
  (50, '{10, 11, 50}', 'Chains'),
  (51, '{10, 11, 51}', 'Cassettes')
;

INSERT INTO products (brand_id, category_id, name) VALUES
(1, 50, 'CN-HG95'),
(1, 50, 'CN-HG54'),

(2, 50, 'PC 1051')
;

INSERT INTO localities (id, name) VALUES
(1, 'Moscow'),
(2, 'Saint-Petersburg')
;

INSERT INTO ads(
  uuid,
  name,
  description,
  user_uuid,
  condition,
  product_id,
  category_id,
  locality_id,
  price,
  currency,
  weight,
  brand_id,
  created_at
  ) VALUES
(
  '5df5b126-1fac-4fe1-a421-972ba56eb17b',
  'PC 1051222',
  'very very cool chain bro',
  'e12087ab-23b9-4d97-8b61-e7016e4e956b',
  'USED',
  2,
  50,
  2,
  50000,
  'RUB',
  200,
  2, 
  now()
)
;
