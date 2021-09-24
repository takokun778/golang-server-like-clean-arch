# golang-server-like-clean-arch

## ディレクトリ構成
```bash
.
├── Makefile
├── README.md
├── app
│   ├── common
│   │   ├── controller
│   │   │   └── error_translator.go
│   │   └── gateway
│   │       └── database.go
│   ├── domain
│   │   ├── common
│   │   │   ├── error.go
│   │   │   ├── id.go
│   │   │   └── time.go
│   │   ├── fuga
│   │   │   ├── fuga.go
│   │   │   ├── fuga_list.go
│   │   │   ├── fuga_repository.go
│   │   │   ├── fuga_test.go
│   │   │   ├── fuga_usecase.go
│   │   │   ├── name.go
│   │   │   └── number.go
│   │   ├── hoge
│   │   │   ├── hoge.go
│   │   │   ├── hoge_list.go
│   │   │   ├── hoge_repository.go
│   │   │   ├── hoge_test.go
│   │   │   ├── hoge_usecase.go
│   │   │   ├── name.go
│   │   │   └── number.go
│   │   └── logger
│   │       └── logger.go
│   ├── fuga
│   │   ├── controller
│   │   │   ├── fuga_controller.go
│   │   │   ├── fuga_controller_test.go
│   │   │   └── fuga_translator.go
│   │   ├── gateway
│   │   │   └── fuga_gateway.go
│   │   └── usecase
│   │       ├── fuga_create_usecase.go
│   │       ├── fuga_delete_usecase.go
│   │       ├── fuga_fetch_all_usecase.go
│   │       ├── fuga_fetch_usecase.go
│   │       ├── fuga_update_usecase.go
│   │       └── fuga_usecase.go
│   ├── hoge
│   │   ├── controller
│   │   │   ├── hoge_controller.go
│   │   │   ├── hoge_controller_test.go
│   │   │   └── hoge_translator.go
│   │   ├── gateway
│   │   │   └── hoge_gateway.go
│   │   └── usecase
│   │       ├── hoge_create_usecase.go
│   │       ├── hoge_delete_usecase.go
│   │       ├── hoge_fetch_all_usecase.go
│   │       ├── hoge_fetch_usecase.go
│   │       ├── hoge_update_usecase.go
│   │       └── hoge_usecase.go
│   └── main.go
├── docker-compose.yml
├── ent
│   ├── client.go
│   ├── config.go
│   ├── context.go
│   ├── ent.go
│   ├── enttest
│   │   └── enttest.go
│   ├── fuga
│   │   ├── fuga.go
│   │   └── where.go
│   ├── fuga.go
│   ├── fuga_create.go
│   ├── fuga_delete.go
│   ├── fuga_query.go
│   ├── fuga_update.go
│   ├── generate.go
│   ├── hoge
│   │   ├── hoge.go
│   │   └── where.go
│   ├── hoge.go
│   ├── hoge_create.go
│   ├── hoge_delete.go
│   ├── hoge_query.go
│   ├── hoge_update.go
│   ├── hook
│   │   └── hook.go
│   ├── migrate
│   │   ├── migrate.go
│   │   └── schema.go
│   ├── mutation.go
│   ├── predicate
│   │   └── predicate.go
│   ├── runtime
│   │   └── runtime.go
│   ├── runtime.go
│   ├── schema
│   │   ├── fuga.go
│   │   └── hoge.go
│   └── tx.go
├── go.mod
├── go.sum
├── migration
│   └── main.go
├── proto
│   ├── fuga
│   │   ├── fuga.pb.go
│   │   ├── fuga_grpc.pb.go
│   │   └── v1
│   │       └── fuga.proto
│   └── hoge
│       ├── hoge.pb.go
│       ├── hoge_grpc.pb.go
│       └── v1
│           └── hoge.proto
└── script
    ├── dev-migrate.sh
    ├── dev-run.sh
    └── test.sh

```

## テスト
以下2点のテストコードを各ドメインモデルで用意する(必須)
- app/domain/${domain_name}_test.go  
    ビジネスロジックの正当性検証
- app/\${domain\_name}/controller/${domain_name}_controller_test.go  
    リクエスト単位の確認 ValueObjectのロジック正当性(バリデーション)もここで検証

## サーボパーティライブラリ
domain/usecase配下ではなるべく依存しないようにする(標準ライブラリで頑張る)

### サーバー
- [gRPC](https://grpc.io/)

### データベース
- [ent](https://entgo.io/)
- [pg](https://pkg.go.dev/github.com/lib/pq)

### ログ  
- [zap](https://pkg.go.dev/go.uber.org/zap)

### テスト
- [testify](https://pkg.go.dev/github.com/stretchr/testify)

## 参照
- [go-clean-arch](https://github.com/bxcodec/go-clean-arch)
