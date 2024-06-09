package main

import (
	"LibWebApi/db"
	"LibWebApi/models"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Response map[string]any

func main() {

	db.Init()

	app := gin.Default()

	app.GET("/books", func(context *gin.Context) {
		fmt.Println("Hello Book")

		result, err := models.GetAllBooks()
		if err != nil {
			context.JSON(400, Response{
				"message": "Cannot serve your request",
			})
			return
		}

		context.JSON(200, Response{
			"message": "All books in the database",
			"books":   result,
		})
	})

	app.POST("/books", func(context *gin.Context) {

		var bookObject models.Book
		err := context.ShouldBindJSON(&bookObject)
		if err != nil {
			context.JSON(400, Response{
				"message": "Invalid object",
			})
			return
		}

		err = bookObject.Save()

		if err != nil {
			context.JSON(400, Response{
				"message": "Cannot insert book object",
			})
			return
		}

		context.JSON(200, Response{
			"message": "Book created successfuly",
			"object":  bookObject,
		})
	})

	app.GET("/books/:id", func(context *gin.Context) {
		id, err := strconv.Atoi(context.Param("id"))
		if err != nil {
			context.JSON(400, Response{
				"message": "Invalid book ID",
			})
			return
		}

		book, err := models.GetBook(int64(id))
		if err != nil {
			context.JSON(404, Response{
				"message": "Book not found",
			})
			return
		}

		context.JSON(200, Response{
			"message": "Book found",
			"book":    book,
		})
	})

	app.PUT("/books/:id", func(context *gin.Context) {
		id, err := strconv.Atoi(context.Param("id"))
		if err != nil {
			context.JSON(400, Response{
				"message": "Invalid book ID",
			})
			return
		}

		var bookObject models.Book
		err = context.ShouldBindJSON(&bookObject)
		if err != nil {
			context.JSON(400, Response{
				"message": "Invalid object",
			})
			return
		}
		bookObject.Id = int64(id)

		err = bookObject.Update()
		if err != nil {
			context.JSON(400, Response{
				"message": "Cannot update book object",
			})
			return
		}

		context.JSON(200, Response{
			"message": "Book updated successfully",
			"object":  bookObject,
		})
	})

	app.DELETE("/books/:id", func(context *gin.Context) {
		id, err := strconv.Atoi(context.Param("id"))
		if err != nil {
			context.JSON(400, Response{
				"message": "Invalid book ID",
			})
			return
		}

		err = models.DeleteBook(int64(id))
		if err != nil {
			context.JSON(400, Response{
				"message": "Cannot delete book object",
			})
			return
		}

		context.JSON(200, Response{
			"message": "Book deleted successfully",
		})
	})

	err := app.Run(":8087")
	if err != nil {
		fmt.Println("SERVER exception")
		fmt.Println(err)
	}
}
