package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Article struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"`
}

func ArticlesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("endpoint hit: /articles")

	articles := []Article{
		{Title: "Hello", Description: "Article description", Content: "Article content"},
		{Title: "Title", Description: "Article description", Content: "Article content"},
	}
	json.NewEncoder(w).Encode(articles)
}
