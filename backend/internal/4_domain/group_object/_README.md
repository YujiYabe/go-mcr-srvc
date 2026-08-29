[<-戻る](../../../../README.md#ドメインロジックの凝集)

# Group Object

複数の type_object を組み合わせて、より大きなドメインモデルを構築するパッケージです。

## 特徴
- 複数の type_object を論理的にグループ化
- DDDのentityとaggregateに相当
- ドメインモデルの整合性を保証
- ビジネスルールの集約

## 実装されている型

- `User`
- `UserEmployment`
- `UserList`
- `Credential`

RequestContext は `backend/internal/1_framework/middleware/request_context` で管理します。group_object には置きません。

[<-戻る](../../../../README.md#ドメインロジックの凝集)
