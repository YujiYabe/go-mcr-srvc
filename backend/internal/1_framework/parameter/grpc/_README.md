# gRPC Parameter

gRPC リクエストのパラメータ処理を管理するパッケージです。  
フレームワーク層でのメッセージパラメータのバリデーションと変換を提供します。

## 主な機能

- プロトコルバッファメッセージの処理
- generated Go code の配置

## 特徴

- 型安全な変換
- gRPC message 定義を framework layer に閉じ込める

## 実装内容

- `user.proto`
- メッセージ構造体の定義
- gRPC service interface の generated code

## 未実装

- ストリーミング RPC
- 共通 validator
- デフォルト値の設定

## 使用方法

このパッケージは、gRPC サービスハンドラーで使用され、クリーンで一貫性のあるパラメータ処理を実現します。  
Clean Architecture の原則に従い、このパッケージの実装の詳細は上位レイヤーから隠蔽されています。
