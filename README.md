# Modular Monolith Sample

このプロジェクトは、Go によるモジュラーモノリスアーキテクチャのサンプル実装です。Connect-RPC（gRPC 互換）を使用した API サーバーとクライアントを提供します。

## アーキテクチャ

このプロジェクトは以下のアーキテクチャ原則に従っています：

- モジュラーモノリス
- クリーンアーキテクチャ
- DDD（ドメイン駆動設計）

### ディレクトリ構造

```
.
├── cmd/                  # エントリーポイント
│   ├── client/           # クライアントアプリケーション
│   └── server/           # サーバーアプリケーション
├── gen/                  # 生成されたコード
├── internal/             # 内部パッケージ
│   ├── client/           # クライアントライブラリ
│   └── server/           # サーバー実装
│       ├── lib/          # 共通ライブラリ
│       ├── module/       # ビジネスモジュール
│       ├── presentation/ # プレゼンテーション層
│       ├── registry/     # 依存性注入
│       └── scenario/     # ユースケースシナリオ
└── proto/                # Protocol Buffersの定義
```

### モジュール構成

- `user`: ユーザー管理モジュール
- `contract`: 契約管理モジュール

## 主な機能

- ユーザー作成（契約も同時に作成）
- ユーザー情報取得
- 契約情報取得

## 技術スタック

- Go 1.21+
- Connect-RPC
- Protocol Buffers
- buf (Protocol Buffers のツールチェーン)

## セットアップ

### 必要条件

- Go 1.21 以上
- buf

### インストール

```bash
# リポジトリのクローン
git clone https://github.com/fumiyakk/modular-monolith-sample.git
cd modular-monolith-sample

# 依存関係のインストール
go mod download

# Protocol Buffersコードの生成
buf generate
```

### 実行方法

サーバーの起動:

```bash
go run cmd/server/main.go
```

クライアントの実行:

```bash
go run cmd/client/main.go
```

#

## アーキテクチャの詳細

### レイヤー構造

1. プレゼンテーション層 (`internal/server/presentation/`)

   - API ハンドラー
   - リクエスト/レスポンスの変換

2. シナリオ層 (`internal/server/scenario/`)

   - 複数のユースケースを組み合わせた処理
   - トランザクション管理

3. モジュール層 (`internal/server/module/`)

   - ビジネスロジック
   - ドメインモデル
   - ユースケース
   - リポジトリインターフェース

4. インフラストラクチャ層 (`internal/server/lib/`)
   - データベース実装
   - ミドルウェア
   - 共通ライブラリ

### 依存性注入

- `internal/server/registry/`で依存関係を管理
- モジュール間の疎結合を実現

## 注意事項

- このプロジェクトはサンプル実装です
- 本番環境での使用には適切なセキュリティ対策が必要です
- メモリ内データストアを使用しているため、永続化は実装されていません
