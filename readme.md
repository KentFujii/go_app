# go_app

https://github.com/sausheong/gwp

https://github.com/GoogleContainerTools/skaffold/tree/master/examples/microservices

https://qiita.com/tomoyamachi/items/660bd7bb3afff8340307

## 構成

下記それぞれのディレクトリがそれぞれ一つのPodとしてk8sに収まるようになっている

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
curl -id "hello=aaa&post=123" http://localhost:50003/process?thread=3
curl -i http://localhost:50003/fileupload -F "uploaded=@test.txt"
curl -i http://localhost:50003/formfile -F "uploaded=@test.txt"
curl localhost:50003/write
curl localhost:50003/writeheader
curl localhost:50003/redirect
curl localhost:50003/json
curl localhost:50003/set_cookie
curl localhost:50003/get_cookie
```

### displaying

テンプレートエンジンを使ったコンテンツの表示

### storing

データの記憶: 構造体・ファイル・データベース

pg_isready -h localhost -U gp
psql -h localhost -U gp gp

### service

GoによるWebサービスの作成: XMLおよびJSONの生成と解析

### test_app

テスト用ライブラリを使ったアプリケーションのテスト

### concurrency

ゴルーチンとチャネルを使った並行処理
