INSERT INTO users (id, name, avatar, profile_type) VALUES
    ('11111111-1111-1111-1111-111111111111', 'Алексей Продавец', 'https://i.pravatar.cc/150?img=1', 'seller'),
    ('22222222-2222-2222-2222-222222222222', 'Мария Покупатель', 'https://i.pravatar.cc/150?img=2', 'buyer'),
    ('33333333-3333-3333-3333-333333333333', 'Иван Ветеран', 'https://i.pravatar.cc/150?img=3', 'veteran'),
    ('44444444-4444-4444-4444-444444444444', 'Елена Новичок', 'https://i.pravatar.cc/150?img=4', 'newbie'),
    ('55555555-5555-5555-5555-555555555555', 'Пётр Универсал', 'https://i.pravatar.cc/150?img=5', 'universal');

INSERT INTO categories (id, name) VALUES
    ('aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa', 'Авто'),
    ('bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb', 'Недвижимость'),
    ('cccccccc-cccc-cccc-cccc-cccccccccccc', 'Электроника'),
    ('dddddddd-dddd-dddd-dddd-dddddddddddd', 'Работа'),
    ('eeeeeeee-eeee-eeee-eeee-eeeeeeeeeeee', 'Услуги'),
    ('ffffffff-ffff-ffff-ffff-ffffffffffff', 'Одежда'),
    ('11111111-1111-1111-1111-aaaaaaaaaaaa', 'Животные'),
    ('22222222-2222-2222-2222-bbbbbbbbbbbb', 'Хобби');

INSERT INTO actions (id, user_id, type, category_id, created_at) VALUES
    -- Алексей (seller)
    (gen_random_uuid(), '11111111-1111-1111-1111-111111111111', 'view', 'cccccccc-cccc-cccc-cccc-cccccccccccc', '2025-01-15 10:00:00'),
    (gen_random_uuid(), '11111111-1111-1111-1111-111111111111', 'view', 'cccccccc-cccc-cccc-cccc-cccccccccccc', '2025-01-16 12:00:00'),
    (gen_random_uuid(), '11111111-1111-1111-1111-111111111111', 'message', 'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa', '2025-02-10 14:00:00'),
    (gen_random_uuid(), '11111111-1111-1111-1111-111111111111', 'purchase', 'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa', '2025-02-11 16:00:00'),
    (gen_random_uuid(), '11111111-1111-1111-1111-111111111111', 'sale', 'cccccccc-cccc-cccc-cccc-cccccccccccc', '2025-03-01 10:00:00'),

    -- Мария (buyer)
    (gen_random_uuid(), '22222222-2222-2222-2222-222222222222', 'view', 'bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb', '2025-01-20 09:00:00'),
    (gen_random_uuid(), '22222222-2222-2222-2222-222222222222', 'view', 'bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb', '2025-01-21 11:00:00'),
    (gen_random_uuid(), '22222222-2222-2222-2222-222222222222', 'favorite', 'bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb', '2025-01-22 13:00:00'),
    (gen_random_uuid(), '22222222-2222-2222-2222-222222222222', 'message', 'bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb', '2025-02-05 15:00:00'),
    (gen_random_uuid(), '22222222-2222-2222-2222-222222222222', 'purchase', 'bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb', '2025-03-10 10:00:00'),

    -- Иван (veteran)
    (gen_random_uuid(), '33333333-3333-3333-3333-333333333333', 'view', 'dddddddd-dddd-dddd-dddd-dddddddddddd', '2025-01-01 08:00:00'),
    (gen_random_uuid(), '33333333-3333-3333-3333-333333333333', 'view', 'eeeeeeee-eeee-eeee-eeee-eeeeeeeeeeee', '2025-01-02 09:00:00'),
    (gen_random_uuid(), '33333333-3333-3333-3333-333333333333', 'view', 'ffffffff-ffff-ffff-ffff-ffffffffffff', '2025-01-03 10:00:00'),
    (gen_random_uuid(), '33333333-3333-3333-3333-333333333333', 'favorite', 'eeeeeeee-eeee-eeee-eeee-eeeeeeeeeeee', '2025-02-01 11:00:00'),
    (gen_random_uuid(), '33333333-3333-3333-3333-333333333333', 'message', 'dddddddd-dddd-dddd-dddd-dddddddddddd', '2025-02-15 12:00:00'),
    (gen_random_uuid(), '33333333-3333-3333-3333-333333333333', 'purchase', 'ffffffff-ffff-ffff-ffff-ffffffffffff', '2025-03-05 14:00:00'),
    (gen_random_uuid(), '33333333-3333-3333-3333-333333333333', 'sale', 'eeeeeeee-eeee-eeee-eeee-eeeeeeeeeeee', '2025-04-01 16:00:00'),

    -- Елена (newbie)
    (gen_random_uuid(), '44444444-4444-4444-4444-444444444444', 'view', 'cccccccc-cccc-cccc-cccc-cccccccccccc', '2025-06-01 10:00:00'),
    (gen_random_uuid(), '44444444-4444-4444-4444-444444444444', 'view', 'cccccccc-cccc-cccc-cccc-cccccccccccc', '2025-06-02 12:00:00'),

    -- Пётр (universal)
    (gen_random_uuid(), '55555555-5555-5555-5555-555555555555', 'view', '11111111-1111-1111-1111-aaaaaaaaaaaa', '2025-01-10 08:00:00'),
    (gen_random_uuid(), '55555555-5555-5555-5555-555555555555', 'view', '22222222-2222-2222-2222-bbbbbbbbbbbb', '2025-01-15 09:00:00'),
    (gen_random_uuid(), '55555555-5555-5555-5555-555555555555', 'favorite', '22222222-2222-2222-2222-bbbbbbbbbbbb', '2025-02-01 10:00:00'),
    (gen_random_uuid(), '55555555-5555-5555-5555-555555555555', 'message', '11111111-1111-1111-1111-aaaaaaaaaaaa', '2025-02-10 11:00:00'),
    (gen_random_uuid(), '55555555-5555-5555-5555-555555555555', 'purchase', '22222222-2222-2222-2222-bbbbbbbbbbbb', '2025-03-01 12:00:00'),
    (gen_random_uuid(), '55555555-5555-5555-5555-555555555555', 'sale', '11111111-1111-1111-1111-aaaaaaaaaaaa', '2025-04-01 14:00:00');