# つまりハンバーガーはclean architectureだったんだよ！
ΩΩΩ < ﾅ、ﾅﾝﾀﾞｯﾃｰ

# この記事で説明したいこと
clean architectureを(実際の実装を含めて)マク○ナル○に例えて説明したいと思います！   


# なんでマク○ナル○？
早い・安い・美味い、エンジニアの気絶のお供と言えばマク○ナル○だからです！（関係ない）
キャッチーなタイトルだと書くモチベがあがりますよね!   

ファーストフードでは誰がオーダーを受けようが誰が料理しようが基本的に同じ味で提供されますね。
例えばキッチンのスタッフが入れ替わっても、レジの人は気にせずキッチンへのオーダーを通す事ができますが、これはclean architectureの疎結合に似ています。  
またclean architecture(に限らない話ですが)はdbの依存関係を極力少なくしているのでdbの移行などでも最小限の工数で済む。というメリットがあります。
今回inputのインターフェースを４つ。outputのインターフェースを計５つ盛りだくさんで実装しています。

# この記事のターゲット
clean architectureについて解説してる有用な記事はたくさんありますが、clean architecture自体が難しい概念なので中々理解できない方が対象です。
かくゆう私も わからなくなることがたまによくあります！てへぺろ


# ソースコードはこちら
https://github.com/yabeyuji/macOnalO-chot-dekiru

