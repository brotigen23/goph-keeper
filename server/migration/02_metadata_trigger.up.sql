CREATE OR REPLACE FUNCTION create_metadata()
RETURNS TRIGGER AS $$
BEGIN
    INSERT INTO metadata
    DEFAULT VALUES;

    NEW.metadata_id := currval('metadata_id_seq');

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER account_metadata
BEFORE INSERT ON accounts_data
FOR EACH ROW
EXECUTE FUNCTION create_metadata();

CREATE TRIGGER text_metadata
BEFORE INSERT ON text_data
FOR EACH ROW
EXECUTE FUNCTION create_metadata();

CREATE TRIGGER binary_metadata
BEFORE INSERT ON binary_data
FOR EACH ROW
EXECUTE FUNCTION create_metadata();

CREATE TRIGGER cards_data
BEFORE INSERT ON cards_data
FOR EACH ROW
EXECUTE FUNCTION create_metadata();
