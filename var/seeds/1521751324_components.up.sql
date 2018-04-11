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
