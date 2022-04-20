## designファイルのビルド
goa gen goa-sample/design
goa example goa-sample/design

go build ./cmd/controller
./controller

## docker起動
docker-compose build --no-cache
docker-compose up -d

## アクセス
### goa-sample API
http://localhost:8000

### Swagger
http://localhost:8002

### phpmyadmin
http://localhost:9000

### mailcatcher
http://localhost:9000