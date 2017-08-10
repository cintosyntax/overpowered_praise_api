package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/cintosyntax/overpowered_praise_api/icndb"
)

const (
	defaultPort         = "9000"
	serviceName         = "Overpowered Praise API"
	defaultResponseType = "in_channel"
)

type slackResponseBody struct {
	ResponseType string `json:"response_type,omitempty"`
	Text         string `json:"text,omitempty"`
}

func main() {
	// joke := icndb.GetRandomJoke("Edwin", "Mak")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")

		praise := icndb.GetRandomJoke(name, "")
		apiResponse := &slackResponseBody{
			ResponseType: defaultResponseType,
			Text:         praise,
		}

		json.NewEncoder(w).Encode(apiResponse)
	})

	fmt.Println(serviceName + " serving on port :" + defaultPort)
	log.Fatal(http.ListenAndServe(":"+defaultPort, nil))
}
