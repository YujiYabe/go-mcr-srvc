# Security Policy

## Supported Versions

このリポジトリは学習・検証用のリファレンス実装です。長期サポート版は設けておらず、セキュリティ修正は原則として`main`ブランチにのみ適用します。

| Version | Supported |
| --- | --- |
| `main` | Yes |
| 過去のcommit、tag、fork | No |

## Reporting a Vulnerability

脆弱性を発見した場合は、公開IssueやPull Requestには詳細を記載せず、GitHubの[Private vulnerability reporting](https://github.com/YujiYabe/go-mcr-srvc/security/advisories/new)から報告してください。

報告には、可能な範囲で次の情報を含めてください。

- 影響を受ける箇所と想定される影響
- 再現手順または概念実証
- 確認したcommit、環境、設定
- 回避策や修正案（把握している場合）

受領後7日以内を目安に確認結果を返信します。修正方針と公開時期は、影響範囲を確認したうえで報告者と調整します。修正が公開されるまでは、脆弱性に関する情報を公開しないようお願いします。

このプロジェクトは本番利用を想定していません。サンプルの認証情報やローカル用設定を本番環境で使用しないでください。
