package main

import (
	"log"
	"my-app/internal/config"
	"my-app/internal/database"
	"my-app/internal/handler"
	"my-app/internal/respository"
	"my-app/internal/router"
	"my-app/internal/service"
	"net/http"
	"time"
)

func main() {
	cfg , err := config.Load()
	if err != nil { 
		log.Fatalf("Config Error : %v" , err)
	}

	db , err := database.ConnectDb(cfg.DSN())

	if err != nil { 
		log.Fatalf("Database Connection Error %v" , err)
	}

	if err := database.Migrate(db); err != nil { 
		log.Fatalf("Database Migration Error : %v" , err)
	}

	// Dependency Injection (Constructor Injection)
	bookRepository := respository.InitBookRepository(db)
	bookService    := service.InitBookService(bookRepository)
	bookHandler    := handler.InitBookHandler(bookService) 


	server := &http.Server { 
		Addr: ":" + cfg.Port,
		Handler: router.New(bookHandler),
		ReadHeaderTimeout: 5 * time.Second,
	}
 
	go func() { 
		log.Printf("App Running on localhost:%v" ,cfg.Port)
		if err := server.ListenAndServe(); err != nil { 
			log.Fatalf("Server Running Error : %v" , err)
		}
	}()
}