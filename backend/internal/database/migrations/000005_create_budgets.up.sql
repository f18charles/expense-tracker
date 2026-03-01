CREATE TABLE budgets (
    id          UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id     UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    category_id UUID NOT NULL REFERENCES categories(id) ON DELETE CASCADE,
    amount      NUMERIC(15,2) NOT NULL,
    spent       NUMERIC(15,2) DEFAULT 0,
    period      VARCHAR(20) DEFAULT 'monthly' CHECK (period IN ('monthly','weekly')),
    start_date  DATE NOT NULL,
    end_date    DATE NOT NULL,
    created_at  TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);