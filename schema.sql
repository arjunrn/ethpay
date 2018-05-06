CREATE TYPE payment_status AS ENUM ('submitted', 'processed', 'confirmed');
CREATE TABLE payments(
    id SERIAL PRIMARY KEY,
    transaction_hash VARCHAR(200) NOT NULL,
    last_updated TIMESTAMP NOT NULL,
    status payment_status NOT NULL,
    block_number INT DEFAULT -1
);