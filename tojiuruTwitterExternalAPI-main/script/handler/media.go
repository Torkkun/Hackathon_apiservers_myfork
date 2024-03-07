package handler

import (
	"app/domain"
	"app/usecase"
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

// POST /media/upload
func mediaUpload(c *gin.Context) {
	mediaId, err := usecase.CreateMediaFile(c)
	if err != nil {
		log.Println(err)
	}
	//新たに作成したIDを返す
	c.JSON(200, domain.Media{Media_id: mediaId})
}

// テスト用
// ioutil.ReadAllで全取得したものを変更するが実行時間等は変わらない気がする
func testmedia(c *gin.Context) {
	uploadfile, header, err := c.Request.FormFile("file")
	if err != nil {
		log.Println(err)
		return
	}
	content, err := ioutil.ReadAll(uploadfile)
	if err != nil {
		log.Println(err)
		return
	}
	reader := bytes.NewReader(content)
	// format
	_, format, err := image.DecodeConfig(reader)
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("Image Format :%s", format)

	// file 保存
	saveFile, err := os.Create("../media/" + header.Filename)
	if err != nil {
		log.Println(err)
		return
	}
	defer saveFile.Close()
	reader = bytes.NewReader(content)
	_, err = io.Copy(saveFile, reader)
	if err != nil {
		log.Println(err)
		return
	}
	reader = bytes.NewReader(content)
	// md5計算
	md5hash := md5.New()
	_, err = io.Copy(md5hash, reader)
	if err != nil {
		log.Println(err)
		return
	}
	md5str := hex.EncodeToString(md5hash.Sum(nil))
	log.Printf("MD5 string: %s", md5str)
}