# 動作サンプル
![demo](https://github.com/yabeyuji/macOnalO-chot-dekiru/blob/main/backend/doc/image/demo.gif?raw=true)

##### 左側:スタッフが確認できるモニター
| 項目     | 概要              |説明 |
| :-       |:-                |:- |
| reserve  | オーダー受付  | オーダーリクエストを受け付けて予約番号が発行されるとこちらに表示 |
| assemble | 調理中           | 調理可能な数(サンプルは3)に空きがあればreserveのオーダーがこちらに移動 |
| complete | オーダー完成      | 調理が完成したオーダーはyummyフォルダにjsonファイルとしておかれこちらに移動 |
| pass     | 引き渡し済み      | yummyフォルダにあるjsonファイルを削除するとこちらに移動 |


##### 右側:オーダーを注文するAPIなど

| 種類     | リクエスト方法             | リクエスト受付 | 補足 |
| :-       |:-                        |:-           |:-|
| mobile   | 1234port http:post:json  | echo  | |
| PC       | 2345port http:post:json  | gin   | |
| delivery | 3456port grpc            | grpc        |ウー○ー的な提携業者を想定 |
| register | jsonファイル指定場所に配置 | ファイル監視 | |


# ディレクトリ設計
全体の設計を<a href="https://github.com/golang-standards/project-layout/blob/master/README_ja.md">Standard Go Project Layout</a>
internal内のディレクトリ設計を<a href="https://blog.tai2.net/the_clean_architecture.html">clean Architecture</a>

#### ディレクトリ設計詳細
**全体**

| ディレクトリ | 説明 |
| --- | --- |
| build | Dockerfile |
| cmd | main.go |
| config | .bashrc |
| internal | 別テーブル参照 |
| pkg | 共通ライブラリ |
| script | entrypoint.shとオーダー用スクリプト |
| storage | ログ・エラーログ |
| web | index.htmlとvueファイルなど |
| yummy | オーダー完成場所(jsonファイル) |

**internal**

| ディレクトリ | 説明 |
| --- | --- |
| 1_framework | Frameworks & Drivers レイヤー |
| 1_framework/db | 食材の保管 |
| 1_framework/db/mysql | MySQL:冷凍庫 パティを保存 |
| 1_framework/db/postgres | PostgreSQL:冷蔵庫 野菜、チーズなどの保存 |
| 1_framework/db/mongo | MongoDB:棚 バンズを保存 |
| 1_framework/external_interface | オーダーモニター、商品の引き渡し |
| 1_framework/external_interface/monitor | websocket&vue:オーダーモニター |
| 1_framework/external_interface/shipment | jsonファイルを商品の引き渡しとする |
| 1_framework/web_ui | オーダー受け付け |
| 1_framework/web_ui/mobile | echo:モバイル注文 |
| 1_framework/web_ui/pc | gin:pc注文 |
| 1_framework/web_ui/delivery | grpc:ウー○ー的な提携業者からの注文 |
| 1_framework/web_ui/register | ファイル監視:レジ注文 |
| 2_adapter | Interface Adapters レイヤー |
| 2_adapter/controller | web_uiからのオーダーをusecaseに渡す |
| 2_adapter/gateway | usecaseからDBへのリクエストを渡す |
| 2_adapter/presenter | usecaseからexternal_interfaceへリクエストを渡す |
| 3_application_business_rule | Application Business Rules レイヤー |
| 3_application_business_rule/usecase | ユースケース |
| 4_enterprise_business_rule | Enterprise Business Rules レイヤー |
| 4_domain | ビジネスロジック 調理（野菜を切る、パティを焼くなど） |


# オーダーフロー
![orderflow](https://github.com/yabeyuji/macOnalO-chot-dekiru/blob/main/backend/doc/image/orderflow.png?raw=true)

# framework_driver/web_ui
[オーダーを受けつけるパート](https://github.com/yabeyuji/macOnalO-chot-dekiru/tree/main/backend/internal/1_framework/web_ui)

- web_uiのデータ型をControllerに持ち込まないようにentityのデータ型に変換
- オーダー番号発行
- オーダー ※オーダー番号を即時返却する必要があるため、Controllerにてチャネルを使用
- オーダー番号を返却

# framework_driver/db
[dbと接続するパート](https://github.com/yabeyuji/macOnalO-chot-dekiru/tree/main/backend/internal/1_framework/db)
- dbの違いはこのパートで解消
- 必要な食材を取得(デクリメント更新)

# framework_driver/external_interface
[db以外の外部と接続するパート](https://github.com/yabeyuji/macOnalO-chot-dekiru/tree/main/backend/internal/1_framework/external_interface)
- オーダー情報を随時更新するモニター(http://localhost:4567/)
- 商品の出荷(backend/yummyディレクトリ)
- 商品の出荷履歴(backend/storage/logディレクトリ)


# interface_adapter/controller
[web_uiからのオーダーを処理するパート](https://github.com/yabeyuji/macOnalO-chot-dekiru/tree/main/backend/internal/2_adapter/controller)
 - オーダー番号発行
 - オーダー処理

# interface_adapter/presenter
[usecaseからexternal_interfaceへ商品またはオーダー更新情報などを渡すパート](https://github.com/yabeyuji/macOnalO-chot-dekiru/tree/main/backend/internal/2_adapter/presenter)


# interface_adapter/gateway
[usecaseからdbへ必要な食材の情報を渡すパート](https://github.com/yabeyuji/macOnalO-chot-dekiru/tree/main/backend/internal/2_adapter/gateway)


# application_business_rule/usecase
[controllerからのオーダーを受け取り、目的に応じて以下に処理を渡すパート](https://github.com/yabeyuji/macOnalO-chot-dekiru/tree/main/backend/internal/3_usecase)

- オーダー内容の解析・調理(データの解析・変更)は enterprise_business_rule/entity
- 食材の取り出し(DB更新)は interface_adapter/gateway
- オーダー情報のモニタ表示・商品の出荷は interface_adapter/presenter

# enterprise_business_rule/entity
[ドメインロジックを扱うパート](https://github.com/yabeyuji/macOnalO-chot-dekiru/tree/main/backend/internal/4_domain)
- 調理に必要な食材数をカウントする
- 食材を調理する

# 動作確認
```
git clone git@github.com:yabeyuji/macOnalO-chot-dekiru.git

cd macOnalO-chot-dekiru
make build
make up

# オーダーは backend/script/order ディレクトリ内を参照
```

# 最後に
うーん。よく見るとハンバーガーのバンズ/レタス/パティ/トマト/バンズ もclean architectureが見えてきますね（タイトル回収）
疲れてるとclean architectureがゲシュタルト崩壊します。(初心者にありがち) 

### こんなときはハンバーガーを食べて寝るに限りますね。
### ちなみに私はサブウェイ派なのでBLTを食べます！
<a href="https://www.subway.co.jp/menu/sandwich/limited_sandwich/2580.html">メキシカンミートタコス</a> おいしそう


# 参考URL
https://qiita.com/t2-kob/items/02a76572693130c9a66e
https://qiita.com/nrslib/items/a5f902c4defc83bd46b8
https://qiita.com/koutalou/items/07a4f9cf51a2d13e4cdc
https://www.m3tech.blog/entry/2020/02/07/110000
https://www.lifull.blog/entry/2020/10/13/110000
