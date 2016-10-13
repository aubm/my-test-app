package app

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"golang.org/x/net/context"
)

func TestGetBooks(t *testing.T) {
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/", nil)

	mockBooksService := &MockBooksService{}
	mockContextProvider := &MockContextProvider{}
	h := BooksHandlers{BooksService: mockBooksService, Context: mockContextProvider}

	h.GetBooks(w, r)

	body := w.Body.String()
	expected := `[{"title":"The Lord of the Rings","author":"J.J.R. Tolkien"},{"title":"Harry Potter","author":"J.K. Rolling"}]
`

	if body != expected {
		t.Errorf("Expected %t, got %t", expected, body)
	}
}

type MockBooksService struct{}

func (m *MockBooksService) FindAll(ctx context.Context) []Book {
	return []Book{
		{Title: "The Lord of the Rings", Author: "J.J.R. Tolkien"},
		{Title: "Harry Potter", Author: "J.K. Rolling"},
	}
}

type MockContextProvider struct{}

func (m *MockContextProvider) Get(r *http.Request) context.Context {
	return context.Background()
}
