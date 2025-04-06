CREATE TABLE users (
    id SERIAL PRIMARY KEY,

    login VARCHAR(64) UNIQUE,
    password VARCHAR(64),

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE accounts_data (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users (id),

    login VARCHAR(64),
    password VARCHAR(64),

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE text_data (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users (id),

    data TEXT,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE binary_data (
    id SERIAL PRIMARY KEY,

    user_id INTEGER REFERENCES users (id),
    data BYTEA,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE cards_data (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users (id),

    number VARCHAR(16),
    cardholder_name VARCHAR(32),
    expire DATE,
    cvv VARCHAR(3),

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE metadata (
    id SERIAL PRIMARY KEY,

    table_name VARCHAR(64),
    row_id INTEGER,

    data TEXT,

    CHECK (table_name IN ('accounts_data, text_data', 'binary_data', 'cards_data')),

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

