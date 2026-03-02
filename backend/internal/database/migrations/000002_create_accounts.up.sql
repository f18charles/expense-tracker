CREATE TABLE accounts (
    id          UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id     UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    name        VARCHAR(225) NOT NULL,
    type        VARCHAR(225) NOT NULL CHECK (type IN ('bank','mpesa','cash')),
    balance     NUMERIC(15,2) DEFAULT 0,
    currency    VARCHAR(10) DEFAULT 'KES',
    created_at  TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);