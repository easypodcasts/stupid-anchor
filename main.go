package main

import (
	"fmt"
	"io/ioutil"
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
			//We Read the response body on the line below.
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprint(w, "error")
				return
			}
			//Convert the body to type string
			sb := string(body)
			fmt.Fprint(w, sb)
		}
	})

	log.Println("listening on", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
