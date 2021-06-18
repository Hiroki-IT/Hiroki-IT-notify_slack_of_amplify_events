package slack

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/domain/entity/eventbridge"

	aws_amplify "github.com/aws/aws-sdk-go/service/amplify"
)

/**
 * Slackメッセージを構成します．
 */
type SlackMessage struct {
	Channel     string       `json:"channel"`
	Text        string       `json:"text"`
	Attachments []Attachment `json:"attachments"`
}

/**/
type Attachment struct {
	Color  string  `json:"color"`
	Blocks []Block `json:"blocks"`
}

/**/
type Block struct {
	Type     string    `json:"type"`
	Text     *Text     `json:"text,omitempty"`
	Elements []Element `json:"elements,omitempty"`
}

/**/
type Text struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

/**/
type Element struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

/**
 * コンストラクタ
 * Messageを作成します．
 */
func NewSlackMessage(eventDetail *eventbridge.EventDetail, branch *aws_amplify.Branch, jobStatusColor *JobStatusColor) *SlackMessage {

	// メッセージを構成します．
	return &SlackMessage{
		Channel: os.Getenv("SLACK_CHANNEL_ID"),
		Text:    "検証用dev環境",
		Attachments: []Attachment{
			Attachment{
				Color: jobStatusColor.PrintStatusColorCode(),
				Blocks: []Block{
					Block{
						Type: "section",
						Text: &Text{
							Type: "mrkdwn",
							Text: "*検証用dev環境*",
						},
					},
					Block{
						Type: "context",
						Elements: []Element{
							Element{
								Type: "mrkdwn",
								Text: fmt.Sprintf(
									"*結果*: %s",
									jobStatusColor.PrintStatusWord(),
								),
							},
						},
					},
					Block{
						Type: "context",
						Elements: []Element{
							Element{
								Type: "mrkdwn",
								Text: fmt.Sprintf(
									"*ブランチ名*: %s",
									eventDetail.BranchName,
								),
							},
						},
					},
					Block{
						Type: "context",
						Elements: []Element{
							Element{
								Type: "mrkdwn",
								Text: fmt.Sprintf(
									"*プルリクURL*: https://github.com/hiroki-it/notify-slack-of-amplify-events/compare/%s",
									eventDetail.BranchName,
								),
							},
						},
					},
					Block{
						Type: "context",
						Elements: []Element{
							Element{
								Type: "mrkdwn",
								Text: fmt.Sprintf(
									"*検証URL*: https://%s.%s.amplifyapp.com",
									aws.StringValue(branch.DisplayName),
									eventDetail.AppId,
								),
							},
						},
					},
					Block{
						Type: "context",
						Elements: []Element{
							Element{
								Type: "mrkdwn",
								Text: fmt.Sprintf(
									":amplify: <https://%s.console.aws.amazon.com/amplify/home?region=%s#/%s/%s/%s|*Amplifyコンソール画面はこちら*>",
									os.Getenv("AWS_AMPLIFY_REGION"),
									os.Getenv("AWS_AMPLIFY_REGION"),
									eventDetail.AppId,
									aws.StringValue(branch.DisplayName),
									eventDetail.JobId,
								),
							},
						},
					},
					Block{
						Type: "divider",
					},
				},
			},
		},
	}
}
