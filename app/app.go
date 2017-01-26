package app

import (
	"encoding/json"
	"net/http"

	"golang.org/x/net/context"
	"google.golang.org/appengine"
)

func init() {
	ctxProvider := &ContextProvider{}
	booksService := &BooksService{}
	booksHandlers := &BooksHandlers{
		BooksService: booksService,
		Context: ctxProvider,
	}

	http.HandleFunc("/books", booksHandlers.GetBooks)
}

type BooksHandlers struct {
	BooksService interface {
		FindAll(ctx context.Context) []Book
	}
	Context interface {
		Get(r *http.Request) context.Context
	}
}

func (h *BooksHandlers) GetBooks(w http.ResponseWriter, r *http.Request) {
	ctx := h.Context.Get(r)

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

type ContextProvider struct{}

func (p *ContextProvider) Get(r *http.Request) context.Context {
	return appengine.NewContext(r)
}

type BooksService struct {}

func (b *BooksService) FindAll(ctx context.Context) []Book {
	return []Book{
		{Title: "The Lord of the Rings", Author: "J.J.R. Tolkien"},
		{Title: "Harry Potter", Author: "J.K. Rolling"},
	}
}
