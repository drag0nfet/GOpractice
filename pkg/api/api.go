package api

import "net/http"

func FillEndpoints() {
	http.HandleFunc("/api/hello", HelloHandler)
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello World"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
