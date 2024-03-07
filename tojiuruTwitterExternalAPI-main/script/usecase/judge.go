package usecase

import (
	"app/database"
	"errors"
)

type Check struct {
	MessageId string
	UserId    string
	Judge     bool
}

func CreateJudge(check *Check) error {
	judge, err := database.CheckJudge(check.MessageId, check.UserId)
	if err != nil {
		return err
	}
	// 投稿済みでない場合追加
	if judge == nil {
		if err = database.InsertJudge(
			&database.Judge{
				MessageId: check.MessageId,
				UserId:    check.UserId,
				Judge:     check.Judge,
			}); err != nil {
			return err
		} else {
			// insertして追加した
			return nil
		}
	}
	// 投稿していたらエラー
	err = errors.New("already voted")
	return err
}
