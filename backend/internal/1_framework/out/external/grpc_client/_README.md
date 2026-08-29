# gRPC Client

gRPC マイクロサービス間通信を管理するパッケージです。  
フレームワーク層でのサービス間連携の実装を提供します。

## 主な機能

- gRPC サービスとの通信
- コネクション管理
- request context の metadata 伝搬
- gRPC message と domain model の変換

## 特徴

- adapter layer の external gateway から利用される
- gRPC の具体実装を framework layer に閉じ込める
- unary RPC を利用する

## 実装内容

- gRPC 接続設定
- プロトコルバッファの実装
- `GetUserListByCondition` の呼び出し
- metadata への RequestContext 変換
- エラーハンドリング

## 未実装

- streaming RPC
- サーキットブレーカー
- client side load balancing
- クライアント側リトライ
- メトリクス収集

## 使用方法

このクライアントは、アダプターレイヤーの gRPC ゲートウェイによって使用され、効率的なマイクロサービス間通信を実現します。  
Clean Architecture の原則に従い、このパッケージの実装の詳細は上位レイヤーから隠蔽されています。
