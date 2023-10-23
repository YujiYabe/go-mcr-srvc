let product;

async function init () {
  product = JSON.parse($("#product").text());
  await buildAdminPrintKitchen();
}



async function buildAdminPrintKitchen () {

  const printZone = $('#print_zone');

  for (let i = 0; i < 12; i++) {
    let divNew = $("<div>")
      .appendTo(printZone)
      .addClass("card col-3 width_18rem")
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


    let cardBody = $("<div>")
      .appendTo(divNew)
      .addClass("card-body")
      ;

    $("<div>")
      .appendTo(cardBody)
      .addClass("card-title")
      .text(product.name_ja)
      ;

    $("<p>")
      .appendTo(cardBody)
      .addClass("card-text")
      .text(product.recipe)
      ;

  }

  attachImage();
}

function attachImage () {
  $("img[name=image]")
    .attr("src", "../../../image/product/" + product.jan_code + ".jpg")
    ;
}

window.addEventListener('load', init);

