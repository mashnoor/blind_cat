package utility

import (
	"encoding/json"
	"fmt"
	"github.com/asmcos/requests"
	"github.com/mashnoor/blind_cat/settings"
)

func SendSlackMessage(serviceName string, down bool, errorCount int64) {

	type TextBlock struct {
		Type string `json:"type"`
		Text string `json:"text"`
	}
	type MainBlock struct {
		Type string    `json:"type"`
		Text TextBlock `json:"text"`
	}
	type ResponseBlock struct {
		Blocks []MainBlock `json:"blocks"`
	}

	downMsg := fmt.Sprintf("*%s Is Down* :crying_cat_face:\n Blind Cat cannot reach the service and marked it as down.\n*Total missing cats: %d*", serviceName, errorCount)
	upMsg := fmt.Sprintf("*%s Is Up!* :smile_cat:\n Blind Cat marked the service up after a hectic downtime.", serviceName)
	sendMsg := upMsg
	if down {
		sendMsg = downMsg
	}
	textBlock := TextBlock{
		Type: "mrkdwn",
		Text: sendMsg,
	}

	r := MainBlock{Type: "section", Text: textBlock}

	slackMsg := ResponseBlock{Blocks: []MainBlock{r}}

	jsonStr, err := json.Marshal(slackMsg)

	hookUrl := settings.SystemAppConfig.SlackUrl
	resp, err := requests.PostJson(hookUrl, string(jsonStr))
	fmt.Println(string(jsonStr))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.Text())
}
