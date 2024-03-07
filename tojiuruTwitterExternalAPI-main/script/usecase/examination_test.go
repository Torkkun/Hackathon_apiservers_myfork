package usecase

import (
	"app/domain"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

const (
	TestMessage  = "Test Message"
	TestDeadline = "Test Deadline"
	TestUserID   = "Test UserID"
	TestPass     = "testpass00!"
)

func TestCreateE(t *testing.T) {
	t.Run("メディアがない場合のテスト", func(t *testing.T) {
		gctx, _ := gin.CreateTestContext(httptest.NewRecorder())
		if err := CreateExamination(gctx, TestUserID, domain.PostExamination{
			Message:  TestMessage,
			Deadline: TestDeadline,
			People:   5,
			MediaID:  nil,
		}); err != nil {
			t.Error(err.Error())
		}
	})
}

func TestGetE(t *testing.T) {
	gete, err := GetExamination(TestUserID)
	if err != nil {
		t.Error(err.Error())
	}
	for _, v := range gete.Examinations {
		t.Log(v.UserId, v.Message_id, v.Message)
	}
}
