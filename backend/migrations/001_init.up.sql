CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    avatar TEXT,
    registered_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    profile_type TEXT CHECK (profile_type IN ('seller', 'buyer', 'veteran', 'newbie', 'universal'))
);

CREATE TABLE IF NOT EXISTS categories (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS actions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    type TEXT NOT NULL,
    category_id UUID REFERENCES categories(id),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS recaps (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    year INT NOT NULL,
    total_views INT DEFAULT 0,
    total_messages INT DEFAULT 0,
    total_favorites INT DEFAULT 0,
    total_purchases INT DEFAULT 0,
    total_sales INT DEFAULT 0,
    activity_days INT DEFAULT 0,
    top_categories JSONB DEFAULT '[]',
    achievements JSONB DEFAULT '[]',
    generated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    UNIQUE(user_id, year)
);

CREATE INDEX IF NOT EXISTS idx_actions_user_created ON actions(user_id, created_at);
CREATE INDEX IF NOT EXISTS idx_actions_user_category ON actions(user_id, category_id);