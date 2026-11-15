package service

import (
	"context"
	"my-app/internal/dto"
	"my-app/internal/models"
	"my-app/internal/respository"
	"strings"
)

type BookService interface {
	Create(ctx context.Context , req dto.CreateBookRequest) (*models.Book , error)
	List(ctx context.Context , query dto.ListBookQuery)([]models.Book , int , error)
	GetById(ctx context.Context , id uint)(*models.Book , error)
	Update(ctx context.Context , id uint , req dto.UpdateBookRequest)(*models.Book , error)
	Delete(ctx context.Context , id uint) error
}

type bookService struct { 
	repo respository.BookRespository
}

func InitBookService(repo respository.BookRespository) BookService{ 
	return &bookService{repo : repo}
}

func (s *bookService) Create(ctx context.Context , req dto.CreateBookRequest)(*models.Book , error) { 
	book := &models.Book{
		Title : strings.TrimSpace(req.Title) ,
		Price: req.Price,
		ISBN: req.ISBN,
		Author: strings.TrimSpace(req.Author),
	}
	if err := s.repo.Create(ctx , book); err != nil { 
		return nil , err
	}
	return book , nil 
}

func (s *bookService) List(ctx context.Context , query dto.ListBookQuery)([]models.Book , int , error) { 
	filter := respository.BookFilter{
		Author: strings.TrimSpace(query.Author),
	}
	return s.repo.FindAll(ctx , filter)
}

func (s *bookService) GetById(ctx context.Context , id uint)(*models.Book , error) { 
	return s.repo.FindById(ctx , id);
}

func (s *bookService) Update(ctx context.Context , id uint , req dto.UpdateBookRequest)(*models.Book , error) { 
	book , err := s.repo.FindById(ctx , id)
	if err != nil { 
		return nil , err
	}
	if req.Title != nil { 
		book.Title = strings.TrimSpace(*req.Title)
	}
	if req.Author != nil { 
		book.Author = strings.TrimSpace(*req.Author)
	}
	if req.Price != nil { 
		book.Price = *req.Price
	}

	if err := s.repo.Update(ctx , book); err != nil { 
		return  nil , err
	}
	return book , nil
}


func (s *bookService) Delete(ctx context.Context , id uint) error { 
	return s.repo.Delete(ctx , id)
}