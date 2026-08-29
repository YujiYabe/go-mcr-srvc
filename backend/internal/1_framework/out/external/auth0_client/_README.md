# Auth0 Client

Auth0 認証サービスとの連携を管理するパッケージです。  
フレームワーク層での Auth0 token endpoint 呼び出しを提供します。

## 主な機能

- Auth0 token endpoint との通信
- client credentials による access token 取得
- Credential domain model から Auth0 request payload への変換
- Auth0 response から AccessToken value object への変換

## 特徴

- HTTP client の具体実装を framework layer に閉じ込める
- context による request cancellation を利用する
- adapter layer の auth gateway から利用される

## 実装内容

- token URL / audience / grant type の設定
- access token 取得 API 呼び出し
- エラーハンドリング

## 未実装

- ユーザー認証フロー
- token 検証
- token refresh
- ユーザー情報取得
- ロールベースアクセス制御
- セッション管理
- レート制限対応

## 使用方法

このクライアントは、アダプターレイヤーの認証ゲートウェイによって使用され、Auth0 から access token を取得します。  
Clean Architecture の原則に従い、このパッケージの実装の詳細は上位レイヤーから隠蔽されています。
