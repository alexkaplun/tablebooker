package sqlite

var sqlInitDatabase = `
	DROP TABLE IF EXISTS tables;
	DROP TABLE IF EXISTS table_book;

	CREATE TABLE tables (id VARCHAR(50) PRIMARY KEY, table_number INT, capacity INT);
    CREATE TABLE table_book (
		id VARCHAR(50) PRIMARY KEY, 
		table_id VARCHAR(50), 
		book_date DATETIME,
		guest_name VARCHAR(255),
		guest_contact TEXT,
		code VARCHAR(50)
	);`

var sqlInsertTable = `
	INSERT INTO tables (id, table_number, capacity)
	VALUES (?, ?, ?)
`

var sqlSelectAllTables = `
	SELECT id, table_number, capacity FROM tables
`

var sqlTableExists = `
	SELECT id FROM tables WHERE id = ?
`

var sqlTableBookedOnDate = `
	SELECT t.id FROM tables t JOIN table_book b 
	ON t.id = b.table_id 
    WHERE t.id = ? and DATE(b.book_date) = DATE(?) 
`

var sqlCreateBook = `
	INSERT INTO table_book (id, table_id, book_date, guest_name, guest_contact, code)
	VALUES (?, ?, ?, ?, ?, ?)
`
