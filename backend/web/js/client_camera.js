async function init () {
  startWebcam();
  if (!isDebug()) { // debugモードでは手動キャプチャ
    captureInterval = setInterval(captureAndUpload, 5000); // n秒ごとにキャプチャーと送信
  }
}

const webcamElement = $('#webcam')[0];
const captureButton = $('#capture')[0];
let isExist = false

let captureInterval;

async function startWebcam () {
  try {
    // メディアデバイスからビデオストリームを取得
    const stream = await navigator.mediaDevices.getUserMedia({
      video: true // PCカメラを使用
      // video: {
      //     facingMode: { exact: "environment" }, // スマホの外部カメラを使用する場合はこの行を有効にする
      // },
    });

    // 取得したストリームをwebcamElementに設定
    $(webcamElement).prop('srcObject', stream);
  } catch (error) {
    // エラーメッセージを詳細にする
    console.error('ウェブカメラへのアクセス中にエラーが発生しました。詳細:', error);
  }
}


async function captureAndUpload () {
  const ratios = 2;
  const queueNo = $('#number').attr('name');

  const canvas = $('<canvas>')[0];
  canvas.width = webcamElement.videoWidth * ratios;
  canvas.height = webcamElement.videoHeight * ratios;
  let ctx = canvas.getContext("2d");
  ctx.drawImage(webcamElement, 0, 0, canvas.width, canvas.height);

  // let img = ctx.getImageData(0, 0, canvas.width, canvas.height);
  // let code = jsQR(img.data, img.width, img.height, { inversionAttempts: "dontInvert" });


  // TODO ブラウザでQRを一つも判別できない場合はAI処理をしない
  // cameraに何も映らない場合はqueueの内容を削除する
  const imageBlob = await new Promise(resolve => canvas.toBlob(resolve, 'image/jpeg', 1.0));

  const formData = new FormData();
  formData.append('image', imageBlob, 'image.jpg');

  try {
    await fetch(`/v1/client/camera/${queueNo}`, {
      method: "POST",
      body: formData,
    });
  } catch (error) {
    console.error('Error uploading image:', error);
  }

}


$(captureButton).on('click', () => {
  captureAndUpload()
});


window.addEventListener('load', init);
