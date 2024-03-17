package db

import (
	"database/sql"
	"log"
	"strings"

	"../models"
	"../utils"
	"github.com/go-sql-driver/mysql"
)

// Creates a book in the database
// Takes the book title, author name, publication date,
// and language.
func (c *DBController) InsertBook(book models.Book) (*models.BookWithID, error) {
	var newBook *models.BookWithID
	var err error
	var genre string = strings.Join(book.Genre, ",")

	db := c.Database

	result, err := db.Exec(
		"INSERT INTO Books (title, author, pubDate, genre, lang, isbn) VALUES (?, ?, ?, ?, ?, ?);",
		book.Title, book.Author, book.PubDate, genre, book.Lang, book.ISBN,
	)

	if err != nil {
		if mysqlerr, ok := err.(*mysql.MySQLError); ok {
			if mysqlerr.Number == utils.MYSQL_DUPLICATE_ERROR {
				return nil, &utils.DataExistsError{
					Data: map[string]models.Book{
						"Book": book,
					},
					Err: mysqlerr,
				}
			}
			return nil, &utils.UnknownError{
				Err: err,
			}
		}
		return nil, &utils.UnknownError{
			Err: err,
		}
	}

	id, err := result.LastInsertId()

	if err != nil {
		return nil, &utils.UnknownError{
			Err: err,
		}
	}

	newBook = &models.BookWithID{
		ObjectWithID: models.ObjectWithID{
			Id: id,
		},
		Book: book,
	}
	return newBook, nil
}

// Deletes a book in the database
// Takes the book id of the book to delete.
func (c *DBController) DeleteBookByID(bid int64) error {
	var err error

	db := c.Database

	result, err := db.Exec("DELETE FROM Books WHERE id = ?;", bid)

	if err != nil {
		log.Println(err.Error())
		return err
	} else if count, _ := result.RowsAffected(); count == 0 {
		return &utils.DataNotFoundError{
			Data: models.ObjectWithID{
				Id: bid,
			},
			Err: err,
		}
	}

	return nil
}

// Gets the book with id in the database
// Returns the book object
func (c *DBController) GetBookByID(bid int64) (*models.BookWithID, error) {
	var book = new(models.BookWithID)
	var err error

	db := c.Database

	row := db.QueryRow(
		"SELECT id, title, author, pubDate, genre, lang, isbn FROM Books WHERE id = ?;",
		bid)

	if err = row.Scan(&book.Id, &book.Title, &book.Author, &book.PubDate, &book.Genre, &book.Lang, &book.ISBN); err != nil {
		log.Println(err)
		if err == sql.ErrNoRows {
			return nil, &utils.DataNotFoundError{
				Data: map[string]int64{"BookId": bid},
			}
		}
		return nil, &utils.UnknownError{
			Err: err,
		}
	}
	return book, nil
}

// Get the books that were exchanged by the user
// Returns the book list the user
func (c *DBController) GetUserBooks(uid int64) []*models.Book {
	return []*models.Book{}
}
