CREATE TABLE transactions (
    id                  UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id             UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    account_id          UUID NOT NULL REFERENCES accounts(id) ON DELETE RESTRICT,
    category_id         UUID REFERENCES categories(id) ON DELETE SET NULL,
    amount              NUMERIC(15, 2) NOT NULL,
    type                VARCHAR(50) NOT NULL CHECK(type IN ('income', 'expense', 'transfer')),
    description         VARCHAR(500),
    payment_method      VARCHAR(50) CHECK (payment_method IN ('cash', 'mpesa', 'card', 'bank_transfer')),
    reference_id        VARCHAR(225),
    status              VARCHAR(50) DEFAULT 'completed' CHECK (status IN ('pending', 'completed', 'failed')),
    transaction_date    TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    created_at          TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE INDEX idx_transaction_user_id ON transactions(user_id);
CREATE INDEX idx_transaction_account_id ON transactions(account_id);
CREATE INDEX idx_transaction_transaction_date ON transactions(transaction_date);