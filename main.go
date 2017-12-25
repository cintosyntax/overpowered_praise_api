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

	http.HandleFunc("/me", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		name := r.URL.Query().Get("user_name")
		name = strings.Title(name)

		if name == "" {
			// Return 400 if no name was given this endpoint expects a name.
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400 - Expected name to be given"))
			return
		}

		cnJoke := icndb.GetRandomJoke()
		praise := ReplaceCN(cnJoke, name)

		// Convert Chuck Norris Joke to praise by replacing all instnaces of his
		// name if it is given.

		apiResponse := &slackResponseBody{
			ResponseType: defaultResponseType,
			Text:         praise,
		}

		json.NewEncoder(w).Encode(apiResponse)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		name := r.URL.Query().Get("name")
		if name == "" {
			name = r.URL.Query().Get("text")
		}

		cnJoke := icndb.GetRandomJoke()
		praise := ReplaceCN(cnJoke, name)

		apiResponse := &slackResponseBody{
			ResponseType: defaultResponseType,
			Text:         praise,
		}

		json.NewEncoder(w).Encode(apiResponse)
	})

	fmt.Println(serviceName + " serving on port :" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// ReplaceCN takes a string with the name Chuck Norris
// and replace all instances with the name provided.
// Note: Some jokes don't make sense even if you do this, for example:
// How much wood would a woodchuck chuck if a woodchuck could Chuck Norris? All of it.
func ReplaceCN(joke string, name string) string {
	joke = strings.Replace(joke, "Chuck Norris ", name+" ", -1)
	return joke
}
