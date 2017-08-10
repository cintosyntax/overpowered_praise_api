package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/cintosyntax/overpowered_praise_api/icndb"
)

const (
	serviceName         = "Overpowered Praise API"
	defaultResponseType = "in_channel"
)

type slackResponseBody struct {
	ResponseType string `json:"response_type,omitempty"`
	Text         string `json:"text,omitempty"`
}

func main() {
	port := os.Getenv("PORT")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")

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
