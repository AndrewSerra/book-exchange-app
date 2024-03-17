use books;

ALTER TABLE Books
    ADD COLUMN isbn VARCHAR(32) NOT NULL;
