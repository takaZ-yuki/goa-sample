# 環境構築手順
1. goaのファイル自動生成を実行
```bash
goa gen goa-sample/design
goa example goa-sample/design
```

2. dockerの起動
```bash
docker-compose up -d
```

## アクセス
### goa-sample API
http://localhost:8000

### Swagger
http://localhost:8002

### phpmyadmin
http://localhost:9000

### mailcatcher
http://localhost:9000

## APIへの接続について
rest.httpファイルにVSCodeのREST拡張仕様のファイルを用意してます
そちらをお使いください