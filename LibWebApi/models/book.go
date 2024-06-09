package models

import (
	"LibWebApi/db"
	"database/sql"
)

type Book struct {
	Id            int64
	Title         string
	ISBN          string
	Author        string
	PublishedYear int
}

var bookCollection []Book

func (m *Book) Save() error {

	statment, err := db.GetDb().Prepare(`
		INSERT INTO 
		books
		    (title, isbn, author, publishedyear)
		VALUES
		    (?, ?, ?,?)
	`)

	defer statment.Close()

	if err != nil {
		return err
	}

	result, err := statment.Exec(m.Title, m.ISBN, m.Author, m.PublishedYear)
	if err != nil {
		return err
	}

	bookId, err := result.LastInsertId()
	m.Id = bookId

	return err
}

func GetAllBooks() ([]Book, error) {

	dbCursor, err := db.GetDb().Query(`SELECT * FROM books`)
	if err != nil {
		return nil, err
	}

	for dbCursor.Next() {

		var bookObject Book
		err := dbCursor.Scan(
			&bookObject.Id,
			&bookObject.Title,
			&bookObject.ISBN,
			&bookObject.Author,
			&bookObject.PublishedYear,
		)

		if err != nil {
			return nil, err
		}

		bookCollection = append(bookCollection, bookObject)
	}

	return bookCollection, nil
}

func GetBook(id int64) (Book, error) {
	var book Book
	err := db.GetDb().QueryRow("SELECT id, title, isbn, author, publishedyear FROM books WHERE id = ?", id).Scan(
		&book.Id, &book.Title, &book.ISBN, &book.Author, &book.PublishedYear)
	if err != nil {
		return book, err
	}
	return book, nil
}

func createBook(db *sql.DB, book *Book) error {
	err := db.QueryRow(
		"INSERT INTO books (title, isbn, author, publishedyear) VALUES ($1, $2, $3, $4) RETURNING id",
		book.Title, book.ISBN, book.Author, book.PublishedYear).Scan(&book.Id)
	return err
}

func (m *Book) Update() error {
	_, err := db.GetDb().Exec(
		"UPDATE books SET title = ?, isbn = ?, author = ?, publishedyear = ? WHERE id = ?",
		m.Title, m.ISBN, m.Author, m.PublishedYear, m.Id)
	return err
}

func DeleteBook(id int64) error {
	_, err := db.GetDb().Exec("DELETE FROM books WHERE id = ?", id)
	return err
}
