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

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		keys, ok := r.URL.Query()["bringme"]
		if ok && len(keys[0]) > 1 {
			resp, err := http.Get(keys[0])
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
