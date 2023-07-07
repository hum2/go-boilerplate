# golang-boilerplate

## 構成
```text
.
├── cmd                         ・・・ コマンド実行用ディレクトリ
├── docs                        ・・・ ドキュメント格納用ディレクトリ
├── ent                         ・・・ ent/ent生成ファイル
├── internal                    ・・・ goのソースコードディレクトリ
│   ├── cmd                     ・・・ コマンド実行用ディレクトリ
│   ├── domain                  ・・・ ドメイン層
│   ├── infrastructure          ・・・ インフラ層
│   ├── interface               ・・・ プレゼンテーション層
│   │   └── controller          ・・・ コントローラー
│   └── usecase                 ・・・ ユースケース層
└── openapi                     ・・・ openapi定義ファイル格納ディレクトリ
    └── assets                  ・・・ openapi定義を分割したファイルの格納ディレクトリ
```

## 各層の概要

![図解](https://storage.googleapis.com/zenn-user-upload/be8ff7c2596e-20220805.png)


## 各種コマンド

### 起動
```shell
docker compose build --no-cache
docker compose up -d
```

### 停止
```shell
docker compose down -v
```

### リリースビルド
```shell
# development環境
docker build . --target dev -t backend-dev:latest --no-cache
# staging環境
docker build . --target stg -t backend-stg:latest --no-cache
# production環境
docker build . --target prd -t backend-prd:latest --no-cache
```

### lint実行
```shell
docker compose exec backend make lint
```

### テスト実行
```shell
docker compose exec backend make test
```

### 自動生成ファイルの再生成
```shell
docker compose exec backend make generate_all
```
