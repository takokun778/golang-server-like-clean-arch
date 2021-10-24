# golang-server-like-clean-arch

# ディレクトリ構成
- クリーンアーキテクチャライクな構成を採用
- ドメインモデル中心でコーディングすることを意識
```bash
├── app
│   ├── main.go
│   ├── domain
│   │   └── xxx
│   │       ├── xxx.go # ドメインモデル
│   │       ├── xxx_repostiroy.go # データストア処理のinterface
│   │       └── xxx_usecase.go # ドメインモデルにおけるユースケース処理のinterface
│   └── xxx
│       ├── controller
│       │   ├── xxx_controller.go # 外->内の変換
│       │   └── xxx_translator.go # 内->外の変換
│       ├── gateway
│       │   └── xxx_gateway.go # repository interface の実装
│       └── usecase
│           └── xxx_usecase.go # usecase interface の実装
├── ent # entによるDB管理
├── logger # 共通で利用するlogger
├── mock # テスト用のmock実装
├── migration # entを利用したDBマイグレーションの実装
├── proto # grpcを使用したAPIの型定義
└── script # 各種操作のスクリプト
```

# [ent](https://entgo.io/)の使い方
1. 下記コマンドにて必要となるテーブルを追加する(単数形でOK)
```bash
go run entgo.io/ent/cmd/ent init ${Xxx}
```
2. `app/ent/schema/xxx.go`というファイルが作成される
3. カラム要素をコーディングする
例)
```go
package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/field"
)

// Fields of the Xxx.
func (Xxx) Fields() []ent.Field {
    return []ent.Field{
        field.UUID("id", uuid.UUID{}),
		field.String("name"),
		field.Time("createdAt").Default(time.Now),
		field.Time("updatedAt").Default(time.Now).UpdateDefault(time.Now),
    }
}
```

# [mock](https://github.com/golang/mock)の使い方
1. 各ドメインモデルにつき`usecase` `repository` それぞれに次の1文をトップに追加
- app/domain/xxx/xxx_usecase.go
- app/domain/xxx/xxx_repository.go
```go
//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../../mock/$GOPACKAGE/$GOFILE
package xxx
```
2. `Makefile`に次の1文を追加
```
go generate ./app/domain/xxx
```
3. `make mock`を実行

# サードパーティライブラリ
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

# 参照
- [go-clean-arch](https://github.com/bxcodec/go-clean-arch)
