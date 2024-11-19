package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/all-files", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			files := []string{"file1.txt", "file2.jpg", "file3.pdf"}

			w.Header().Set("Content-Type", "application/json")

			json.NewEncoder(w).Encode(files)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	fmt.Printf("running on http://localhost:3888/all-files")
	http.ListenAndServe(":3888", nil)
}
