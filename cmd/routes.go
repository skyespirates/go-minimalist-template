package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/skyespirates/go-minimalist-template/internal/delivery/http/handler"
	"github.com/skyespirates/go-minimalist-template/internal/infra/pgsql"
	"github.com/skyespirates/go-minimalist-template/internal/usecase"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	taskHandler := handler.NewTaskHandler(usecase.NewTaskUsecase(pgsql.NewTaskRepository(app.db)))

	router.HandlerFunc(http.MethodGet, "/", index)
	router.HandlerFunc(http.MethodGet, "/healthcheck", healthcheck)

	router.HandlerFunc(http.MethodGet, "/v1/tasks", taskHandler.GetAll)
	router.HandlerFunc(http.MethodGet, "/v1/tasks/:id", taskHandler.GetById)

	return app.loggerMiddleware(router)
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, Skyes! 😎"))
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("All iz well 👌"))
}
