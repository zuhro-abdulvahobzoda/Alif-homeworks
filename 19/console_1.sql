CREATE TABLE IF NOT EXISTS books ( -- first task
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    author TEXT NOT NULL,
    price INT NOT NULL,
    pages INT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    CONSTRAINT unique_title_author UNIQUE (title, author)
);


INSERT INTO books (title, author, price, pages) VALUES -- second task
('The Hobbit', 'J.R.R. Tolkien', 499, 310),
('1984', 'George Orwell', 999, 328),
('To Kill a Mockingbird', 'Harper Lee', 1299, 281),
('The Great Gatsby', 'F.S. Fitzgerald', 899, 180),
('Dune', 'Frank Herbert', 1599, 617),
('The Catcher in the Rye', 'J.D. Salinger', 1099, 234),
('Brave New World', 'Aldous Huxley', 119, 268),
('Fahrenheit 451', 'Ray Bradbury', 1399, 256),
('Pride and Prejudice', 'Jane Austen', 799, 279),
('The Alchemist', 'Paulo Coelho', 1450, 208),
('Project Hail Mary', 'Andy Weir', 899, 476),
('Sapiens', 'Yuval Noah Harari', 1999, 512),
('Atomic Habits', 'James Clear', 420, 320),
('Educated', 'Tara Westover', 40, 352),
('The Silent Patient', 'Alex Michaelides', 1250, 336)
ON CONFLICT (title, author) DO NOTHING;

/* Full table */
SELECT * FROM books;
/* Title, author */
SELECT title, author FROM books;
/* Title, price */
SELECT title, price FROM books;

/* Find author */
SELECT title FROM books WHERE author = 'Robert Martin';
/*Price > 100*/
SELECT title FROM books WHERE price > 100;
/*Price < 50*/
SELECT title FROM books WHERE price < 50; -- price < 50
/*Pages > 500*/
SELECT title FROM books WHERE pages > 500; -- pages > 500
/*Price fixation */
SELECT title FROM books WHERE price > 100 AND price < 200; -- price 100 && 200

/*price AND pages*/
SELECT title FROM books WHERE price > 100 AND pages > 300;
/*price OR pages*/
SELECT title FROM books WHERE price < 100 OR pages < 200;
/*author AND price*/
SELECT * FROM books WHERE author = 'George Orwell' AND price > 150; -- find a certain author + price > 150


/*ASCEND*/
SELECT * FROM books ORDER BY price ASC;

/*DESCEND*/
SELECT * FROM books ORDER BY price DESC;

/*DESC by pages*/
SELECT * FROM books ORDER BY pages DESC;

/*Alphabetic order*/
SELECT * FROM books ORDER BY title;



/*Limit (task7)*/
SELECT * FROM books ORDER BY price DESC LIMIT 5;

/*Limit by pages*/
SELECT * FROM books ORDER BY pages DESC LIMIT 3;



UPDATE books SET price = 25 WHERE id = 1;
/*Update 1*/
SELECT * FROM books WHERE id = 1;

UPDATE books SET author = 'George Not Found' WHERE id = 2;
/*Update 2*/
SELECT * FROM books WHERE id = 2;

UPDATE books SET title = 'Go Set a Watchman', price = 1499, pages = 278 WHERE id = 3;
/*Update 3*/
SELECT * FROM books WHERE id = 3;


UPDATE books SET price = 150 WHERE price < 100;
/*Update-Where*/
SELECT * FROM books;


DELETE FROM books
WHERE id = 11;
/*Delete with id*/
SELECT * FROM books;

DELETE FROM books WHERE price < 50;
/*Delete with condition*/
SELECT * FROM books;

DELETE FROM books WHERE pages < 100;
/*Delete with condition pt2*/
SELECT * FROM books;


/*Count total books*/
SELECT COUNT(*) FROM books;

/*Count with condition*/
SELECT COUNT(*) FROM books WHERE price > 100;

/*Count with condition*/
SELECT COUNT(*) FROM books WHERE author = 'Robert Martin';

/*Like 1*/
SELECT * FROM books WHERE title LIKE 'Go';

/*Like 2*/
SELECT * FROM books WHERE author LIKE 'Luke_';