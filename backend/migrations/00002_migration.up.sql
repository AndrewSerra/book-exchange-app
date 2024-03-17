use books;

ALTER TABLE Books
    ADD COLUMN isbn VARCHAR(16) NOT NULL;
