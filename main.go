package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	http.HandleFunc("/books", booksHandler)
	http.ListenAndServe(":8080", nil)
}

func booksHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"meta": map[string]string{
			"timestamp":   "2023-10-03T10:30:00",
			"api_version": "1.5.2",
			"status":     "success",
		},
		"data": map[string]interface{}{
			"books": []map[string]interface{}{
				{
					"id": "B123",
					"title": "Mystery of the Missing Key",
					"authors": []map[string]string{
						{"id": "A456", "name": "John Doe"},
						{"id": "A457", "name": "Jane Doe"},
					},
					"reviews": []map[string]interface{}{
						{"id": "R789", "rating": 4, "text": "Good book!"},
						{"id": "R790", "rating": 5, "text": "Excellent!"},
					},
					"borrowers": []map[string]interface{}{
						{"id": "BR001", "name": "Alice", "due_date": "2023-11-03"},
						{"id": "BR002", "name": "Bob", "due_date": "2023-11-10"},
					},
				},
				{
					"id": "B124",
					"title": "The AI Chronicles",
					"authors": []map[string]string{
						{"id": "A458", "name": "Ella Smith"},
					},
					"reviews":  []map[string]string{},
					"borrowers": []map[string]interface{}{
						{"id": "BR003", "name": "Charlie", "due_date": "2023-11-15"},
					},
				},
			},
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
