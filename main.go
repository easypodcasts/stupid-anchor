package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	token := os.Getenv("TOKEN")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		url_key, ok_url := r.URL.Query()["bringme"]
		token_key, ok_token := r.URL.Query()["token"]
		if ok_url && len(url_key[0]) > 1 && ok_token && len(token_key[0]) > 1 && token_key[0] == token {
			resp, err := http.Get(url_key[0])
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprint(w, "error")
				return
			}
			io.Copy(w, resp.Body)
		}
	})

	log.Println("listening on", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
