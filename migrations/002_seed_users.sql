-- +migrate Up
INSERT INTO users (username, password, created_by) VALUES (
    'admin',
    '$2a$10$qjIXhljgW1OaDJYLZEKoxOmQFWClcKG9J3lFKBtjvn75eQ2.fgcze',
    'system'
);

-- +migrate Down
DELETE FROM users WHERE username = 'admin';