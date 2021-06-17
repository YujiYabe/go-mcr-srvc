# オーダー方法について

以下４つの方法でオーダー可能です。
- mobile
- pc
- delivery
- register


# mobile
mobile.restから任意のメニューをpostリクエストします。   
※rest Client拡張機能が必要です。

<br><br>

# pc
pc.restから任意のメニューをpostリクエストします。   
※rest Client拡張機能が必要です。

<br><br>

# delivery
delivery.shをcliから実行します。   
※処理実体はdelivery/delivery.goを参照

### grpcurl を使って調査
#### grpcインターフェイス確認コマンド   
grpcurl -plaintext localhost:3456 list

#### grpc送信コマンド(空オーダー)   
grpcurl -plaintext localhost:3456 delivery.DeliveryService/DeliveryRPC 

<br><br>
# register
registerディレクトリ内にjsonファイルを配置。   
order.exampleをコピーして拡張子をjsonに変更。   
<br><br>


