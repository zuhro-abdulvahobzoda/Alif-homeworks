CREATE TABLE IF NOT EXISTS book (
    id SERIAL PRIMARY KEY,
    title TEXT,
    author TEXT,
    price INT,
    pages INT,
    created_at TIMESTAMP DEFAULT NOW()
);

INSERT INTO book (title, author, price, pages) VALUES
('The Alchemist', 'Paulo Coelho', 114, 197),
('Jane Eyre', 'Charlotte Bronte', 192, 500),
('The Chronicles of Narnia', 'C.S. Lewis', 250, 768),
('Brave New World', 'Aldous Huxley', 130, 268),
('The Odyssey', 'Homer', 450, 544),
('Crime and Punishment', 'Fyodor Dostoevsky', 360, 671),
('The Little Prince', 'Antoine de Saint-Exupéry', 123, 96),
('The Picture of Dorian Gray', 'Oscar Wilde', 490, 254),
('Frankenstein', 'Mary Shelley', 141, 280),
('The Da Vinci Code', 'Dan Brown', 617, 489);

/*Full table*/
SELECT * FROM book;

/*Title, author*/
SELECT title, author FROM book;

/*Fifth book*/
SELECT * FROM book WHERE id = 5;

UPDATE book SET price = 350 WHERE id = 3;
/*Update 3rd book price*/
SELECT * FROM book;

UPDATE book SET price = price + 100;
/*Add 100*/
SELECT * FROM book;

UPDATE book SET pages = 500 WHERE id = 6;
/*Change 6th pages value*/
SELECT * FROM book;

DELETE FROM book WHERE id = 8;
/*Deleted book*/
SELECT * FROM book;

DELETE FROM book WHERE price < 250;
/*Deleted book*/
SELECT * FROM book;

/*Select certain books*/
SELECT * FROM book WHERE price > 400;

/*Select again*/
SELECT * FROM book WHERE pages > 300;

/*Sorting*/
SELECT * FROM book ORDER BY price DESC;

/*Sorting AGAIN*/
SELECT * FROM book LIMIT 5;
