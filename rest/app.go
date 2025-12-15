package rest

import (
	"cloudego/storage"
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	r *mux.Router
}

func New(r *mux.Router) *App {
	return &App{r: r}
}

func (a *App) Start() {
	const op = "app.Start"

	a.r.HandleFunc("/v1/{key}", a.KeyValuePutHandler).Methods("PUT")
	a.r.HandleFunc("/v1/{key}", a.KeyValueGetHandler).Methods("GET")
	a.r.HandleFunc("/v1/{key}", a.KeyValueDeleteHandler).Methods("DELETE")

	log.Printf("%s, server is started at :8080", op)

	if err := http.ListenAndServe(":8080", a.r); err != nil {
		log.Fatalf("%s, something went wrong", op)
	}
}

func (a *App) KeyValuePutHandler(w http.ResponseWriter, r *http.Request) {
	const op = "app.KeyValuePutHandler"
	vars := mux.Vars(r)
	key := vars["key"]

	value, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Printf("%s, body request err", op)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = storage.Put(key, string(value))
	if err != nil {
		log.Printf("%s, storage err", op)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (a *App) KeyValueGetHandler(w http.ResponseWriter, r *http.Request) {
	const op = "app.KeyValueGetHandler"
	vars := mux.Vars(r)
	key := vars["key"]

	value, err := storage.Get(key)
	if errors.Is(err, storage.ErrorNoSuchKey) {
		log.Printf("%s: %w", op, err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if err != nil {
		log.Printf("%s: %w", op, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte(value))
}

func (a *App) KeyValueDeleteHandler(w http.ResponseWriter, r *http.Request) {
	const op = "app.KeyValueDeleteHandler"
	vars := mux.Vars(r)
	key := vars["key"]

	storage.Delete(key)
	w.WriteHeader(http.StatusAccepted)
}
