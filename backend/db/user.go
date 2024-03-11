package db

import (
	"database/sql"
	"time"

	"../models"
	"../utils"
)

// Gets a user from the database
// Returns a User object
func (c *DBController) GetUserByID(uid int) (models.UserWithID, error) {
	var user models.UserWithID
	var err error

	db := c.Database

	row := db.QueryRow(
		"SELECT id, firstName, lastName, dob, email from Users WHERE id = ?;",
		uid)

	if err = row.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Dob, &user.Email); err != nil {
		if err == sql.ErrNoRows {
			return user, &utils.DataNotFoundError{
				Data: map[string]int{"UserId": uid},
			}
		}
		return user, &utils.UnknownError{
			Err: err,
		}
	}

	return user, nil
}

// Creates a user in the database
// Takes the first name, last name, date of birth,
// and email of the user.
func (c *DBController) InsertUser(firstName string, lastName string, email string, date time.Time) {}

// Deletes a user in the database
// Takes the user id of the user to delete.
func (c *DBController) DeleteUser(uid int) {}

// Updates the user email in the database
// Takes the user id and the new email.
func (c *DBController) UpdateEmail(uid int, newEmail string) {}

// Marks the user email as verified in the database
// Takes user id as input.
func (c *DBController) VerifyEmail(uid int) {}
