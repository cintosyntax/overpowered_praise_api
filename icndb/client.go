package icndb

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type response struct {
	Type  string        `json:"type,omitempty"`
	Value responseValue `json:"value,omitempty"`
}

type responseValue struct {
	ID         int      `json:"id,omitempty"`
	Joke       string   `json:"joke,omitempty"`
	Categories []string `json:"categories,omitempty"`
}

const (
	baseURL = "https://api.icndb.com"
)

// GetRandomJoke fetches a joke from the ICNDB API (http://www.icndb.com/api/) and
// and returns the resultant string w/o escaped characters.
func GetRandomJoke() string {
	// Defer replacing 'Chuck Norris' to after the call has been made.
	url := baseURL + "/jokes/random?escape=javascript&exclude=%5Bexplicit%5D"

	resp, _ := http.Get(url)
	body, _ := ioutil.ReadAll(resp.Body)

	var r response
	json.Unmarshal(body, &r)

	joke := r.Value.Joke

	return joke
}
