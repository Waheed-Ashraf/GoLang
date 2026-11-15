package dto

import "my-app/internal/models"

// Json Body (POST /api/v1/books)
type CreateBookRequest struct { 
	Title string `json:"title" binding:"required,max=200"`
	Author string `json:"author" binding:"required,max=100"`
	ISBN string `json:"isbn" binding:"required , min=8 ,max=20"`
	Price float64 `json:"price" binding:"required , gt=0"`
} 

type UpdateBookRequest struct { 
	Title *string `json:"title" binding:"omitempty,max=200"`
	Author *string `json:"author" binding:"omitempty,max=100"`
	ISBN *string `json:"isbn" binding:"omitempty,min=8 ,max=20"`
	Price *float64 `json:"price" binding:"omitempty,gt=0"`
}

type BookResponse struct { 
	ID uint `json:"id"`
	Title string `json:"title"`
	Price float64 `json:"price"`
	Author string `json:"author"`
	ISBN string `json:"isbn"`
}

// ?page=1&page_size=2
type ListBookQuery struct { 
	// Page int `form:"page,default=1" binding:"gte=1"`
	// PageSize int `form:"page_size,default=12" binding:"gte=1"`
	Author string `form:"author" binding:"max=100"`
}

type BookListResponse struct { 
	Data []BookResponse `json:"data"`
	Total int64 `json:"total"`
	// Page int `json:"page"`
	// PageSize int `json:"page_size"`
}


func DisplayBookResponse(book *models.Book) BookResponse {
	return BookResponse{
		ID: book.ID,
		Title: book.Title,
		Price: book.Price,
		Author: book.Author,
		ISBN: book.Author,
	}
}


func DisplayBooksResponses(books []models.Book) []BookResponse { 
	res := make([]BookResponse , 0 , len(books))

	for i := range books { 
		res = append(res, DisplayBookResponse(&books[i]))
	}

	return res
}