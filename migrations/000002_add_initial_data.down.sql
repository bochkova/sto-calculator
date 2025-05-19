-- Remove parameters
DELETE FROM parameters WHERE calculation_id IN (1, 2, 3, 4, 5, 6, 7, 8, 9);

-- Remove calculations
DELETE FROM calculations WHERE id IN (1, 2, 3, 4, 5, 6, 7, 8, 9);

-- Remove units
DELETE FROM units WHERE id IN (1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23);