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


