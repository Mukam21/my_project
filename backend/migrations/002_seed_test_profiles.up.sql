INSERT INTO users (id, name, avatar, profile_type) VALUES
    ('11111111-1111-1111-1111-111111111111', 'Алексей Продавец', 'https://i.pravatar.cc/150?img=1', 'seller'),
    ('22222222-2222-2222-2222-222222222222', 'Мария Покупатель', 'https://i.pravatar.cc/150?img=2', 'buyer'),
    ('33333333-3333-3333-3333-333333333333', 'Иван Ветеран', 'https://i.pravatar.cc/150?img=3', 'veteran'),
    ('44444444-4444-4444-4444-444444444444', 'Елена Новичок', 'https://i.pravatar.cc/150?img=4', 'newbie'),
    ('55555555-5555-5555-5555-555555555555', 'Пётр Универсал', 'https://i.pravatar.cc/150?img=5', 'universal')
ON CONFLICT (id) DO NOTHING;

INSERT INTO categories (id, name) VALUES
    ('aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa', 'Авто'),
    ('bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb', 'Недвижимость'),
    ('cccccccc-cccc-cccc-cccc-cccccccccccc', 'Электроника'),
    ('dddddddd-dddd-dddd-dddd-dddddddddddd', 'Работа'),
    ('eeeeeeee-eeee-eeee-eeee-eeeeeeeeeeee', 'Услуги'),
    ('ffffffff-ffff-ffff-ffff-ffffffffffff', 'Одежда'),
    ('11111111-1111-1111-1111-aaaaaaaaaaaa', 'Животные'),
    ('22222222-2222-2222-2222-bbbbbbbbbbbb', 'Хобби')
ON CONFLICT (id) DO NOTHING;

-- Масштабируемые тестовые действия: пороги ачивок достижимы на демо.
WITH profile_actions (user_id, action_type, action_count, day_span, category_id) AS (
    VALUES
        -- Алексей (seller): много sales + views
        ('11111111-1111-1111-1111-111111111111'::uuid, 'view',     620, 200, 'cccccccc-cccc-cccc-cccc-cccccccccccc'::uuid),
        ('11111111-1111-1111-1111-111111111111'::uuid, 'message',   40, 200, 'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa'::uuid),
        ('11111111-1111-1111-1111-111111111111'::uuid, 'favorite',  30, 200, 'cccccccc-cccc-cccc-cccc-cccccccccccc'::uuid),
        ('11111111-1111-1111-1111-111111111111'::uuid, 'purchase',   3, 200, 'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa'::uuid),
        ('11111111-1111-1111-1111-111111111111'::uuid, 'sale',      12, 200, 'cccccccc-cccc-cccc-cccc-cccccccccccc'::uuid),

        -- Мария (buyer): views + favorites + purchases
        ('22222222-2222-2222-2222-222222222222'::uuid, 'view',    1100, 180, 'bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb'::uuid),
        ('22222222-2222-2222-2222-222222222222'::uuid, 'message',   55, 180, 'bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb'::uuid),
        ('22222222-2222-2222-2222-222222222222'::uuid, 'favorite', 160, 180, 'bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb'::uuid),
        ('22222222-2222-2222-2222-222222222222'::uuid, 'purchase',  14, 180, 'bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb'::uuid),
        ('22222222-2222-2222-2222-222222222222'::uuid, 'sale',       1, 180, 'ffffffff-ffff-ffff-ffff-ffffffffffff'::uuid),

        -- Иван (veteran): почти каждый день
        ('33333333-3333-3333-3333-333333333333'::uuid, 'view',    1300, 320, 'dddddddd-dddd-dddd-dddd-dddddddddddd'::uuid),
        ('33333333-3333-3333-3333-333333333333'::uuid, 'message',   70, 320, 'eeeeeeee-eeee-eeee-eeee-eeeeeeeeeeee'::uuid),
        ('33333333-3333-3333-3333-333333333333'::uuid, 'favorite',  90, 320, 'eeeeeeee-eeee-eeee-eeee-eeeeeeeeeeee'::uuid),
        ('33333333-3333-3333-3333-333333333333'::uuid, 'purchase',  11, 320, 'ffffffff-ffff-ffff-ffff-ffffffffffff'::uuid),
        ('33333333-3333-3333-3333-333333333333'::uuid, 'sale',       6, 320, 'eeeeeeee-eeee-eeee-eeee-eeeeeeeeeeee'::uuid),

        -- Елена (newbie): мало активности
        ('44444444-4444-4444-4444-444444444444'::uuid, 'view',      45,  20, 'cccccccc-cccc-cccc-cccc-cccccccccccc'::uuid),
        ('44444444-4444-4444-4444-444444444444'::uuid, 'message',    3,  20, 'cccccccc-cccc-cccc-cccc-cccccccccccc'::uuid),
        ('44444444-4444-4444-4444-444444444444'::uuid, 'favorite',   4,  20, 'cccccccc-cccc-cccc-cccc-cccccccccccc'::uuid),

        -- Пётр (universal): сбалансированно высокий
        ('55555555-5555-5555-5555-555555555555'::uuid, 'view',     800, 150, '11111111-1111-1111-1111-aaaaaaaaaaaa'::uuid),
        ('55555555-5555-5555-5555-555555555555'::uuid, 'message',   60, 150, '22222222-2222-2222-2222-bbbbbbbbbbbb'::uuid),
        ('55555555-5555-5555-5555-555555555555'::uuid, 'favorite',  80, 150, '22222222-2222-2222-2222-bbbbbbbbbbbb'::uuid),
        ('55555555-5555-5555-5555-555555555555'::uuid, 'purchase',  10, 150, '22222222-2222-2222-2222-bbbbbbbbbbbb'::uuid),
        ('55555555-5555-5555-5555-555555555555'::uuid, 'sale',       8, 150, '11111111-1111-1111-1111-aaaaaaaaaaaa'::uuid)
),
expanded AS (
    SELECT
        pa.user_id,
        pa.action_type,
        pa.category_id,
        pa.day_span,
        generate_series(1, pa.action_count) AS seq
    FROM profile_actions pa
)
INSERT INTO actions (id, user_id, type, category_id, created_at)
SELECT
    md5(user_id::text || ':' || action_type || ':' || seq::text)::uuid,
    user_id,
    action_type,
    category_id,
    TIMESTAMPTZ '2025-01-01 00:00:00+00'
        + ((seq - 1) % day_span) * INTERVAL '1 day'
        + ((seq * 37) % 86400) * INTERVAL '1 second'
FROM expanded
ON CONFLICT (id) DO NOTHING;