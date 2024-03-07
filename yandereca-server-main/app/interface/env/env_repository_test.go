package env

import (
	"testing"
)

type ApiClientMock struct{}

func (a *ApiClientMock) ReadEnv(data string) (string, error) {
	return "http://localhost:8080", nil
}

func TestReadEnv(t *testing.T) {
	testInstance := EnvRepository{}
	testInstance.EnvHandler = &ApiClientMock{}
	expectText := "http://localhost:8080"
	result, _ := testInstance.ReadEnvValue("CORS_URL")
	if result.Value != expectText {
		t.Error("Result:", result, "Expected:", expectText)
	}
}
