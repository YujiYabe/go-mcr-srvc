# Contributing

go-mcr-srvcへの貢献を検討いただき、ありがとうございます。小さな修正を含め、Pull Requestを歓迎します。

## はじめに

大きな機能追加やアーキテクチャ変更を提案する場合は、実装前にIssueで目的と方針を共有してください。バグ報告では、再現手順、期待する結果、実際の結果、確認した環境を記載してください。

セキュリティ上の問題は公開Issueにせず、[`SECURITY.md`](SECURITY.md)の手順で報告してください。

## 開発環境

必要なソフトウェアと起動手順は[`README.md`](README.md)を参照してください。ローカル設定はサンプルから作成します。

```bash
cp backend/internal/env/local.env.example backend/internal/env/local.env
make install-tools
```

実際の認証情報、秘密鍵、token、ローカルデータはcommitしないでください。

## 実装方針

- `backend/internal`の依存方向は`1_framework → 2_adapter → 3_usecase → 4_domain`を維持してください。
- domain層にDB、HTTP、gRPC、`context.Context`、loggerなどの外部依存を持ち込まないでください。
- 業務ルールには適切なユニットテストを追加してください。
- 生成コードの変更は必要最小限にし、生成元も同時に更新してください。
- 既存コードの命名やフォーマットに合わせ、変更目的と無関係な修正を混在させないでください。

詳しい設計・実装規約は[`agents.md`](agents.md)と[`README.md`](README.md)を参照してください。

## 動作確認

Pull Requestを作成する前に、少なくとも次のコマンドを実行してください。

```bash
make gotest
make lint
make govulncheck
```

並行処理に関係する変更では、race detectorも実行してください。

```bash
cd backend && GOTOOLCHAIN=go1.27.0 go test -race ./...
```

## Pull Request

Pull Requestには次の内容を記載してください。

- 変更の目的と概要
- 主な設計判断やトレードオフ
- 実行したテストと結果
- 関連するIssue

Pull Requestは`main`ブランチを対象にし、CIが成功する状態にしてください。レビューで追加の変更をお願いする場合があります。
