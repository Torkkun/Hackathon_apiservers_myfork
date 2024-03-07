package usecase

import (
	"app/database"
	"app/domain"
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateExamination(c *gin.Context, userID string, examination domain.PostExamination) error {
	messageId, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	// トランザクション張る
	_, err = Transaction(c, database.Db, func(tx *sql.Tx) (interface{}, error) {
		// mediaIDが0じゃなかった場合
		if examination.MediaID != nil {
			if err := database.UpdateMediaByMessageIDTx(tx, messageId.String()); err != nil {
				return nil, err
			}
		}
		if err := database.CreateExaminationTx(
			tx,
			&database.Examination{
				MessageId: messageId.String(),
				Message:   examination.Message,
				People:    examination.People,
				UserId:    userID,
				Deadline:  examination.Deadline,
			}); err != nil {
			return nil, err
		}
		return nil, nil
	})
	return err
}

func GetExamination(userID string) (*domain.ResponseExaminations, error) {
	// 全データをDBから取得
	examinations, err := database.SelectExamination()
	if err != nil {
		return nil, fmt.Errorf("select examination error: %s", err.Error())
	}
	// Media用のリスト
	//var messageIDs []string

	// response用のリスト
	var resExami domain.ResponseExaminations
	for _, v := range *examinations {
		// Media用
		//messageIDs = append(messageIDs, v.MessageId)

		// 審査tweetごとのトジウル内良いね数悪いね数を計算
		judgemap, err := database.CountJudge(v.MessageId)
		if err != nil {
			return nil, fmt.Errorf("count judge error: %s", err.Error())
		}
		// 審査tweetごとのstate計算のための値を取得
		// GETしてきたユーザーの状態を返す
		judgeData, err := database.CheckJudge(v.MessageId, userID)
		if err != nil {
			return nil, fmt.Errorf("check judge error: %s", err.Error())
		}
		// レスポンス用state値を決める
		var state int
		if judgeData == nil {
			state = -1
		} else {
			if judgeData.Judge {
				// 良いね
				state = 1
			} else {
				// 危ないね
				state = 0
			}
		}
		// 構造体に格納
		resExami.Examinations = append(resExami.Examinations, domain.Examination{
			Message_id: v.MessageId,
			Message:    v.Message,
			People:     v.People,
			Good_num:   judgemap[true],
			Bad_num:    judgemap[false],
			CreatedAt:  v.CreatedAt,
			Deadline:   v.Deadline,
			UserId:     v.UserId,
			Username:   v.UserName,
			State:      state,
		})
	}
	// Media用　辛いので無視
	/* if messageIDs != nil {
		var mediares domain.MediaResponse
		for _, id := range messageIDs {
			// media url get
			media, err := database.SelectMediaFindByMessageId(id)
			if err != nil {
				return nil, err
			}
			mediares.MessageID = id
			// レスポンス用のurlリストを作成
			for _, m := range *media {
				// mediaIDとフォーマットを組み合わせてURLを作成
				mediares.MediaURL = append(mediares.MediaURL, m.MediaId+m.Format)
			}
			resExami.Media = append(resExami.Media, mediares)
		}
	} */
	return &resExami, nil
}

// OwnerまたはERRORの場合はfalse,それ以外ならtrue。しかしエラーの場合はfalseとerrorが入るのでerrorが先に処理される
func CheckUserIdFromExamination(messageID, userID string) (bool, error) {
	idfromdb := database.FindUserIDFromExaminationByMessageID(messageID)
	if idfromdb == "" {
		return false, fmt.Errorf("no userid from examination tabele")
	}
	if idfromdb == userID {
		return false, nil
	}
	return true, nil
}
