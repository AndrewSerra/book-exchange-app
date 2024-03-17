package models

import "../utils"

type ObjectWithID struct {
	Id int64 `json:"id"  binding:"required"`
}

type Book struct {
	Title   string              `json:"title" binding:"required"`
	Author  string              `json:"author" binding:"required"`
	PubDate utils.MonthYearDate `json:"pubDate" binding:"required"`
	Lang    string              `json:"lang" binding:"required"`
	ISBN    string              `json:"isbn" binding:"required"`
}

type BookWithID struct {
	ObjectWithID
	Book
}

type User struct {
	FirstName string         `json:"firstName" binding:"required"`
	LastName  string         `json:"lastName" binding:"required"`
	Dob       utils.FullDate `json:"dob" binding:"required"`
	Email     string         `json:"email" binding:"required"`
}

type UserWithID struct {
	ObjectWithID
	User
}

type Review struct {
	BookId    int `json:"bookId" binding:"required"`
	UserId    int `json:"userId" binding:"required"`
	ExhangeId int `json:"exchangeId" binding:"required"`
	Rating    int `json:"rating" binding:"required"`
}

type ReviewWithID struct {
	ObjectWithID
	Review
}

type Address struct {
	Addr1      string `json:"address" binding:"required"`
	Addr2      string `json:"address2"`
	District   string `json:"district" binding:"required"`
	City       string `json:"city" binding:"required"`
	Country    string `json:"country" binding:"required"`
	PostalCode string `json:"postalCode" binding:"required"`
	Default    bool   `json:"isDefault" binding:"required"`
}

type AddressWithID struct {
	ObjectWithID
	Address
}
