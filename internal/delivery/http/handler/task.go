package handler

import (
	"encoding/json"
	"net/http"

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
