package handler

import (
	"encoding/json"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/entities/amplify"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/entities/eventbridge"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/entities/slack"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/usecases/logger"
)

/**
 * Lambdaハンドラー関数
 */
func HandleRequest(event events.CloudWatchEvent) {

	eventDetail := eventbridge.NewEventDetail()

	// eventbridgeから転送されたJSONを構造体にマッピングします．
	err := json.Unmarshal([]byte(event.Detail), eventDetail)

	if err != nil {
		logger.ErrorLog(err)
	}

	amplifyApi, err := amplify.NewAmplifyAPI(os.Getenv("AWS_REGION"))

	if err != nil {
		logger.ErrorLog(err)
	}

	amplifyClient := amplify.NewAmplifyClient(amplifyApi)

	getBranchInput := amplifyClient.CreateGetBranchInput(eventDetail)

	getBranchOutput, err := amplifyClient.GetBranchFromAmplify(getBranchInput)

	if err != nil {
		logger.ErrorLog(err)
	}

	slackClient := slack.NewSlackClient()

	message := slackClient.BuildMessage(
		eventDetail,
		getBranchOutput.Branch,
	)

	err = slackClient.PostMessage(message)

	if err != nil {
		logger.ErrorLog(err)
	}

	log.Println("Exit")
}
