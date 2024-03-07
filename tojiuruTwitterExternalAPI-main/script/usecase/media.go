package usecase

import (
	"app/database"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"image"
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func CreateMediaFile(c *gin.Context) (string, error) {
	// imageのフォーマットをチェック
	format, err := checkImageFormat(c)
	if err != nil {
		return "", err
	}
	md5str, err := conversionMD5(c)
	if err != nil {
		return "", err
	}
	// ここからトランザクションにするか？
	// md5で検索し存在すればmediaIDを返す
	media, err := database.FindByMD5(md5str)
	if err != nil {
		return "", err
	}
	// Nosqlだった場合
	if media != nil {
		//存在していればファイルを検索
		isExist, err := isExistFile(media)
		if err != nil {
			return "", err
		}
		if isExist {
			//既に存在しているIDを返す
			return media.MediaID, nil
		} else {
			//DBに存在しているもののファイルに存在していない
			//消去する
			database.DeleteMediaFile(media.MediaID)
		}
	}
	//同ファイルが存在していなければ新たにDBに追加
	// トランザクションを張る
	// MediaIDをstring型としてdataに返す
	// interfaceで返されるのでstringで展開
	data, err := Transaction(c, database.Db, func(tx *sql.Tx) (interface{}, error) {
		mediaID, err := database.CreateMediaFileTx(tx, &database.MediaFile{
			Md5:    md5str,
			Format: format,
		})
		if err != nil {
			return nil, err
		}
		// createExaminationのタイミングでtweetIDを付与
		if err := database.CreateMediaTx(tx, &database.Media{
			MediaId: mediaID,
			Format:  format,
		}); err != nil {
			return nil, err
		}
		return mediaID, nil
	})
	if err != nil {
		return "", err
	}
	//ファイルを保存
	if err := createFile(c, data.(string), format); err != nil {
		return "", err
	}
	return data.(string), nil
}

func checkImageFormat(c *gin.Context) (string, error) {
	file, err := requestdata(c)
	if err != nil {
		return "", err
	}
	_, format, err := image.DecodeConfig(file)
	if err != nil {
		return "", err
	}
	return format, err
}

//受信したファイルのmd5を計算
func conversionMD5(c *gin.Context) (md5str string, err error) {
	md5hash := md5.New()
	file, err := requestdata(c)
	if err != nil {
		return
	}
	_, err = io.Copy(md5hash, file)
	if err != nil {
		return
	}
	md5str = hex.EncodeToString(md5hash.Sum(nil))
	return
}

//ファイルを保存
func createFile(c *gin.Context, id string, format string) error {
	saveFile, err := os.Create("../media/" + id + "." + format)
	if err != nil {
		return err
	}
	defer saveFile.Close()

	file, err := requestdata(c)
	if err != nil {
		return err
	}
	//保存コピー
	io.Copy(saveFile, file)
	return nil
}

func isExistFile(media *database.MediaFile) (isExist bool, err error) {
	file, err := os.Open("../media/" + media.MediaID + "." + media.Format)
	if err != nil {
		return
	}
	md5hash := md5.New()
	_, err = io.Copy(md5hash, file)
	if err != nil {
		return
	}
	md5str := hex.EncodeToString(md5hash.Sum(nil))
	//DBのmd5とimageディレクトリ内のファイルが等しいかをハッシュ値で確かめる
	if md5str == media.Md5 {
		isExist = true
		return
	} else {
		log.Println("other image")
		return
	}
}

// データを取得(流用するとバイナリがおかしくなるのでとりあえず一回一回とってくる)
func requestdata(c *gin.Context) (io.Reader, error) {
	uploadfile, _, err := c.Request.FormFile("file")
	if err != nil {
		return nil, err
	}
	defer uploadfile.Close()
	return uploadfile, nil
}
