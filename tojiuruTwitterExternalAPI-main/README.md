# TwitterExternalAPI

## dockerでの動作検証方法
### コンテナを作るまで
1. build\_testブランチをローカル環境にcloneなりでコピー
1. scriptディレクトリ直下に.envファイルを配置
1. "docker image build -t (任意のイメージ名):latest (Dockerfileのあるディレクトリ)"を実行
  1. "docker images"でイメージが作られたかを確認可
1. "docker run --name (任意のコンテナ名) -it -d -p 8080:8080 (利用するイメージ名 or ID)"
  1. "docker ps"で動作中のコンテナを確認可
### サーバサイドプログラムを動かすまで
いずれかの方法でプログラムを実行
- "docker exec -i (コンテナ名 or ID) sh -c "cd script && go run ."
- "docker exec -it (コンテナ名 or ID) bash"でコンテナに入る
  - main.goを実行
### コンテナを止める
- "docker stop (コンテナ名 or ID)"

### imageファイルアップロード用curlコマンド
#### base64形式
- curl -X POST -H "Content-Type: application/json" -d '{"key":"test.png", “(base64でエンコードした文字列)” localhost/tojiuru/image64
#### multipart/form-data形式
- curl -X POST -F file1=@/var/tmp/sample.jpg localhost/tojiuru/image-multiport
=======
## APIエンドポイント
### Examinaton
- POST   /examination/create  
```
{
  "message":string,
  "deadline":string,
  "people":int,
  "media_id",string 含んでなくても良い
}
```
- GET    /examination/data  
[]Examination  
```
type Examination {
  "message_id":string,
  "message":string,
  "people":int,
  "good_num":int,
  "bad_num":int,
  "created_at":string,
  "deadline":string,
  "user_id":string,
  "username":string,
  "state":int
}
```
- DELETE /examination/delete
### Judge
- POST   /judge/create
```
{
  "judge":bool,
  "tweet_id":string
}
```
### Media
- POST   /media/upload
```
{
  "media_id":string
}
```
### Reply "まだ調整出来てない"
- POST   /reply/create
```
{
  "tweet_id":string,
  "reply_text":string
}
```
- GET    /reply/data  
[]Reply  
```
type Reply {
  "tweet_id":string,
  "reply_id":string,
  "reply_text":string,
  "to_uid":string,
  "from_uid":string,
  "username":string,
  "created_at":string
}
```
### Tojiuru
- POST   /tojiuru/authenticate "ログイン"  
```
{
  "name":string,
  "password":string
}
```
- POST   /tojiuru/signup_account "アカウント作成"  
```
{
  "name":"",
  "password":""
}
```
- GET    /tojiuru/signout_account
- GET    /tojiuru/data

