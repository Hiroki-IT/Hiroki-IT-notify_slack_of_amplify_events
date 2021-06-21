module github.com/hiroki-it/notify-slack-of-amplify-events

go 1.15

require (
	github.com/aws/aws-lambda-go v1.23.0
	github.com/aws/aws-sdk-go v1.25.40
	github.com/stretchr/testify v1.7.0
	go.uber.org/zap v1.16.0
)

replace (
	github.com/hiroki-it/notify-slack-of-amplify-events/cmd/usecase/service/amplify => /cmd/usecase/service/amplify
	github.com/hiroki-it/notify-slack-of-amplify-events/cmd/usecase/service/notification => /cmd/usecase/service/notification
	github.com/hiroki-it/notify-slack-of-amplify-events/cmd/domain/entity/event_detail => /cmd/domain/entity/event_detail
	github.com/hiroki-it/notify-slack-of-amplify-events/cmd/domain/entity/file => /cmd/infrastructure/file
	github.com/hiroki-it/notify-slack-of-amplify-events/cmd/presentation/controllers => /cmd/presentation/controllers
	github.com/hiroki-it/notify-slack-of-amplify-events/mock/amplify => /mock/amplify
)
