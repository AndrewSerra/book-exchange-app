package db

import (
	"database/sql"
	"log"

	"../models"
	"../utils"
	"github.com/go-sql-driver/mysql"
)

// Gets a user from the database
// Returns a User object
func (c *DBController) GetUserByID(uid int64) (*models.UserWithID, error) {
	var user = new(models.UserWithID)
	var err error

	db := c.Database

	row := db.QueryRow(
		"SELECT id, firstName, lastName, DATE_FORMAT(dob,'%Y-%m-%d'), email FROM Users WHERE id = ?;",
		uid)

	if err = row.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Dob, &user.Email); err != nil {
		log.Println(err)
		if err == sql.ErrNoRows {
			return nil, &utils.DataNotFoundError{
				Data: map[string]int64{"UserId": uid},
			}
		}
		return nil, &utils.UnknownError{
			Err: err,
		}
	}
	return user, nil
}

// Creates a user in the database
// Takes the first name, last name, date of birth,
// and email of the user.
func (c *DBController) InsertUser(user models.User) (*models.UserWithID, error) {
	var newUser = new(models.UserWithID)
	var err error

	db := c.Database

	result, err := db.Exec(
		"INSERT INTO Users (firstName, lastName, dob, email) VALUES (?, ?, ?, ?);",
		user.FirstName, user.LastName, user.Dob, user.Email,
	)

	if err != nil {
		if mysqlerr, ok := err.(*mysql.MySQLError); ok {
			if mysqlerr.Number == utils.MYSQL_DUPLICATE_ERROR {
				return nil, &utils.DataExistsError{
					Data: map[string]models.User{
						"User": user,
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

	newUser = &models.UserWithID{
		ObjectWithID: models.ObjectWithID{
			Id: id,
		},
		User: user,
	}
	return newUser, nil
}

// Deletes a user in the database
// Takes the user id to delete.
func (c *DBController) DeleteUser(uid int64) error {
	var err error

	db := c.Database

	result, err := db.Exec("DELETE FROM Users WHERE id = ?;", uid)

	if err != nil {
		log.Println(err.Error())
		return err
	} else if count, _ := result.RowsAffected(); count == 0 {
		return &utils.DataNotFoundError{
			Data: models.ObjectWithID{
				Id: uid,
			},
			Err: err,
		}
	}

	return nil
}

// Updates the user email in the database
// Takes the user id and the new email.
func (c *DBController) UpdateEmail(uid int64, newEmail string) {}

// Marks the user email as verified in the database
// Takes user id as input.
func (c *DBController) VerifyEmail(uid int64) {}
