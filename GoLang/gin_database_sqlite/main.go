package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "modernc.org/sqlite" // Blank Import
)

// Product Model
type Product struct {
	ID     int64  `json:"id"`
	Title  string `json:"title" binding:"required"`
	Price  int    `json:"price" binding:"required"`
	Vendor string `json:"vendor" binding:"required"`
}

// Server Struct Shared By Handlers
type Server struct {
	db *sql.DB
}

func main() {
	db, err := sql.Open("sqlite", "file:products.db?_pragma=busy_timeout(4000)")
	if err != nil {
		log.Fatalf("Database Connection Failed : %v", err)
	}
	defer db.Close()

	db.SetMaxOpenConns(1) // Sqlite Allow One Writer At A Time

	if err := db.Ping(); err != nil {
		log.Fatalf("Ping Error %v ", err)
	}

	if err := migrate(db); err != nil {
		log.Fatalf("Migration Error %v", err)
	}

	router := gin.Default()

	s := &Server{db: db}

	productRouter := router.Group("/api/v1/products")
	{
		productRouter.GET("", s.listProducts)
		productRouter.POST("", s.createProduct)
		productRouter.GET("/:id", s.getProduct)
	}

	router.Run(":8080")
}

// parse id helper function
func parseID(c *gin.Context) (int64, error) {
	id := c.Param("id")
	if id == "" {
		return 0, fmt.Errorf("id is required")
	}
	return strconv.ParseInt(id, 10, 64)
}

func migrate(db *sql.DB) error {
	const schema = `
        CREATE TABLE IF NOT EXISTS products(
            id INTEGER PRIMARY KEY AUTOINCREMENT ,
            title TEXT NOT NULL ,
            price INTEGER NOT NULL ,
            vendor TEXT NOT NULL
        )
    `
	_, err := db.Exec(schema)
	return err
}

func (s *Server) createProduct(c *gin.Context) {
	var product Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	res, err := s.db.ExecContext(c.Request.Context(), `
        INSERT INTO products(title , price , vendor) VALUES(? , ? , ?)
    `, product.Title, product.Price, product.Vendor)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}

	id, _ := res.LastInsertId()

	product.ID = id

	c.JSON(http.StatusCreated, gin.H{
		"status":        "ok",
		"storedProduct": product,
	})
}

func (s *Server) listProducts(c *gin.Context) {
	rows, err := s.db.QueryContext(c.Request.Context(), `
        SELECT id , title , price , vendor FROM products`)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	defer rows.Close()

	products := make([]Product, 0)

	for rows.Next() {
		var p Product
		if err := rows.Scan(&p.ID, &p.Title, &p.Price, &p.Vendor); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		products = append(products, p)
	}

	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":        "Ok",
		"totalProducts": len(products),
		"products":      products,
	})
}

func (s *Server) getProduct(c *gin.Context) {
	id, err := parseID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var product Product
	err = s.db.QueryRowContext(c.Request.Context(), `
        SELECT id , title , price , vendor FROM products WHERE id = ?
    `, id).Scan(&product.ID, &product.Title, &product.Price, &product.Vendor)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "Ok",
		"product": product,
	})
}
