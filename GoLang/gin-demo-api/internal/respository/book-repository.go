package respository

import (
	"context"
	"errors"
	"my-app/internal/apperrors"
	"my-app/internal/models"
	"gorm.io/gorm"
)

type BookFilter struct { 
	Author string
}

type BookRespository interface {
	Create(ctx context.Context , book *models.Book) error
	FindById(ctx context.Context, id uint) (*models.Book , error)
	FindAllBook(ctx context.Context)([]models.Book , error)
	FindAll(ctx context.Context , filter BookFilter)([]models.Book , int , error)
	Update(ctx context.Context , book *models.Book) error
	Delete(ctx context.Context , id uint) error
}

type bookRepository struct { 
	db *gorm.DB
}

func InitBookRepository(db *gorm.DB) BookRespository{ 
	return &bookRepository{db : db}
}

func (r *bookRepository) Create(ctx context.Context , book *models.Book) error { 
	err := r.db.WithContext(ctx).Create(book).Error
	return err
}

func (r *bookRepository) FindAllBook(ctx context.Context) ([]models.Book , error) {
	var books []models.Book
	query := r.db.WithContext(ctx).Model(&models.Book{})
	query.Find(&books)
	return books , nil
}

func (r *bookRepository) FindAll(ctx context.Context , filter BookFilter)([]models.Book , int , error) { 
	query := r.db.WithContext(ctx).Model(&models.Book{})

	if filter.Author != "" { 
		query.Where("author ILIKE ?" , "%" + filter.Author + "%")
	}

	var total int64
	var books []models.Book

	if err := query.Count(&total).Error; err != nil {
		return nil , 0, err
	}

	if err := query.Find(&books).Error; err != nil { 
		return nil , 0 , err
	}

	return books , int(total) , nil
	
}

func (r *bookRepository) FindById(ctx context.Context , id uint) (*models.Book, error) { 
	var book models.Book
	if err := r.db.WithContext(ctx).First(&book , id).Error; err != nil { 
		return nil , translateError(err)
	}
	return &book , nil
}

func (r *bookRepository) Update(ctx context.Context , book *models.Book) error { 
	err := r.db.WithContext(ctx).Save(book).Error
	return translateError(err)
}

func (r *bookRepository) Delete(ctx context.Context , id uint) error { 
	result := r.db.WithContext(ctx).Delete(&models.Book{} , id)

	if result.Error != nil {
		return  result.Error
	}

	if result.RowsAffected == 0 { 
		return apperrors.ErrNotFound
	}

	return nil
}


func translateError(err error) error { 
	switch {
		case errors.Is(err , gorm.ErrRecordNotFound):
			return apperrors.ErrNotFound
		case errors.Is(err , gorm.ErrDuplicatedKey):
			return apperrors.ErrConflict
		default:
			return err
	}
}