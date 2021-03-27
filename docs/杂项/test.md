score 默认值

[code](http)



```
docker run -p 9000:9000 \
  -e "MINIO_ROOT_USER=gqzcl" \
  -e "MINIO_ROOT_PASSWORD=systemlogin" \
  minio/minio server /data
```