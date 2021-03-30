注册 score 默认值

文章

在线编译

连接堵塞(长连接？)



[code](http)



```
docker run -p 9000:9000 \
  -e "MINIO_ROOT_USER=gqzcl" \
  -e "MINIO_ROOT_PASSWORD=systemlogin" \
  minio/minio server /data
```