async function init () {
  buildAdminPrintLanguage();
}

function buildAdminPrintLanguage () {

  const languageCode = $('#language_code').attr('name');
  const languageName = $('#language_name').attr('name');

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

    $("<div>")
      .appendTo(divNew)
      .addClass("d-block h1 mx-auto color-secondary pt-2")
      .text(languageName)
      ;


    $("<div>")
      .appendTo(divNew)
      .addClass("product-img d-block mx-auto")
      .attr("name", "qr")
      ;

    $("<br>")
      .appendTo(divNew)
      ;
    // $("<img>")
    //   .appendTo(divNew)
    //   .addClass("bar-code product-img d-block mx-auto ")
    //   .attr("name", "barcode")
    //   ;

  }



  attachQRCode(languageCode, commonSize);

}


function attachQRCode (languageCode, commonSize) {
  $("div[name=qr]")
    .html("")
    .qrcode({ width: commonSize, height: commonSize, text: languageCode })
    ;
}




window.addEventListener('load', init);

