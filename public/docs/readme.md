# はじめに
<a href="../../readme.md">こちら</a>のブラウザ側の解説です。
<br><br><br>

# 概要
### 以下の機能を提供
- ウェブサーバとwebSocketで接続してリアルタイム通信
- vueLoaderでコンパイルなしでvueコンポーネントを提供
<br><br>


# ファイル構成
***vue.jsなどの外部js,cssについてオフラインでも使えるようにcdnではなくlocalから取得するようにしています***

| No | ファイル名         | 概要 |
| -- | ------------------ | ---- |
| 1  | index.html         | トップページ | 
| 2  | js/index.js        | メインjs | 
| 3  | home.vue           | ルートコンポーネント。4のXXX.vueを切り替え | 
| 4  | vue/room/XXX.vue   | 部屋コンポーネント。5のYYY.vueを切り替え| 
| 5  | vue/device/YYY.vue | 家電コンポーネント。 | 
| 6  | vue/icon/ZZZ.vue   | アイコンコンポーネント。3,4,5で使われる | 
<br><br>


# 赤外線情報
vue/device/YYY.vue内に定義してあるdata-valueがデバイスサーバの赤外線のマップ※のキーとなっている。
※backend/device/internal/1_infrastructure/deviceapp/ir_data.go
<br><br>

# 部屋情報
ブラウザ側でlocal storageに各部屋の選択した機器を保存し、部屋の切り替えの際に元の機器を選択した状態にする、
<br><br>


# リンク
## セットアップ
<a href="docs/setup.md">こちら参照</a>
<br><br>

## 共通設計
<a href="docs/common_structure.md">こちら参照</a>
<br><br>

## サーバ
以下参照   
<a href="backend/ws/docs/readme.md">ウェブサーバ</a>   
<a href="backend/file/docs/readme.md">ファイルサーバ</a>   
<a href="backend/media/docs/readme.md">メディアサーバ</a>   
<a href="backend/device/docs/readme.md">デバイスサーバ</a>   
<br><br>

