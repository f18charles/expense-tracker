CREATE EXTENSION IF NOT EXISTS "uuid-ossp"

CREATE TABLE users (
    id              UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    email           VARCHAR(225) UNIQUE NOT NULL,
    password_hash   VARCHAR(225) NOT NULL,
    full_name       VARCHAR(225) NOT NULL,
    currency        VARCHAR(10) DEFAULT 'KES',
    created_at      TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at      TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
);