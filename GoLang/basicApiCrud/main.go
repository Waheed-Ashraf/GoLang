package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"
	"sync"
	"time"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Auhtor string `json:"author"`
	Price  int    `json:"price"`
}

type Store struct {
	mu     sync.RWMutex // allow many reader but only write
	books  map[int]Book // books map (nil map cause panic)
	nextId int          // auto increment counter
}

// Initialization Constuctor Share Same Store With Id Starter = 1
func NewStore() *Store {
	return &Store{books: make(map[int]Book), nextId: 1}
}

func (s *Store) create(book Book) Book {
	s.mu.Lock()
	defer s.mu.Unlock()
	book.ID = s.nextId
	s.nextId++
	s.books[book.ID] = book
	return book
}

// [{ID : 5} , {ID : 2} , {ID : 4}]
// cmp.compare()
func (s *Store) List() []Book {
	s.mu.RLock()
	defer s.mu.RUnlock()
	booksSlice := make([]Book, 0, len(s.books))

	for _, book := range s.books {
		booksSlice = append(booksSlice, book)
	}
	// Sort
	sort.Slice(booksSlice, func(i, j int) bool {
		return booksSlice[i].ID < booksSlice[j].ID
	})

	return booksSlice
}

func (s *Store) GetById(id int) (Book, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	book, ok := s.books[id]
	return book, ok
}

func (s *Store) Update(id int, book Book) (Book, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.books[id]; !ok {
		return Book{}, false
	}
	book.ID = id
	s.books[id] = book
	return book, true
}

func (s *Store) Delete(id int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.books[id]; !ok {
		return false
	}
	delete(s.books, id)
	return true
}

// Write JSon
func writeJSON(w http.ResponseWriter, status int, val any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(val); err != nil {
		log.Printf("Something Wrong While Encoding %v ", err)
	}
}

// Write Error
func writError(w http.ResponseWriter, status int, msg string) {
	writeJSON(w, status, map[string]string{"errorMessage": msg,
		"status": "error"})
}

func main() {
	mux := http.NewServeMux()

	srv := http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  8 * time.Second,
		WriteTimeout: 8 * time.Second,
	}

	fmt.Println("Server Running on localhost:8080")
	log.Fatal(srv.ListenAndServe())
}
