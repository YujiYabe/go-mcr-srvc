# Primitive Object

基本データ型のドメインロジックを提供するパッケージです。  
各型に対してバリデーションやフォーマット処理を実装しています。

## 実装されている型

- PrimitiveInt - 整数型
- PrimitiveInt64 - 64bit整数型
- PrimitiveUint - 符号なし整数型
- PrimitiveString - 文字列型
- PrimitiveSliceInt - 整数配列型
- PrimitiveSliceString - 文字列配列型

## 主な機能

### バリデーション

- 数値型: 最小/最大桁数チェック (minDigit/maxDigit)
- 文字列型: 最小文字列長チェック (minLength)
- 配列型: nil 状態チェック、要素のバリデーション

### 型変換

- スライス型の値取得メソッド (ToSliceInt(), ToSliceString())
- 各プリミティブ型の値取得メソッド

### エラーハンドリング

- バリデーションエラーの保持と取得
- エラー状態の管理

## 使用例

このパッケージは value_object パッケージの基盤として使用され、より具体的なドメインモデルの実装に活用されます。
