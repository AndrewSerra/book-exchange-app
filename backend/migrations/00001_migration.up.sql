use books;

CREATE TABLE IF NOT EXISTS Books (
    id         INT AUTO_INCREMENT NOT NULL,
    title      VARCHAR(128) NOT NULL,
    author     VARCHAR(255) NOT NULL,
    genre      VARCHAR(255) NOT NULL,
    pubDate    DATE NOT NULL,
    lang       VARCHAR(2) NOT NULL,
    createdAt TIMESTAMP NOT NULL DEFAULT NOW(),
    updatedAt TIMESTAMP NOT NULL DEFAULT NOW() ON UPDATE now(),
    PRIMARY KEY (id)  
);

CREATE TABLE IF NOT EXISTS Users (
    id         INT AUTO_INCREMENT NOT NULL,
    firstName VARCHAR(255) NOT NULL,
    lastName VARCHAR(255) NOT NULL,
    dob DATE NOT NULL,
    email VARCHAR(255) NOT NULL,
    createdAt TIMESTAMP NOT NULL DEFAULT NOW(),
    updatedAt TIMESTAMP NOT NULL DEFAULT NOW() ON UPDATE now(),
    PRIMARY KEY (id),
    CONSTRAINT UQEmail UNIQUE (email)
);

CREATE TABLE IF NOT EXISTS Exchanges (
    id           INT AUTO_INCREMENT NOT NULL,
    bookId INT NOT NULL,
    fromUserId   INT NOT NULL,
    toUserId     INT NOT NULL,
    isReceived BIT NOT NULL DEFAULT 0,
    createdAt TIMESTAMP NOT NULL DEFAULT NOW(),
    updatedAt TIMESTAMP NOT NULL DEFAULT NOW() ON UPDATE now(),
    PRIMARY KEY (id),
    FOREIGN KEY (bookId) REFERENCES Books(id),
    FOREIGN KEY (fromUserId) REFERENCES Users(id),
    FOREIGN KEY (toUserId) REFERENCES Users(id)
);

CREATE TABLE IF NOT EXISTS Reviews (
    id INT AUTO_INCREMENT NOT NULL,
    bookId INT NOT NULL,
    userId INT NOT NULL,
    exchangeId INT NOT NULL,
    rating INT CHECK (rating BETWEEN 0 AND 5),
    createdAt TIMESTAMP NOT NULL DEFAULT NOW(),
    updatedAt TIMESTAMP NOT NULL DEFAULT NOW() ON UPDATE now(),
    PRIMARY KEY (id),
    FOREIGN KEY (bookId) REFERENCES Books(id),
    FOREIGN KEY (userId) REFERENCES Users(id),
    FOREIGN KEY (exchangeId) REFERENCES Exchanges(id)
);

CREATE TABLE IF NOT EXISTS Addresses (
    id INT AUTO_INCREMENT NOT NULL,
    userId INT NOT NULL,
    isDefault BIT NOT NULL DEFAULT 0,
    address VARCHAR(255) NOT NULL,
    address2 VARCHAR(255),
    district VARCHAR(128) NOT NULL,
    city VARCHAR(128) NOT NULL,
    country VARCHAR(128) NOT NULL,
    postalCode VARCHAR(128) NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (userId) REFERENCES Users(id)
);
