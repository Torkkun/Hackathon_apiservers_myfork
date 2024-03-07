package usecase

import "testing"

const (
	TestMessageID      = "0000"
	TestOwner          = "1001"
	TestUser1          = "0001"
	TestUser2          = "0002"
	TestReplyID        = "1111"
	TestReplyMessageID = "AAAA"
)

func TestCreateRepU(t *testing.T) {
	//ホントは値もちゃんと確かめるべきだが
	t.Run("Ownerではないユーザーの場合のテスト", func(t *testing.T) {
		isOther, err := CheckUserIdFromExamination(TestMessageID, TestUser1)
		if err != nil {
			t.Error(err.Error())
		}
		t.Log(isOther)
		if isOther {
			if err := CreateReplyUser(TestMessageID, TestUser1); err != nil {
				t.Error(err.Error())
			}
		}
	})
	t.Run("Ownerだった場合のテスト実行されない", func(t *testing.T) {
		isOther, err := CheckUserIdFromExamination(TestMessageID, TestOwner)
		if err != nil {
			t.Error(err.Error())
		}
		t.Log(isOther)
		if isOther {
			t.Error("Owner create self reply")
			if err := CreateReplyUser(TestMessageID, TestOwner); err != nil {
				t.Error(err.Error())
			}
		}
	})
}

func TestGetReplyU(t *testing.T) {
	list, err := GetReplyUser(TestMessageID)
	if err != nil {
		t.Error(err.Error())
	}
	for _, v := range *list {
		t.Log(v.ReplyID, v.FromUserID, v.FromUserName, v.CreatedAt)
	}
}

func TestCreateReplyM(t *testing.T) {
	t.Run("User2疎通できるユーザー", func(t *testing.T) {
		if err := CheckUserIdFromReplyUser(TestUser2, TestReplyID); err != nil {
			t.Error(err.Error())
		}
		CreateReplyMessage(TestReplyID, TestUser2, "ユーザー2 メッセージ")
	})
	t.Run("Owner疎通できるオーナー", func(t *testing.T) {
		if err := CheckUserIdFromReplyUser(TestOwner, TestReplyID); err != nil {
			t.Error(err.Error())
		}
		CreateReplyMessage(TestReplyID, TestOwner, "オーナー　メッセージ")
	})
	t.Run("User1疎通できないユーザー", func(t *testing.T) {
		if err := CheckUserIdFromReplyUser(TestUser1, TestReplyID); err != nil {
			t.Log(err.Error())
		}
		CreateReplyMessage(TestReplyID, TestUser1, "ユーザー1 メッセージ")
	})
}

func TestGetReplyM(t *testing.T) {
	data, err := GetReplyMessage(TestReplyID)
	if err != nil {
		t.Error(err.Error())
	}
	for _, v := range *data {
		t.Log(v.ReplyID, v.ReplyMessageID, v.Message, v.UserID, v.UserName, v.CreatedAt)
	}
}
