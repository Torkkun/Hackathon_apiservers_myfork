package infra

import (
	"testing"
)

func TestReadEnv(t *testing.T) {
	testInstance := NewEnvHandler()
	//ここは環境変数に指定した値に適宜変更
	expectText := "http://localhost:8080"
	result, _ := testInstance.ReadEnv("CORS_URL")
	if result != expectText {
		t.Error("Result:", result, "Expected:", expectText)
	}
}
