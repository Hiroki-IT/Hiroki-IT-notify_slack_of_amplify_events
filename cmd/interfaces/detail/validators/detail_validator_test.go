package validators

import (
	"encoding/json"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/domain/file/entities"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/domain/file/values"
	"testing"

	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/infrastructure/file/repositories"
	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {

	t.Helper()

	// テストケース
	cases := []struct {
		// テストケース名
		name string
		// テストデータ
		detail json.RawMessage
	}{
		{
			name:   "TestValidate_Valid_ReturnEmpty",
			detail: repositories.NewFileRepository(entities.NewFile(values.NewPath("./test_data/valid.json"))).ByteLoad(),
		},
		{
			name:   "TestValidate_InvalidRequired_ExceptionError",
			detail: repositories.NewFileRepository(entities.NewFile(values.NewPath("./test_data/invalid_required.json"))).ByteLoad(),
		},
	}

	// 反復処理で全てのテストケースを検証します．
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {

			t.Log(string(tt.detail))

			v := NewDetailValidator()

			err := json.Unmarshal(tt.detail, v)

			if err != nil {
				t.Fatal(err.Error())
			}

			errorMessage := v.Validate()

			t.Log(errorMessage)

			if errorMessage != nil {
				// 異常系テストのアサーション
				assert.Error(t, errorMessage)
			} else {
				// 正常系テストのアサーション
				assert.Nil(t, errorMessage)
			}
		})
	}
}
