package app

import (
	"encoding/json"
	"net/http"

	"golang.org/x/net/context"
	"google.golang.org/appengine"
)

func init() {
	booksHandlers := &BooksHandlers{}

	http.HandleFunc("/books", booksHandlers.GetBooks)
}

type BooksHandlers struct {
	BooksService interface {
		FindAll(ctx context.Context) []Book
	}
}

func (h *BooksHandlers) GetBooks(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	books := h.BooksService.FindAll(ctx)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(books); err != nil {
		http.Error(w, "An error occured when encoding JSON", 500)
	}
}

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}
