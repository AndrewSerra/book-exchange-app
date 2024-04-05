package api

import (
	"log"
	"net/http"
	"strconv"

	"github.com/AndrewSerra/book-exchange/models"
	"github.com/AndrewSerra/book-exchange/utils"
	"github.com/gin-gonic/gin"
)

func (api *APIController) GetBookByID(c *gin.Context) {
	var book *models.BookWithID
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return
	}

	book, err = api.DBController.GetBookByID(int64(id))

	if err != nil {
		switch err.(type) {
		case *utils.DataNotFoundError:
			c.Status(http.StatusNotFound)
			return
		case *utils.UnknownError:
		default:
			c.Status(http.StatusInternalServerError)
			return
		}
	}

	c.JSON(http.StatusOK, book)
}

func (api *APIController) GetUsersConnectedToBook(c *gin.Context) {
	var users []models.User
	id, err := strconv.Atoi((c.Param("id")))

	if err != nil {
		log.Println(err)
		c.Status(400)
		return
	}

	users, err = api.DBController.GetUsersInteractedWithBook(int64(id))

	if err != nil {
		switch err.(type) {
		case *utils.QueryProcessingError:
		case *utils.UnknownError:
		default:
			c.Status(http.StatusInternalServerError)
			return
		}
	}

	c.JSON(http.StatusOK, users)
}

func (api *APIController) CreateBook(c *gin.Context) {
	var book models.Book
	var createdBook *models.BookWithID
	var err error

	if err = c.BindJSON(&book); err != nil {
		log.Println(err)
		return
	}

	createdBook, err = api.DBController.InsertBook(book)

	if err != nil {
		switch err.(type) {
		case *utils.DataExistsError:
			c.Status(http.StatusConflict)
			return
		case *utils.UnknownError:
		default:
			c.Status(http.StatusInternalServerError)
			return
		}
	}

	c.JSON(http.StatusCreated, createdBook)
}

func (api *APIController) DeleteBookByID(c *gin.Context) {
	var err error
	id, _ := strconv.Atoi(c.Param("id"))

	err = api.DBController.DeleteBookByID(int64(id))

	if err != nil {
		switch err.(type) {
		case *utils.DataNotFoundError:
			c.Status(http.StatusNotFound)
			return
		case *utils.UnknownError:
		default:
			c.Status(http.StatusInternalServerError)
			return
		}
	}

	c.Status(http.StatusNoContent)
}
