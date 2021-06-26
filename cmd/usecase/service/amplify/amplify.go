package amplify

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/amplify/amplifyiface"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/domain/detail"

	aws_amplify "github.com/aws/aws-sdk-go/service/amplify"
)

type AmplifyClientInterface interface {
	CreateGetBranchInput(*detail.Detail) *aws_amplify.GetBranchInput
	GetBranchFromAmplify(*detail.Detail) (*aws_amplify.GetBranchOutput, error)
}

type AmplifyClient struct {
	AmplifyClientInterface
	api amplifyiface.AmplifyAPI
}

// NewAmplifyClient コンストラクタ
func NewAmplifyClient(amplifyApi amplifyiface.AmplifyAPI) *AmplifyClient {

	return &AmplifyClient{
		api: amplifyApi,
	}
}

// GetBranchFromAmplify Amplifyからブランチ情報を取得します．
func (cl *AmplifyClient) GetBranchFromAmplify(detail *detail.Detail) (*aws_amplify.GetBranchOutput, error) {

	gbi := &aws_amplify.GetBranchInput{
		AppId:      aws.String(detail.AppId.Value()),
		BranchName: aws.String(detail.BranchName.Value()),
	}

	// ブランチ情報を構造体として取得します．
	gbo, err := cl.api.GetBranch(gbi)

	if err != nil {
		return nil, err
	}

	return gbo, nil
}
