CREATE OR REPLACE FUNCTION update_timestamp()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_updated_at
BEFORE UPDATE ON users
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();

CREATE TRIGGER set_updated_at
BEFORE UPDATE ON accounts_data
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();

CREATE TRIGGER set_updated_at
BEFORE UPDATE ON binary_data
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();

CREATE TRIGGER set_updated_at
BEFORE UPDATE ON cards_data
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();

CREATE TRIGGER set_updated_at
BEFORE UPDATE ON text_data
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();

CREATE TRIGGER set_updated_at
BEFORE UPDATE ON metadata
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();