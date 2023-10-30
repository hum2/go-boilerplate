# コードジェネレータ
API作成で必要なファイルをCLIで生成します。

## コントローラ作成
```shell
# go run cmd/generator/*.go controller ${タグ名}
# `/user/offer`エンドポイントのGetメソッド作成
go run cmd/generator/*.go controller user/offer
# `/user/offer`エンドポイントのGET,POST,PUT,DELETEメソッド作成
go run cmd/generator/*.go controller user/offer --method=GET,POST,PUT,DELETE
```

※ 現時点では追加されたエンドポイントのYAMLファイルを`openapi/oapi-codegen.yml`に記載する必要があります。
```yaml
tags:
  - name: ${タグ名}
    description: Change me
# ...
paths:
# ...
  /${タグ名}:
    $ref: ./assets/path/${タグ名}/root.yaml
# ...
components:
  parameters:
  # Parameter
  # TODO: 必要に応じて記載
# ...
  schemas:
  # Request
  # TODO: 必要に応じて記載
# ...
  # Response
  # TODO: 必要に応じて記載
```

## ドメイン作成
```shell
go run cmd/generator/*.go domain user/offer
```

## リポジトリ作成
```shell
go run cmd/generator/*.go repository user/offer
```
