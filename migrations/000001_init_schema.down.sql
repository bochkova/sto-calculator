-- Drop triggers
DROP TRIGGER IF EXISTS update_parameters_updated_at ON parameters;
DROP TRIGGER IF EXISTS update_calculations_updated_at ON calculations;
DROP TRIGGER IF EXISTS update_units_updated_at ON units;

-- Drop function
DROP FUNCTION IF EXISTS update_updated_at_column();

-- Drop tables
DROP TABLE IF EXISTS parameters;
DROP TABLE IF EXISTS calculations;
DROP TABLE IF EXISTS units; 