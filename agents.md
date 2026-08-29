# Agents Guide

このリポジトリで作業するエージェント向けの実装ガイドです。
設計方針は [README.md](README.md) と [doc/ca.drawio.png](doc/ca.drawio.png) を正とします。

## 基本方針

- クリーンアーキテクチャを優先する。
- DDD やマイクロサービスの考え方と衝突する場合も、依存方向と責務分離はクリーンアーキテクチャを優先する。
- `backend/internal` 以下は層を意識して実装する。
- 依存方向は外側から内側へ向ける。
- domain は framework / adapter / usecase に依存させない。
- usecase から usecase を呼び出す設計は避ける。

## 層ごとの責務

### framework

外部依存を扱う層。

- HTTP / gRPC / PubSub の入出力
- PostgreSQL / Redis / 外部 HTTP / gRPC client
- DB ライブラリ、認証 SDK、通信 SDK などの具体実装

この層では外部技術への変換や呼び出しを担当する。
業務ルールは置かない。

### adapter

入力差分や外部 gateway の違いを吸収する層。

- controller は framework の入力を usecase に渡す。
- gateway は usecase から見た永続化・外部サービス呼び出しの入口になる。
- 基本的に業務判断は書かない。

### usecase

アプリケーション固有の処理順序をまとめる層。

- context の状態確認
- 権限、トランザクション、gateway 呼び出し順序の制御
- domain service / entity / value object のルール適用
- 複数 repository / gateway 操作の orchestration

業務ルールそのものを直接増やしすぎない。
複数オブジェクト間の純粋な業務判断は domain service へ寄せる。

### domain

ドメインロジックを凝集する層。

- `primitive_object`: 基本型のバリデーションやフォーマット
- `type_object`: ドメイン固有の値オブジェクト
- `group_object`: Entity / Aggregate 相当
- `service_object`: Entity や Value Object 単体に収まらない業務ルール

domain は左側の層から参照される。
domain から usecase / adapter / framework を呼び出さない。

## DDD の置き方

- Entity / Aggregate は `group_object` に置く。
- 識別子、ライフサイクル、状態変更に関する振る舞いは Entity に寄せる。
- 値そのものの制約は `type_object` に置く。
- 複数 Entity / Value Object 間の整合性は `service_object` に置く。
- DB テーブル名ではなく、業務上の言葉を package / type / method 名へ反映する。

例:

- `User`
- `UserEmployment`
- `PrimaryEmploymentAssignmentPolicy`
- `UpdateUserProfileWithPrimaryEmployment`

## Usecase と Domain の境界

usecase に置くもの:

- トランザクション境界
- gateway / repository 呼び出し順序
- 外部サービス呼び出しの orchestration
- request context や timeout などアプリケーション処理上の制御

domain に置くもの:

- ユーザーが更新可能な状態か
- 主所属として割り当て可能か
- ID、メールアドレス、権限などの値制約
- Entity / Aggregate の不変条件

判断に迷う場合は、外部 I/O や処理順序に関係するものは usecase、純粋な業務整合性は domain に置く。

## DB 実装

- `backend/internal/1_framework/out/db/postgres_client` はテーブル単位にファイルを分割する。
- トランザクションなどテーブル横断の処理は共通ファイルへ分ける。
- PostgreSQL / GORM 固有の処理は framework 層に閉じ込める。
- usecase / domain に GORM や SQL の型を漏らさない。
- GORM の DB model はテーブル表現に専念させ、`json` タグは付けない。HTTP/API の JSON 表現は DTO 側で定義する。

## Redis キャッシュ

- PostgreSQL などの永続化先を正とし、Redis は cache-aside 方式のキャッシュとして扱う。
- キャッシュの TTL は原則として無期限にする。
- POST / PUT / DELETE などでデータが変化する場合は、先に永続化先へ反映し、成功後に関連する Redis キーを削除する。
- 永続化に失敗した場合は Redis キーを削除しない。
- キャッシュ削除後の次回読み込みでは、永続化先から最新版を取得して Redis に保存する。
- 更新処理内でキャッシュ値を直接書き換えず、キー削除による無効化を基本とする。

## トランザクション

- 複数テーブルを同時更新する場合は usecase で transaction 境界を張る。
- transaction 内では gateway 経由で複数の DB 操作を呼び出す。
- domain は transaction を知らない。
- rollback の判断は gateway の error を usecase へ返して行う。

## 命名

- `GatewayDB` のような技術寄りの名前を application 側へ広げすぎない。
- 将来複数テーブルを扱う repository / gateway は、特定テーブル名へ寄せすぎない。
- HTTP DTO、gRPC message、DB model は外側の都合として扱う。
- domain/usecase の名前は業務語彙を優先する。
- 層をまたいで依存を受け渡す interface は、プロジェクト内ルールとして `ToXXXX` 形式で統一する。
- `ToXXXX` は「現在の層から XXXX 側へ依存を向けるための境界」を表す。
- 一般的な Go の慣習で責務名の interface が自然な場合でも、このリポジトリでは既存規約との一貫性を優先する。

## テスト

- type_object / group_object / service_object の業務ルールにはユニットテストを置く。
- usecase では transaction 境界、呼び出し順序、gateway が呼ばれない条件をテストする。
- framework の外部依存は必要に応じて integration test として切り分ける。
- 変更後は少なくとも以下を実行する。

```bash
make gotest
make lint
```

## 変更時の注意

- 既存の層構造を崩さない。
- domain へ context、DB、HTTP、gRPC、logger などの外部都合を持ち込まない。
- generated code は必要最小限に触る。
- `task.memo` や作業メモはユーザーから明示されない限りマージ対象にしない。
