use books;

ALTER TABLE Books 
    ADD CONSTRAINT UQ_ISBN UNIQUE (isbn);
