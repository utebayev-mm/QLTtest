-- DROP TABLE IF EXISTS transactions;
-- DROP TABLE IF EXISTS categories;

CREATE TABLE IF NOT EXISTS transactions (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,		
		name TEXT,
		price INTEGER,
		date TEXT,
		type BOOLEAN,
		comments TEXT,
		category_id INTEGER
	  );

CREATE TABLE IF NOT EXISTS categories (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name TEXT UNIQUE
	  );

-- INSERT INTO transactions (name, price) values("penis2",123)