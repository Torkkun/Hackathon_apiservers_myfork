package handler

import (
  "github.com/gin-gonic/gin"
  "encoding/base64"
  "os"
  "io"
  "fmt"
  "os/exec"
  //"app/database"
  //"encoding/json"
)

type FileData struct {
  Filename string `json:"key"`
  Fileinfo string `json:"attachement"`
}

type BotData_typetalk struct {
  Tool  string `json:"tool"`
  Token string `json:"token"`
  Send_To   string `json:"send_to"`
}

func savefile_base64(c *gin.Context) {
  var data FileData
  if err := c.ShouldBindJSON(&data); err != nil {
    c.JSON(500, err)
    return
  }
  info, _ := base64.StdEncoding.DecodeString(data.Fileinfo)
  save_info := "/go/src/script/image/"
  save_info += data.Filename
  file, _ := os.Create(save_info)
  defer file.Close()

  file.Write(info)
}


func savefile_multiport(c *gin.Context) {
    image, header, _ := c.Request.FormFile("image")
    saveFile, _ := os.Create("./image/" + header.Filename)
    defer saveFile.Close()
    io.Copy(saveFile, image)

    /*
    jsonStr := c.Request.FormValue("formData")
    var p Post
    json.Unmarshal([]byte(jsonStr), &p)
    fmt.Println(p.Title, p.Author, p.Description)

    c.JSON(201, p)
    */
}

func savefile_botinfo(c *gin.Context) {
  var data BotData_typetalk
  if err := c.ShouldBindJSON(&data); err != nil {
    c.JSON(500, err)
    return
  }
  //info := fmt.Sprintf("TOKEN=%s\nURL=%s\n", data.Token, data.Url)
  info := fmt.Sprintf("{\n\"TOKEN\":\"%s\",\n\"URL\":\"%s\"\n}\n", data.Token, data.Send_To)
  save_info := "/go/src/script/image/"
  save_info += "access.json"
  file, _ := os.Create(save_info)
  defer file.Close()

  for _, info := range info {
    _, err := file.WriteString(string(info))
    if err != nil {
      c.JSON(500, err)
      return
    }
  }
  ps,err := exec.Command("sh","-c","/go/src/script/send_to_typetalk.sh 連携が完了しました").CombinedOutput()
  if err != nil {
    c.JSON(500, err)
    fmt.Println(string(ps),err)
    return
  }
}

/*
func savefile_botinfo(c *gin.Context) {
  var data BotData_typetalk
  err := json.NewDecoder(c.Request.Body).Decode(&data)
  if err != nil {
    fmt.Printf("miss decode json :%v", err)
    c.JSON(500, err)
    return
  }

 u := database.Bot{
   Tool: data.Tool,
   Token: data.Token,
   Send_To:   data.Send_To,
 }
 err = u.InsertBot()

 if err != nil {
		fmt.Println(err)
		c.JSON(500, "This name is not available")
		return
	}
	c.JSON(200, "ok")

  notification("連携が完了しました!")

  if err != nil {
    c.JSON(500, err)
    return
  }
}
*/
