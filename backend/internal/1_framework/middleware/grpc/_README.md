# gRPC Middleware

gRPC サービスのミドルウェアを管理するパッケージです。  
フレームワーク層での横断的な関心事を実装します。

## 主な機能

- リクエスト/レスポンスのインターセプト
- RequestContext と gRPC metadata の相互変換
- gRPC message と domain model の変換補助

## 特徴

- エラーハンドリング
- コンテキスト管理
- TraceID の context 設定

## 実装内容

- unary server interceptor
- incoming metadata から RequestContext への変換
- RequestContext から outgoing metadata への変換
- User の gRPC parameter / domain model 変換

## 未実装

- 認証・認可
- メトリクス収集
- レート制限
- リカバリー処理
- OpenTelemetry などの分散トレーシング

## 使用方法

このパッケージは、gRPC サーバーの設定時に使用され、一貫性のあるクロスカッティングコンサーンを提供します。  
Clean Architecture の原則に従い、このパッケージの実装の詳細は上位レイヤーから隠蔽されています。
