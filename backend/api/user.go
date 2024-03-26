package api

import (
	"log"
	"net/http"
	"strconv"

	"../models"
	"../utils"
	"github.com/gin-gonic/gin"
)

func (api *APIController) GetUserByID(c *gin.Context) {
	var user *models.UserWithID
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return
	}

	user, err = api.DBController.GetUserByID(int64(id))

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

	c.JSON(http.StatusOK, user)
}

func (api *APIController) GetBooksOwnedByUser(c *gin.Context) {

}

func (api *APIController) CreateUser(c *gin.Context) {
	var user models.User
	var createdUser *models.UserWithID
	var err error

	if err = c.BindJSON(&user); err != nil {
		log.Println(err)
		return
	}

	createdUser, err = api.DBController.InsertUser(user)

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

	c.JSON(http.StatusCreated, createdUser)
}

func (api *APIController) DeleteUserByID(c *gin.Context) {
	var err error
	id, err := strconv.Atoi(c.Param("id"))

	err = api.DBController.DeleteUser(int64(id))

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
