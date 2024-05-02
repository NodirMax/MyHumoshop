-- ALTER TABLE users ALTER COLUMN password TYPE varchar(100); -- Устанавливаем размер в 50. Проверим, укладываются ли в него наши данные!
SELECT * FROM users WHERE login='HomerSimpson'