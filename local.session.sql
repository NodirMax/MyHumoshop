-- ALTER TABLE users ALTER COLUMN password TYPE varchar(100); -- Устанавливаем размер в 50. Проверим, укладываются ли в него наши данные!
-- INSERT INTO users(name, login, password) VALUES ('Nodir', 'admin', 'admin')
-- ALTER TABLE orders ADD COLUMN product_name varchar(100)
SELECT * FROM product
-- ALTER TABLE orders DROP COLUMN product_id
-- ALTER TABLE orders DROP COLUMN products
-- ALTER TABLE orders ADD COLUMN products JSON[];
-- SELECT * FROM users WHERE login
