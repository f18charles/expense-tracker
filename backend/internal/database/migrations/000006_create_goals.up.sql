CREATE TABLE goals (
    id              UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id         UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    name            VARCHAR(225) NOT NULL,
    target_amount   NUMERIC(15,2) NOT NULL,
    current_amount  NUMERIC(15,2) DEFAULT 0,
    deadline        DATE,
    created_at      TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);