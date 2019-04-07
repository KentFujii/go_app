# go_app

golang application programs

## getting started

開発環境にk8sを、CI/CDにskaffoldを使用しています

```
skaffold dev
```

## 構成

下記それぞれのディレクトリがそれぞれ一つのPod/Deployment/Serviceとしてk8sに収まるようになっている

### primary

簡単なWebアプリケーション

```
curl localhost:50001
```

### handling

net/httpライブラリを使ったリクエストの受け付け

```
curl localhost:50002/handler

curl localhost:50002/multihandler_hello
curl localhost:50002/multihandler_world

curl localhost:50002/handlerfunc_hello
curl localhost:50002/handlerfunc_world

curl localhost:50002/chain_handlerfunc

curl localhost:50002/chain_handler
```

### processing

リクエストのデータ構造とその処理

```ocurl
curl localhost:50003/headers

curl -id "first_name=Kent&last_name=Fujii" localhost:50003/body

curl -id "hello=aaa&post=123" http://localhost:50003/form?thread=3

curl -i http://localhost:50003/fileupload -F "uploaded=@test.txt"

curl -i http://localhost:50003/formfile -F "uploaded=@test.txt"

curl -i localhost:50003/write
curl -i localhost:50003/writeheader
curl -i localhost:50003/redirect
curl -i localhost:50003/json

curl -i -c cookie.txt localhost:50003/set_cookie
curl -i -b cookie.txt localhost:50003/get_cookie

curl -i -c cookie.txt localhost:50003/set_message
curl -i -c cookie.txt -b cookie.txt localhost:50003/show_message
```

### displaying

テンプレートエンジンを使ったコンテンツの表示

```
curl localhost:50004/trigger_template

curl localhost:50004/random_number

curl localhost:50004/iterator

curl localhost:50004/set_dot

curl localhost:50004/include

curl localhost:50004/custom_function

curl localhost:50004/context_aware

curl localhost:50004/nested_1

curl localhost:50004/nested_2
```

### storing

データの記憶: 構造体・ファイル・データベース

```
curl localhost:50005/map_store

curl localhost:50005/read_write_file

curl localhost:50005/csv_store

curl localhost:50005/gob_store

curl localhost:50005/sql_store1

curl localhost:50005/sql_store2

curl localhost:50005/gorm_store
```

### service

GoによるWebサービスの作成: XMLおよびJSONの生成と解析

```
curl localhost:50006/xml_parsing_unmarshal_1

curl localhost:50006/xml_parsing_unmarshal_2

curl localhost:50006/xml_parsing_decoder

curl localhost:50006/xml_creating_marshal

curl localhost:50006/xml_creating_encoder

curl localhost:50006/json_parsing_unmarshal

curl localhost:50006/json_parsing_decoder

curl -i -X POST -H "Content-Type: application/json" -d '{"content":"Myfirstpost", "author":"KentFujii"}' localhost:50006/web_service/
curl -i -X GET localhost:50006/web_service/1
curl -i -X PUT -H "Content-Type: application/json" -d '{"content":"MySecondpost", "author":"KentFujii"}' localhost:50006/web_service/1
curl -i -X DELETE localhost:50006/web_service/1

```
### test

テスト用ライブラリを使ったアプリケーションのテスト

```

kubectl exec -it test -- bash -c "cd unittest ; go test"
kubectl exec -it test -- bash -c "cd unittest ; go test parallel_test.go -v -parallel 3"
kubectl exec -it test -- bash -c "cd unittest ; go test -bench ."

curl -i -X POST -H "Content-Type: application/json" -d '{"content":"Myfirstpost", "author":"KentFujii"}' localhost:50007/httptest_1/
kubectl exec -it test -- bash -c "cd httptest_1 ; go test -v"

curl -i -X POST -H "Content-Type: application/json" -d '{"content":"Myfirstpost", "author":"KentFujii"}' localhost:50007/httptest_2/
kubectl exec -it test -- bash -c "cd httptest_2 ; go test -v"

kubectl exec -it test -- bash -c "cd di ; go test -v"
```

### concurrency

ゴルーチンとチャネルを使った並行処理

```
kubectl exec -it concurrency -- bash -c "cd goroutine ; go test -v"

kubectl exec -it concurrency -- bash -c "cd wait_group ; go test -v"

kubectl exec -it concurrency -- bash -c "cd channel_wait ; go test -v"

kubectl exec -it concurrency -- bash -c "cd channel_message ; go test -v"

kubectl exec -it concurrency -- bash -c "cd channel_select ; go test -v"

curl http://localhost:50008/mosaic -F "tile_size=5" -F "uploaded=@cat.jpg" -o mosaic.jpg
```

### db

各プロジェクトが利用するDB

```
pg_isready -h localhost -U gp
psql -h localhost -U gp gp
```

### sidecar

### monitor

各プロジェクトのメトリクスを取得する

### dashboard

メトリクスを可視化
