package models

import "time"

type ObjectWithID struct {
	Id int `json:"id"`
}

type Book struct {
	Title   string    `json:"title"`
	Author  string    `json:"author"`
	PubDate time.Time `json:"pubDate"`
	Lang    string    `json:"lang"`
}

type BookWithID struct {
	ObjectWithID
	Book
}

type User struct {
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Dob       time.Time `json:"dob"`
	Email     string    `json:"email"`
}

type UserWithID struct {
	ObjectWithID
	User
}

type Review struct {
	BookId    int `json:"bookId"`
	UserId    int `json:"userId"`
	ExhangeId int `json:"exchangeId"`
	Rating    int `json:"rating"`
}

type ReviewWithID struct {
	ObjectWithID
	Review
}

type Address struct {
	Addr1      string `json:"address"`
	Addr2      string `json:"address2"`
	District   string `json:"district"`
	City       string `json:"city"`
	Country    string `json:"country"`
	PostalCode string `json:"postalCode"`
	Default    bool   `json:"isDefault"`
}

type AddressWithID struct {
	ObjectWithID
	Address
}
