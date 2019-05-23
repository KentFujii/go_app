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


