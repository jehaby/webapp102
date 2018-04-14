INSERT INTO users (uuid, name, email, password, created_at)
VALUES
  (
    'e12087ab-23b9-4d97-8b61-e7016e4e956b',
    'urf',
    'u@j.com',
    '$2a$10$R2iIpKeBPb12wcF3cZnzDuzlWKbM4fyFQo01S2d5eiNEXMO.8t7cS',
    now()
  );

INSERT INTO manufacturers (id, name) VALUES 
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
  (51, '{10, 11, 50}', 'Cassettes')
;

INSERT INTO components (id, manufacturer_id, category_id, name) VALUES 
(1, 1, 50, 'CN-HG95'),
(2, 1, 50, 'CN-HG54'),

(3, 2, 50, 'PC 1051')
;

INSERT INTO ads(uuid, name, description, user_uuid, component_id, category_id, created_at) VALUES
('5df5b126-1fac-4fe1-a421-972ba56eb17b', 'Cool chain', 'very very cool chain bro', 'e12087ab-23b9-4d97-8b61-e7016e4e956b', 2, 50, now())
;
