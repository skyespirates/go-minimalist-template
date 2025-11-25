package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	"github.com/skyespirates/go-minimalist-template/internal/delivery/http/handler"
	"github.com/skyespirates/go-minimalist-template/internal/infra/pgsql"
	"github.com/skyespirates/go-minimalist-template/internal/logger"
	"github.com/skyespirates/go-minimalist-template/internal/usecase"
)

type application struct {
	logger *logger.Logger
}

func main() {
	godotenv.Load()

	db, err := pgsql.InitDb()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	log.Println("database connection pool established")

	logger := logger.New(os.Stdout)

	app := &application{
		logger: logger,
	}

	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World! 🥶"))
	})

	handler.NewTaskHandler(router, usecase.NewTaskUsecase(pgsql.NewTaskRepository(db)))

	log.Println(fmt.Sprintf("server running on port %s", os.Getenv("PORT")))

	app.serve(app.loggerMiddleware(router))
}
