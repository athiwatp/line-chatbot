package line

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/agungdwiprasetyo/go-utils/debug"
	"github.com/agungdwiprasetyo/line-chatbot/src/linebot/domain"
)

const url = "https://api.line.me/v2/bot/message/push"

// SendMessage push message to channel
func SendMessage(message *domain.Message) error {
	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(true)
	encoder.Encode(message)

	req, _ := http.NewRequest("POST", url, buffer)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("LINE_CHANNEL_TOKEN")))

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	debug.Println(string(body))
	return nil
}
