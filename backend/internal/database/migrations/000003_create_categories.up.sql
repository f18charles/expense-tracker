CREATE TABLE categories (
    id          UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id     UUID REFERENCES users(id) ON DELETE CASCADE,
    name        VARCHAR(225) NOT NULL,
    type        VARCHAR(50) NOT NULL CHECK(type IN ('income', 'expense')),
    color       VARCHAR(20),
    icon        VARCHAR(50),
    is_default  BOOLEAN DEFAULT false,
    created_at  TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
);

INSERT INTO categories (name, type, is_default) VALUES 
    ('Salary', 'income', true),
    ('Freelance', 'income', true),
    ('Business', 'income', true),
    ('Food & Drinks', 'expense', true),
    ('Transport', 'expense', true),
    ('Rent', 'expense', true),
    ('Utilities', 'expense', true),
    ('Entertainment', 'expense', true),
    ('Healthcare', 'expense', true),
    ('Shopping', 'expense', true),
    ('Savings', 'expense', true),
    ('Other', 'expense', true);