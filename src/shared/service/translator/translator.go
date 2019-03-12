package translator

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// Translate utils for translate language
func Translate(from, to, text string) (result string) {
	value := url.Values{}
	value.Set("key", os.Getenv("TRANSLATOR_KEY"))
	value.Set("lang", from+"-"+to)
	value.Add("text", text)

	resp, err := http.PostForm(os.Getenv("TRANSLATOR_HOST"), value)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	var response struct {
		Code int      `json:"code"`
		Lang string   `json:"lang"`
		Text []string `json:"text"`
	}

	json.Unmarshal(b, &response)
	return strings.Join(response.Text, " ")
}
