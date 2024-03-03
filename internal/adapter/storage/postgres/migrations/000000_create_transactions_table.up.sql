CREATE
EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE transactions
(
    id            UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name          VARCHAR,
    amount        VARCHAR,
    transactionAt TIMESTAMP
);