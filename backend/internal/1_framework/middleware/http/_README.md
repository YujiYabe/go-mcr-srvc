# HTTP Middleware

HTTP サーバーのミドルウェアを管理するパッケージです。  
フレームワーク層での横断的な関心事を実装します。

## 主な機能

- RequestContext の作成
- HTTP request からの ClientIP / UserAgent / Locale / TimeZone 取得
- TraceID の context 設定
- JWT 検証

## 特徴

- チェーン可能なミドルウェア
- エラーハンドリング
- Auth0 JWKS を使った JWT 検証に対応
- HMAC secret を使った JWT 検証に対応

## 実装内容

- `ContextMiddleware`
- `JWTMiddleware`
- `JWTMiddlewareAuth0`
- RequestContext を request context に保存

## 未実装

- 認可
- メトリクス収集
- レート制限
- CORS 制御
- セッション管理
- リカバリー処理
- OpenTelemetry などの分散トレーシング

## 使用方法

このパッケージは、HTTP サーバーの設定時に使用され、一貫性のあるクロスカッティングコンサーンを提供します。  
Clean Architecture の原則に従い、このパッケージの実装の詳細は上位レイヤーから隠蔽されています。
