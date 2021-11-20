# golang-server-like-clean-arch

# アーキテクチャ
- クリーンアーキテクチャライクな構成を採用
- ドメインモデルを中心に構築することを意識
- [bxcodec/go-clean-arch](https://github.com/bxcodec/go-clean-arch)

## Domain
- ビジネルロジックを実装
- Valueオブジェクトを利用
- 必要に応じて配列をラップした`List`モデルを実装
- コンストラクタには`New`関数を内部には`new()`を利用
- イミュータブル厳守

## Controller
- 外部内部とのやりとり(入出力)を管理
- スキーマ駆動により実装

## Gateway(Repository)
- Domainモデルを永続化を管理
- DB管理には[ent.](https://entgo.io/)を採用

## Usecase
- アプリケーションのユースケースを管理
- シーケンス図のコーディング
- `for`による動作/`if`利用のための詳細ロジックはDomainモデルに記述できないか検討

## ディレクトリ構成
```bash
├── app                 # アプリケーション(メイン実装管理)
│   ├── main.go         # アプリケーション起点
│   ├── domain          # ドメイン(ビジネスロジック)
│   │   ├── xxx         # Xxxモデル
│   │   └── yyy         # Yyyモデル
│   ├── xxx
│   │   ├── controller  # Xxxモデルにおける入出力変換機能
│   │   ├── gateway     # Xxxモデルにおける永続化機能
│   │   └── usecase     # Xxxモデルにおけるユースケース
│   └── yyy
│       ├── controller  # Yyyモデルにおける入出力変換機能
│       ├── gateway     # Yyyモデルにおける永続化機能
│       └── usecase     # Yyyモデルにおけるユースケース
│
├── ent                 # entによるDB管理
├── logger              # zapによるログ形式管理
├── migration           # entを利用したDBマイグレーション管理
├── mock                # テスト用のモック実装(自動生成)
├── proto               # grpcを使用したAPIの型定義
└── script              # 各種操作のスクリプト
```

## ファイル
```bash
├── main.go
├── domain
│   └── xxx
│       │   # ドメインモデル
│       ├── xxx.go
│       │   # ドメインモデルの複数形
│       ├── xxx_list.go
│       │   # ドメインのテスト
│       ├── xxx_test.go
│       │   # 永続化機能の抽象モデル
│       ├── xxx_repository.go
│       │   # ユースケースの抽象モデル
│       └── xxx_usecase.go
└── xxx
    │   # クライアント間通信の実装
    ├── controller
    │   │   # 外部モデル -> domainモデル
    │   ├── xxx_controller.go
    │   │   # 各種操作メソッド実装とそのテスト
    │   ├── xxx_controller_create.go
    │   ├── xxx_controller_create_test.go
    │   ├── xxx_controller_read.go
    │   ├── xxx_controller_read_test.go
    │   ├── xxx_controller_update.go
    │   ├── xxx_controller_update_test.go
    │   ├── xxx_controller_delete.go
    │   ├── xxx_controller_delete_test.go
    │   │   # domainモデル -> 外部モデル
    │   └── xxx_translator.go
    │   # repositoryの実装
    ├── gateway
    │   │   # entを利用した実装とそのテスト
    │   ├── xxx_gateway.go
    │   ├── xxx_gateway_test.go
    │   │   # entモデル <-> domainモデル
    │   └── xxx_entity.go
    │   # usecaseの実装
    └── usecase
        │   # 各種操作メソッド実装とそのテスト
        ├── xxx_usecase.gp
        ├── xxx_usecase_create.go
        ├── xxx_usecase_create_test.go
        ├── xxx_usecase_read.go
        ├── xxx_usecase_read_test.go
        ├── xxx_usecase_update.go
        ├── xxx_usecase_update_test.go
        ├── xxx_usecase_delete.go
        └── xxx_usecase_delete_test.go
```

# 命名規則
## ディレクトリ(Package)
- 名詞1単語(単数形)
## ファイル
- スネークケースを利用
## 構造体(クラス)
- アッパーキャメルケース(パスカルケース)
## 変数名
- ローワーキャメルケース(キャメルケース)
- 配列は複数形を利用(xxxListは利用しない)
## メソッド(レシーバー)
- 動詞はじまり厳守
- 内部利用はローワーキャメルケース(キャメルケース)
- 外部利用はアッパーキャメルケース(パスカルケース)

## レシーバ名
- 構造体に利用した命名から英語1文字or2文字を利用

例)
```go
type Xxx struct {
    hoge string
    fuga string
}

func New() Xxx {
    xxx := new(Xxx)
    xxx.hoge = "hoge"
    xxx.fuga = "fuga"
    return *xxx
}

(x Xxx) execute() {}

type XxxList []Xxx

func NewList(xxxs []Xxx) XxxList {
    xxxList := XxxList(xxxs)
    return xxxList
}
```
※参照:[他言語プログラマが最低限、気にすべきGoのネーミングルール](https://zenn.dev/keitakn/articles/go-naming-rules)

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
