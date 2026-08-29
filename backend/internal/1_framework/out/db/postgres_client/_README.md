# PostgreSQL Client

PostgreSQL データベースとの接続を管理するパッケージです。  
フレームワーク層でのデータベース接続の実装を提供します。

## 主な機能

- データベース接続の確立
- トランザクション制御
- クエリの実行
- ドメインモデルと DB model の変換

## 特徴

- 設定の外部化
- 接続確立時のリトライ
- エラーハンドリング
- PostgreSQL / GORM 固有処理の隠蔽

## 実装内容

- データベース接続設定
- GORM によるクエリ実行
- context 経由での transaction 伝搬
- users / user_employments / validation_word_rules の永続化
- validation_word_rules の upsert / update / delete

## 未実装

- クエリ単位のリトライ
- クエリ単位のタイムアウト設定
- メトリクス収集
- Transactional Outbox

## 使用方法

このクライアントは、アダプターレイヤーのデータベースゲートウェイによって使用され、実際のデータベース操作を行います。  
Clean Architecture の原則に従い、このパッケージの実装の詳細は上位レイヤーから隠蔽されています。
