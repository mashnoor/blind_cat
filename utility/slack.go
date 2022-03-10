package utility

import (
	"encoding/json"
	"fmt"
	"github.com/asmcos/requests"
)

func SendSlackMessage(serviceName string, down bool) {

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

	textBlock := TextBlock{
		Type: "mrkdwn",
		Text: fmt.Sprintf("*%s* Seems Down* :crying_cat_face:\n Blind Cat cannot reach the service and marked as down :sob:. Send your search team to find the cat.", serviceName),
	}

	r := MainBlock{Type: "section", Text: textBlock}

	slackMsg := ResponseBlock{Blocks: []MainBlock{r}}

	jsonStr, err := json.Marshal(slackMsg)

	hookUrl := "https://hooks.slack.com/services/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"
	resp, err := requests.PostJson(hookUrl, string(jsonStr))
	fmt.Println(string(jsonStr))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.Text())
}
