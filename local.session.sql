-- ALTER TABLE users ALTER COLUMN password TYPE varchar(100); -- Устанавливаем размер в 50. Проверим, укладываются ли в него наши данные!
-- INSERT INTO users(name, login, password) VALUES ('Nodir', 'admin', 'admin')
-- ALTER TABLE orders ADD COLUMN product_name varchar(100)
SELECT * FROM orderproducts
-- ALTER TABLE orders DROP COLUMN product_id
-- ALTER TABLE orders DROP COLUMN products
-- ALTER TABLE orderproducts ADD COLUMN product_count INTEGER;
-- SELECT * FROM users WHERE login
