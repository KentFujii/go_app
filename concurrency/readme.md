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


