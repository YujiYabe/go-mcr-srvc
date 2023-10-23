async function init () {
  buildAdminClient();
}

function buildAdminClient () {

  const janCode = $('#jan_code').attr('name');

  const printZone = $('#print_zone');
  const commonSize = 90;

  for (let i = 0; i < 12; i++) {
    let divNew = $("<div>")
      .addClass("card col-3 width_18rem")
      .appendTo(printZone)
      ;

    $("<div>")
      .appendTo(divNew)
      .addClass("d-block mx-auto color-secondary pt-2")
      .text("○")
      ;

    $("<img>")
      .appendTo(divNew)
      .addClass("product-img d-block mx-auto")
      .attr("name", "image")
      ;

    $("<div>")
      .appendTo(divNew)
      .addClass("product-img d-block mx-auto")
      .attr("name", "qr")
      ;


    $("<img>")
      .appendTo(divNew)
      .addClass("bar-code product-img d-block mx-auto ")
      .attr("name", "barcode")
      ;

  }

  attachImage(janCode, commonSize);
  attachQRCode(janCode, commonSize);
  attachBarCode(janCode, commonSize);
}

function attachImage (janCode, commonSize) {
  $("img[name=image]")
    .attr("src", "../../../image/product/" + janCode + ".jpg")
    ;
}


function attachQRCode (janCode, commonSize) {
  $("div[name=qr]")
    .html("")
    .qrcode({ width: commonSize, height: commonSize, text: janCode })
    ;
}


function attachBarCode (janCode, commonSize) {
  JsBarcode("img[name=barcode]", janCode, { format: "ean13" });
}


window.addEventListener('load', init);

