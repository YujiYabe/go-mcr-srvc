[<-戻る](../../../../README.md#ドメインロジックの凝集)

# Type Object

- primitive_object を利用して、ドメイン固有の値オブジェクトを定義するパッケージです。
- type_object は、group_object でより大きなドメインモデルを構築する際の構成要素として使用されます。
- 各 type_object は、それぞれのドメインルールを独立して実装し、ビジネスロジックの整合性を保証します。

## 特徴

- 不変性(Immutable)を保証
- ドメインルールに基づいたバリデーション
- 自己完結的な振る舞いの実装

## 実装例

### 実装されている主な値

- `Email`
- `Name`
- `ID`
- `UserID`
- `TenantID`
- `Permission`
- `PermissionList`
- `AccessToken`
- `ClientID`
- `ClientSecret`
- `ClientIP`
- `UserAgent`
- `Locale`
- `TimeZone`
- `TraceID`
- `RequestStartTime`
- `TimeOutMillSecond`

### 例

- `Email` はメールアドレス形式を検証します。
- `Name` は文字列長と blacklist を検証します。
- `TraceID` は未指定時に UUID を生成します。
- `RequestStartTime` と `TimeOutMillSecond` はリクエスト伝搬時の timeout 計算に使います。


[<-戻る](../../../../README.md#ドメインロジックの凝集)
