package textmining

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const (
	defaultError = "Mohon coba lagi"
)

// ProcessText for mining input text
func ProcessText(text string) string {
	req, _ := http.NewRequest("GET", os.Getenv("CHATBOT_AI_HOST"), nil)
	uri := req.URL.Query()
	uri.Add("input", text)
	req.URL.RawQuery = uri.Encode()

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return defaultError
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return defaultError
	}

	var output struct {
		Output string `json:"output"`
	}
	json.Unmarshal(body, &output)

	return strings.TrimLeftFunc(output.Output, func(r rune) bool {
		return r == '-' || r == ' '
	})
}
