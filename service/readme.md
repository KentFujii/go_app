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
