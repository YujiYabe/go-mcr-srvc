# gRPC Parameter

gRPC リクエストのパラメータ処理を管理するパッケージです。  
フレームワーク層でのメッセージパラメータのバリデーションと変換を提供します。

## 主な機能

- プロトコルバッファメッセージの処理
- メタデータの処理
- ストリーミングパラメータの処理
- コンテキスト情報の処理

## 特徴

- 型安全な変換
- バリデーションルール
- メッセージの整合性チェック
- カスタムバリデータ

## 実装内容

- メッセージ構造体の定義
- バリデーションロジック
- 型変換処理
- デフォルト値の設定
- エラーハンドリング

## 使用方法

このパッケージは、gRPC サービスハンドラーで使用され、クリーンで一貫性のあるパラメータ処理を実現します。  
Clean Architecture の原則に従い、このパッケージの実装の詳細は上位レイヤーから隠蔽されています。
