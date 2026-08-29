# Database Gateway Layer

Clean Architecture のアダプターレイヤーにおけるデータベースゲートウェイを実装するパッケージです。

永続化層とドメイン層の橋渡しを行います。

## 主な役割

- usecase から見た永続化の入口
- PostgreSQL と Redis cache-aside の組み合わせ
- トランザクション境界の委譲
- cache hit / miss と cache invalidation の制御

## 特徴

- usecase へ PostgreSQL / Redis の詳細を漏らさない
- 永続化先を正とし、Redis はキャッシュとして扱う
- 永続化の詳細の隠蔽

## 実装内容

- PostgreSQL gateway への委譲
- キャッシュ制御
- validation_word_rules の cache-aside
- validation_word_rules 更新後の Redis key 削除

## 未実装

- MongoDB などの NoSQL gateway
- マイグレーション管理
- 分散トランザクション

## 使用方法

- ユースケース層からの要求に応じてデータの永続化を行う
- データベースからの取得結果をドメインモデルに変換
- トランザクション境界の制御
- データアクセスの最適化
