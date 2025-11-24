package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	"github.com/skyespirates/go-minimalist-template/internal/delivery/http/handler"
	"github.com/skyespirates/go-minimalist-template/internal/infra/pgsql"
	"github.com/skyespirates/go-minimalist-template/internal/usecase"
)

type application struct{}

func main() {
	godotenv.Load()

	db, err := pgsql.InitDb()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	log.Println("database connection pool established")

	router := httprouter.New()

	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.Write([]byte("Hello, World!"))
	})

	ctx := context.Background()

	handler.NewTaskHandler(ctx, router, usecase.NewTaskUsecase(pgsql.NewTaskRepository(db)))

	app := &application{}

	log.Println(fmt.Sprintf("server running on port %s", os.Getenv("PORT")))

	app.serve(router)
}
