-- Habilitar extensão de UUID
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Inserir categorias
INSERT INTO categories (name) VALUES
('Streaming de Vídeo'),
('Streaming de Esportes'),
('Streaming de Entretenimento'),
('Serviços Esportivos');

-- Inserir produtos
INSERT INTO products (name, category_id, monthly_price) VALUES
('Globoplay Padrão', 1, 19.90),
('Globoplay Premium', 1, 29.90),
('Premiere', 2, 79.90),
('Telecine', 3, 34.90),
('Cartola PRO', 4, 29.90);

-- Inserir 5000 usuários
INSERT INTO users (id, name, email, created_at)
SELECT 
    uuid_generate_v4(),
    'Usuário ' || seq,
    'usuario' || seq || '@exemplo.com',
    CURRENT_TIMESTAMP - (random() * 365 || ' days')::interval
FROM generate_series(1,10) seq;

-- Inserir assinaturas (múltiplas e realmente aleatórias por usuário)
WITH user_subscriptions AS (
    SELECT 
        u.id as user_id,
        CASE 
            WHEN random() < 0.40 THEN 1
            WHEN random() < 0.75 THEN 2
            WHEN random() < 0.95 THEN 3
            ELSE 4
        END as subscription_count
    FROM users u
),
subscription_details AS (
    SELECT 
        us.user_id,
        p.id AS product_id,
        CURRENT_TIMESTAMP - (random() * 365 || ' days')::interval AS start_date,
        random() > 0.2 AS active,
        random() AS rand_order
    FROM user_subscriptions us
    JOIN products p ON random() < 0.6  -- controla chance de incluir cada produto (~60%)
),
ranked_subscriptions AS (
    SELECT 
        user_id,
        product_id,
        start_date,
        active,
        ROW_NUMBER() OVER (PARTITION BY user_id ORDER BY rand_order) AS rn
    FROM subscription_details
)
INSERT INTO subscriptions (user_id, product_id, start_date, active)
SELECT 
    user_id,
    product_id,
    start_date,
    active
FROM ranked_subscriptions
WHERE rn <= (
    SELECT subscription_count FROM user_subscriptions u2 WHERE u2.user_id = ranked_subscriptions.user_id
);

-- Garantir pelo menos 1 assinatura por usuário (caso algum tenha ficado sem)
INSERT INTO subscriptions (user_id, product_id, start_date, active)
SELECT
    u.id,
    p.id,
    CURRENT_TIMESTAMP - (random() * 365 || ' days')::interval,
    random() > 0.2
FROM users u
JOIN LATERAL (
    SELECT id FROM products ORDER BY random() LIMIT 1
) p ON TRUE
WHERE NOT EXISTS (
    SELECT 1 FROM subscriptions s WHERE s.user_id = u.id
);
