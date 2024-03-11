package db

import (
	"time"

	"../models"
)

// Creates a book in the database
// Takes the book title, author name, publication date,
// and language.
func (c *DBController) InsertBook(title string, author string, pubDate time.Time, lang string) {

}

// Deletes a book in the database
// Takes the book id of the book to delete.
func (c *DBController) DeleteBook(uid int) {}

// Gets the book with id in the database
// Returns the book object
func (c *DBController) GetBook(bid int) *models.Book {
	return nil
}

// Get the books that were exchanged by the user
// Returns the book list the user
func (c *DBController) GetUserBooks(uid int) []*models.Book {
	return []*models.Book{}
}
