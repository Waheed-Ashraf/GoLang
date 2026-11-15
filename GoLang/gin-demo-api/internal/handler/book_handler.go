package handler

import (
	"errors"
	"my-app/internal/apperrors"
	"my-app/internal/dto"
	"my-app/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	service service.BookService
}

func InitBookHandler(service service.BookService) *BookHandler { 
	return &BookHandler{service: service}
}

func (h *BookHandler) CreateBook(c *gin.Context) { 
	var req dto.CreateBookRequest

	if err := c.ShouldBindJSON(&req); err != nil { 
		c.JSON(http.StatusBadRequest , gin.H { 
			"error" : err.Error(),
		})
		return
	}

	book , err := h.service.Create(c.Request.Context() , req)

	if err != nil { 
		respondError(c , err)
		return
	}

	c.JSON(http.StatusCreated , gin.H { 
		"status" : "created" , 
		"book" : dto.DisplayBookResponse(book),
	})

} 


func (h *BookHandler) ListBooks(c *gin.Context) { 
	var query dto.ListBookQuery

	if err := c.ShouldBindQuery(&query); err != nil { 
		c.JSON(http.StatusBadRequest , gin.H {
			"error" : err.Error(),
		})
		return
	}

	books , total , err := h.service.List(c.Request.Context() , query)

	if err != nil { 
		respondError(c , err)
		return
	}

	c.JSON(http.StatusOK , gin.H { 
		"status" : "ok" , 
		"response" : dto.BookListResponse{
			Data: dto.DisplayBooksResponses(books),
			Total: int64(total),
		},
	})
}


func (h *BookHandler) GetBookById(c *gin.Context) { 
	id , ok := parseId(c)
	if(!ok) {return}

	book , err := h.service.GetById(c.Request.Context() , id)
	if err != nil { 
		respondError(c , err)
		return
	}
	c.JSON(http.StatusOK , dto.DisplayBookResponse(book))
}


func (h *BookHandler) Update(c *gin.Context) { 
	id , ok := parseId(c)
	if !ok { return }

	var req dto.UpdateBookRequest

	if err := c.ShouldBindJSON(&req); err != nil { 
		c.JSON(http.StatusBadRequest , gin.H{ 
			"error" : err.Error(),
		})
		return
	}

	book , err := h.service.Update(c.Request.Context() , id , req)
	if err != nil { 
		respondError(c , err)
		return
	}

	c.JSON(http.StatusOK , dto.DisplayBookResponse(book))
}


func (h *BookHandler) Delete(c *gin.Context) { 
	id , ok := parseId(c)
	if !ok {return}

	if err := h.service.Delete(c.Request.Context() , id); err != nil { 
		respondError(c , err)
		return
	}

	c.Status(http.StatusNoContent)
}


func parseId(c *gin.Context) (uint , bool) { 
	id , err := strconv.ParseUint(c.Param("id") , 10 , 64)

	if err != nil || id <= 0 { 
		c.JSON(http.StatusBadRequest , gin.H { 
			"error" : "Invalid id value",
		})
		return 0 , false
	}
	return uint(id) , true
}


func respondError(c *gin.Context , err error) { 
	switch {
	case errors.Is(err , apperrors.ErrNotFound):
		c.JSON(http.StatusNotFound , gin.H {
			"error" : "Book Not Found",
		})
	case errors.Is(err , apperrors.ErrConflict):
		c.JSON(http.StatusConflict , gin.H { 
			"error" : "Book Already Exists",
		})
	default :
		c.JSON(http.StatusInternalServerError , gin.H { 
			"error" : err.Error(),
		})
	}
}