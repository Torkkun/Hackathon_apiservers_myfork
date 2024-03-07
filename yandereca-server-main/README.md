# yandere-server

## テストタスク登録

curl -i -H "Origin:http://localhost" localhost:8080/task/test/create<br>
curl -i -H "Origin:http://localhost" localhost:8080/task/test/read<br>
curl -i -H "Origin:http://localhost" localhost:8080/task/progress<br>
curl -X POST -H "Content-Type: application/json" -H "Origin:http://localhost" -d '{"task_id":"example-task-id-1","task":"example_task","desc":"hogehoge"}' localhost:8080/task/create

## テストユーザー登録
curl -i -H "Origin:http://localhost" localhost:8080/user/test/create<br>
curl -i -H "Origin:http://localhost" localhost:8080/user/test/read<br>
curl -X DELETE -H "Origin:http://localhost" localhost:8080/user/delete<br>
curl -X POST -H "Content-Type: application/json" -H "Origin:http://localhost" -d '{"id":"example-id","name":"example-name","email":"example-email","token":"example-token","refresh_token":"example-refresh-token","google_uid":"example-google-uid"}' localhost:8080/user/create

## GoogleTodoユーザー認証
curl -i -H "Origin:http://localhost" localhost:8080/googletask/request<br>
curl -X POST -H "Content-Type: application/json" -H "Origin:http://localhost" -d '{"code":"**コード**"}' localhost:8080/googletask/auth<br>
curl -i -H "Origin:http://localhost" localhost:8080/googletask/progress?uid=**ユーザーid**<br>

## Docker滅びの呪文（SQL関連の更新が来たときはこれを実行することを推奨）

docker-compose down --rmi all --volumes --remove-orphans

## テスト方法

run.shを実行してください

- 登録用API：http://localhost:8080/task/create
- 更新用API：http://localhost:8080/task/update
- 進捗率計算用API：http://localhost:8080/task/progress
