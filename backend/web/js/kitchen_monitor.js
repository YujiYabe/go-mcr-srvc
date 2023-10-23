// グローバル変数
let products;

async function init () {
  products = JSON.parse($('#product_list').text());
}

async function buildOnWs (eventData) {
  buildKitchenMonitor(JSON.parse(JSON.parse(eventData)));
}

function createDivWithClass (className) {
  return $("<div>").addClass(className);
}


async function buildKitchenMonitor (eventData) {
  [targetProductList, soldNo] = await getRequestKitchenQueue(eventData);
  printDebug("targetProductList", targetProductList); // -----------------------------

  let queueContent = $('#queue_content');
  if (targetProductList.length == 0) {
    queueContent.empty();
    return;
  }

  let divOrderHead = createDivWithClass("h2 border-0");

  if (soldNo != 0) {
    const soldNoString = zeroFill(soldNo);
    divOrderHead.text("注文番号 : " + soldNoString).addClass("alert alert-info text-center");
    divOrderHead.addClass(" ");
  } else {
    divOrderHead.addClass(" ");
  }

  const divParent = createDivWithClass("row");

  targetProductList.forEach(targetProduct => {
    if (targetProduct == null || targetProduct == undefined) {
      return;
    }

    const divProduct = $("<div>")
      .appendTo(divParent)
      .addClass("col-4 row border border-black-50 border-1 margin-0")
      ;

    // 商品カウント
    $("<div>")
      .appendTo(divProduct)
      .addClass("col-5 text-center badge bg-light text-dark font-150px")
      .text(targetProduct.count)
      ;

    // 商品画像
    const divImage = $("<div>")
      .appendTo(divProduct)
      .addClass("col-7")
      ;
    $("<img>")
      .appendTo(divImage)
      .addClass("img-fluid rounded-start")
      .attr("src", "../../../image/product/" + targetProduct["jan_code"] + ".jpg")
      ;

    // 商品場所
    $("<div>")
      .appendTo(divProduct)
      .addClass("col-6 h1 text-center font-70px")
      .text(targetProduct.place)
      ;

    // レシピ
    $("<div>")
      .appendTo(divProduct)
      .addClass("col-6 h3 text-center font-70px")
      .text(targetProduct.recipe)
      ;

  });

  queueContent.empty();

  divOrderHead.appendTo(queueContent);
  divParent.appendTo(queueContent);
}

window.addEventListener('load', init);
