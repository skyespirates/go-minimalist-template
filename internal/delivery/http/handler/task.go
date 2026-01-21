package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/skyespirates/go-minimalist-template/internal/usecase"
)

type taskHandler struct {
	uc usecase.TaskUsecase
}

func NewTaskHandler(uc usecase.TaskUsecase) *taskHandler {
	return &taskHandler{uc}
}

func (th *taskHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	tasks, err := th.uc.GetAll(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := map[string]any{
		"data": tasks,
	}

	json.NewEncoder(w).Encode(resp)
}

func (th *taskHandler) GetById(w http.ResponseWriter, r *http.Request) {
	id := httprouter.ParamsFromContext(r.Context()).ByName("id")

	task, err := th.uc.GetById(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := map[string]any{
		"data": task,
	}

	json.NewEncoder(w).Encode(resp)
}

func (th *taskHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Title string `json:"title"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}

	task, err := th.uc.Create(r.Context(), req.Title)
	if err != nil {
		log.Printf("error: %s\n", err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := make(map[string]any)
	response["message"] = "task created successfully"
	response["task"] = task
	json.NewEncoder(w).Encode(response)
}

func (th *taskHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := httprouter.ParamsFromContext(r.Context()).ByName("id")

	todoId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	deletedId, err := th.uc.Delete(r.Context(), todoId)
	if err != nil {
		log.Printf("error: %s\n", err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	response := make(map[string]any)
	response["message"] = "todo deleted successfully"
	response["id"] = deletedId
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
