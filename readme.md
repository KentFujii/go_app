# go_app

https://github.com/sausheong/gwp

https://github.com/GoogleContainerTools/skaffold/tree/master/examples/microservices

https://qiita.com/tomoyamachi/items/660bd7bb3afff8340307

## 構成

下記それぞれのディレクトリがそれぞれ一つのPodとしてk8sに収まるようになっている

### primary

簡単なWebアプリケーション

### handling

net/httpライブラリを使ったリクエストの受け付け

### processing

リクエストのデータ構造とその処理

### displaying

テンプレートエンジンを使ったコンテンツの表示

## storing

データの記憶: 構造体・ファイル・データベース

## web_service

GoによるWebサービスの作成: XMLおよびJSONの生成と解析

## test_app

テスト用ライブラリを使ったアプリケーションのテスト

## concurrency

ゴルーチンとチャネルを使った並行処理
