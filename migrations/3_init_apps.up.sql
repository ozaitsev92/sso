INSERT INTO apps (id, name, secret)
VALUES (1, 'test', 'dGVzdC1zZWNyZXQ=')
ON CONFLICT DO NOTHING;