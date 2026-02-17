package main

import (
	"embed"
	"encoding/json"
	"io/fs"
	"log"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/skyespirates/go-minimalist-template/internal/delivery/http/handler"
	"github.com/skyespirates/go-minimalist-template/internal/infra/pgsql"
	"github.com/skyespirates/go-minimalist-template/internal/usecase"
	"github.com/skyespirates/go-minimalist-template/internal/utils"
)

//go:embed dist
var embeddedFiles embed.FS

func (app *application) routes() http.Handler {
	router := httprouter.New()

	distFS, err := fs.Sub(embeddedFiles, "dist")
	if err != nil {
		log.Fatal(err)
	}

	fileServer := http.FileServer(http.FS(distFS))

	taskHandler := handler.NewTaskHandler(usecase.NewTaskUsecase(pgsql.NewTaskRepository(app.db)))
	userHandler := handler.NewUserHandler(usecase.NewUserUsecase(pgsql.NewUserRepository(app.db)))

	router.HandlerFunc(http.MethodGet, "/healthcheck", healthcheck)

	router.HandlerFunc(http.MethodPost, "/v1/auth/register", userHandler.Register)
	router.HandlerFunc(http.MethodPost, "/v1/auth/login", userHandler.Login)

	router.HandlerFunc(http.MethodGet, "/v1/tasks", app.authenticate(taskHandler.GetAll))
	router.HandlerFunc(http.MethodGet, "/v1/tasks/:id", taskHandler.GetById)
	router.HandlerFunc(http.MethodPost, "/v1/tasks", app.authenticate(taskHandler.Create))
	router.HandlerFunc(http.MethodPut, "/v1/tasks/:id", app.authenticate(taskHandler.Update))
	router.HandlerFunc(http.MethodDelete, "/v1/tasks/:id", taskHandler.Delete)

	router.HandlerFunc(http.MethodPost, "/api/generate-key", func(w http.ResponseWriter, r *http.Request) {
		res := make(map[string]string)

		res["key"] = utils.GenerateKey()

		json.NewEncoder(w).Encode(res)
	})

	router.HandlerFunc(http.MethodPost, "/api/encrypt", func(w http.ResponseWriter, r *http.Request) {
		var input struct {
			Key  string `json:"key"`
			Text string `json:"text"`
		}

		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}
		if input.Key == "" || input.Text == "" {
			http.Error(w, "bad request, key and text are required", http.StatusBadRequest)
			return
		}

		result := utils.Encrypt(input.Key, input.Text)

		res := make(map[string]string)
		res["encrypted"] = result
		err = json.NewEncoder(w).Encode(res)
		if err != nil {
			http.Error(w, "error on json encoder", http.StatusInternalServerError)
		}
	})

	router.HandlerFunc(http.MethodPost, "/api/decrypt", func(w http.ResponseWriter, r *http.Request) {
		var input struct {
			Key  string `json:"key"`
			Text string `json:"text"`
		}

		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}
		if input.Key == "" || input.Text == "" {
			http.Error(w, "bad request, key and encrypted text are required", http.StatusBadRequest)
			return
		}

		decoded := utils.Decrypt(input.Key, input.Text)
		res := make(map[string]string)
		res["decrypted"] = decoded

		json.NewEncoder(w).Encode(res)

	})

	// router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	// Try to serve static file first
	// 	_, err := distFS.Open(strings.TrimPrefix(r.URL.Path, "/"))
	// 	if err == nil {
	// 		fileServer.ServeHTTP(w, r)
	// 		return
	// 	}

	// 	// If file doesn't exist → serve index.html (SPA fallback)
	// 	r.URL.Path = "/"
	// 	fileServer.ServeHTTP(w, r)
	// })

	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// try to serve static file
		_, err := distFS.Open(strings.TrimPrefix(r.URL.Path, "/"))
		if err == nil {
			fileServer.ServeHTTP(w, r)
			return
		}

		// fallback to index.html
		r.URL.Path = "/index.html"
		fileServer.ServeHTTP(w, r)
	})

	return app.loggerMiddleware(router)
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("All iz well 👌"))
}
