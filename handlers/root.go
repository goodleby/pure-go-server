package handlers

import (
	"fmt"
	"net/http"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("endpoint hit: root")

	fmt.Fprintf(w, "Welcome to the root!")
}
