CREATE TABLE users (
    id SERIAL PRIMARY KEY,

    login VARCHAR(64) UNIQUE NOT NULL,
    password VARCHAR(64) NOT NULL,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE accounts_data (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users (id),

    metadata TEXT,

    login VARCHAR(64),
    password VARCHAR(64),

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE text_data (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users (id),

    metadata TEXT,

    data TEXT,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE binary_data (
    id SERIAL PRIMARY KEY,

    user_id INTEGER REFERENCES users (id),

    metadata TEXT,

    data BYTEA,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE cards_data (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users (id),

    metadata TEXT,

    number VARCHAR(16),
    cardholder_name VARCHAR(32),
    expire DATE,
    cvv VARCHAR(3),

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


