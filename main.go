package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/cintosyntax/overpowered_praise_api/icndb"
)

const (
	serviceName         = "Overpowered Praise API"
	defaultResponseType = "in_channel"
	defaultPort         = "9000"
)

type slackResponseBody struct {
	ResponseType string `json:"response_type,omitempty"`
	Text         string `json:"text,omitempty"`
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		name := r.URL.Query().Get("name")
		if name == "" {
			name = r.URL.Query().Get("text")
		}

		praise := icndb.GetRandomJoke(name, "")
		apiResponse := &slackResponseBody{
			ResponseType: defaultResponseType,
			Text:         praise,
		}

		json.NewEncoder(w).Encode(apiResponse)
	})

	http.HandleFunc("/me", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		name := r.URL.Query().Get("user_name")
		name = strings.Title(name)

		praise := icndb.GetRandomJoke(name, "")
		apiResponse := &slackResponseBody{
			ResponseType: defaultResponseType,
			Text:         praise,
		}

		json.NewEncoder(w).Encode(apiResponse)
	})

	fmt.Println(serviceName + " serving on port :" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
