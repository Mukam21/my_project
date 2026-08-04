CREATE TABLE IF NOT EXISTS profile_details (
    user_id UUID PRIMARY KEY REFERENCES users(id) ON DELETE CASCADE,
    bio TEXT,
    city TEXT,
    phone TEXT,
    telegram TEXT,
    website TEXT
);

INSERT INTO profile_details (user_id, bio, city, phone, telegram, website) VALUES
    ('11111111-1111-1111-1111-111111111111', 'Продаю всё, что можно продать', 'Москва', '+7-999-111-22-33', '@alexey_seller', 'alexey.ru'),
    ('22222222-2222-2222-2222-222222222222', 'Люблю покупать интересные вещи', 'Санкт-Петербург', '+7-999-222-33-44', '@maria_buyer', NULL),
    ('33333333-3333-3333-3333-333333333333', 'Здесь с самого начала', 'Новосибирск', '+7-999-333-44-55', '@ivan_veteran', 'ivan.ru'),
    ('44444444-4444-4444-4444-444444444444', 'Только начинаю свой путь', 'Екатеринбург', '+7-999-444-55-66', '@elena_newbie', NULL),
    ('55555555-5555-5555-5555-555555555555', 'Готов к любым сценариям', 'Казань', '+7-999-555-66-77', '@petr_universal', 'petr.ru');