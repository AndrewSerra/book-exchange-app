package api

import (
	"net/http"
	"strconv"

	"../models"
	"../utils"
	"github.com/gin-gonic/gin"
)

func (api *APIContoller) GetUserByID(c *gin.Context) {
	var user models.UserWithID
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return
	}

	user, err = api.DBController.GetUserByID(id)

	if err != nil {
		switch err.(type) {
		case *utils.DataNotFoundError:
			c.JSON(http.StatusNotFound, user)
		case *utils.UnknownError:
		default:
			c.JSON(http.StatusInternalServerError, user)
		}
	}

	c.JSON(http.StatusOK, user)
}

func (api *APIContoller) CreateUser(c *gin.Context) {

}
