// グローバル変数
let janCodes = [];
// let swiper;
let previousResponseHash = "";
let products;
let allergies;

// 初期化関数
async function init () {
  products = JSON.parse($("#product_list").text());
  allergies = JSON.parse($("#allergy_list").text());

  
  if (isDebug()){
    $("#webcam").css("display", "block");
    $("#canvas").css("display", "block");
    $(".camera_zone").css("display", "block");
  }
}

async function buildOnWs (eventData) {
  buildClientMonitor(eventData);
}



// クライアントモニターを構築する関数
async function buildClientMonitor (eventData) {
  const divParent = $("<div>")
  // .addClass("swiper-wrapper")
  ;
  printDebug("eventData", eventData) ; // -----------------------------
  data = JSON.parse(JSON.parse(eventData));

  targetProductList = await convertJANtoProduct(data['jan_code_list']);
  printDebug("targetProductList", targetProductList) ; // -----------------------------

  if (targetProductList == null) {
    $("#queue_content").empty();
    return;
  }

  selectedLanguageCode = data['language_code'];
  let selectedAllergy = await parseAllergyLang(selectedLanguageCode, allergies);

  hashString = data['language_code'] + "-" + data['jan_code_list'].join('-');

  const currentResponseHash = await calculateSHA256Hash(hashString);

  if (previousResponseHash != currentResponseHash ) {
    await setHashContentLocalStorage (currentResponseHash);
    location.reload();
  };

  let totalPrice = 0;
  let totalCount = 0;

  let taxName;
  await targetProductList.forEach((targetProduct) => {
    if (!targetProduct) return;


    let targetProductName
    let targetProductInformation
    let calorieName
    let allergyName


    switch (selectedLanguageCode) {
      case 1041: // japanese
        targetProductName = targetProduct['name_ja'];
        targetProductInformation = targetProduct['information_ja'];
        calorieName = " カロリー";
        allergyName = "アレルギー情報";
        taxName = " 円(税込み)";
        // totalAmount = " 点 合計金額 ";
        break;
      case 1033: // us english
        targetProductName = targetProduct['name_en'];
        targetProductInformation = targetProduct['information_en'];
        calorieName = " calorie";
        allergyName = "allergy info";
        taxName = " yen(tax included)";
        // totalAmount = "total amount";
        break;
      case 1034: // spanish
        targetProductName = targetProduct['name_es'];
        targetProductInformation = targetProduct['information_es'];
        calorieName = " caloria";
        allergyName = "Información sobre alergias";
        taxName = " yen(impuesto incluido)";
        // totalAmount = "cantidad total";
        break;
      case 2052:// 中国語 (中華人民共和国)
        targetProductName = targetProduct['name_zh'];
        targetProductInformation = targetProduct['information_zh'];
        calorieName = " 卡路里";
        allergyName = "过敏信息";
        taxName = " yen(含税)";
        // totalAmount = "总金额";
        break;
      case 14337:// アラビア語 (U.A.E.)
        targetProductName = targetProduct['name_ar'];
        targetProductInformation = targetProduct['information_ar'];
        calorieName = " سعرات حرارية";
        allergyName = "معلومات الحساسية";
        taxName = " yen(شامل الضريبة)";
        // totalAmount = "المبلغ الإجمالي";
        break;
      default:
        console.error("error status : ", value);
        break;
    }

    const price =
      targetProduct.retail_price +
      targetProduct.operation_cost +
      targetProduct.energy_cost;

    totalPrice += (price * targetProduct.count);
    totalCount += targetProduct.count;

    janCodes.push(targetProduct.jan_code);

    const divCard = $("<div>")
      .appendTo(divParent)
      // .addClass("card swiper-slide p-8")
      .addClass("card p-8")
      ;

    const divRow = $("<div>")
      .appendTo(divCard)
      .addClass("row")
      ;

    const divCardLeft = $("<div>")
      .appendTo(divRow)
      .addClass("col-md-3")
      ;

    $("<img>")
      .appendTo(divCardLeft)
      .addClass("img-fluid rounded-start ")
      .attr("src", "../../../image/product/" + targetProduct.jan_code + ".jpg")
      ;


    const divCardRight = $("<div>")
      .appendTo(divRow)
      .addClass("col-md-9")
      ;

    const divCardBody = $("<div>")
      .appendTo(divCardRight)
      .addClass("card-body")
      ;


    // price
    const liPrice = $("<div>")
      .appendTo(divCardBody)
      .addClass("h1 card-title ")
      ;

    $("<span>")
      .appendTo(liPrice)
      .addClass("h1 card-text")
      .text(targetProduct.count + " Ｘ")
      ;

    $("<span>")
      .appendTo(liPrice)
      .addClass("h1 text-center")
      .text("  " + price + "")
      ;

    $("<span>")
      .appendTo(liPrice)
      .addClass("h1 card-text")
      .text(" ＝ " + price * targetProduct.count + taxName)
      ;

    // calorie
    const liCalorie = $("<div>")
      .appendTo(divCardBody)
      .addClass("h6 card-title ")
      ;

    $("<span>")
      .appendTo(liCalorie)
      .addClass("h6 card-text")
      .text(targetProduct.count + " Ｘ")
      ;

    $("<span>")
      .appendTo(liCalorie)
      .addClass("h6 text-center")
      .text("  " + targetProduct.calorie + " ＝ ")
      ;

    $("<span>")
      .appendTo(liCalorie)
      .addClass("h6 card-text")
      .text("  " + targetProduct.calorie * targetProduct.count + calorieName)
      ;

    // 名前
    $("<div>")
      .appendTo(divCardBody)
      .addClass("h6 card-text")
      .text(targetProductName)
      ;

    // 説明
    $("<div>")
      .appendTo(divCardBody)
      .addClass("h6 card-text")
      .text(targetProductInformation)
      ;

    $("<div>")
      .appendTo(divCardBody)
      .addClass("h6 card-text")
      .attr("id", "calorie_" + targetProduct.jan_code)
      ;

    const li_allergy = $("<div>")
      .appendTo(divCardBody)
      .addClass("h6 card-text")
      ;

    $("<div>")
      .appendTo(li_allergy)
      .addClass("badge bg-danger my-2")
      .text(allergyName)
      ;

    const cardTop = $("<div>")
      .appendTo(li_allergy)
      .addClass("row mx-auto col-12 ")
      ;

    const allergyList = JSON.parse(targetProduct.allergy);


    for (const key in allergyList) {
      if (allergyList[key]) {
        const card = $("<div>")
          .appendTo(cardTop)
          .addClass("card col-1 width_18rem")
          ;

        $("<img>")
          .appendTo(card)
          .addClass("img-thumbnail card-img-top")
          .attr("src", "../../../image/allergy/" + key + ".gif")
          ;

        const cardBody = $("<div>")
          .appendTo(card)
          .addClass("text-center")
          ;

        // $("<p>")
        //   .appendTo(cardBody)
        //   .addClass("text-center")
        //   .text(key)
        //   ;

        $("<p>")
          .appendTo(cardBody)
          .addClass("text-center")
          .text(selectedAllergy[key])
          ;
      }
    }
  });

  $("#queue_content").empty();
  divParent.appendTo("#queue_content");
  // $("#queue_content").addClass("swiper-container");
  $("#queue_content");

  $("#total_price").empty();

  if (totalPrice != 0) {
    discountPrice = 0 ;

    if (totalCount > 1) {
      discountPrice = 100 * (totalCount - 1)
      discountText = "(" + totalPrice + "-" + discountPrice +") " 
      $("<span>")
        .appendTo( $("#total_price"))
        .addClass("text-secondary")
        .text(discountText)
        ;
    }

    $("<span>")
      .appendTo( $("#total_price"))
      // .addClass("text-center")
      .text((totalPrice - discountPrice) + taxName)
      ;

      
    // $("#total_price").text(totalPrice + taxName + discountText);
  }


  
  // autoVerticalScroll();
}

// local storage にあれば使う
if (localStorage.hasOwnProperty("hash_content")) {
  previousResponseHash = localStorage.getItem("hash_content");
}

async function setHashContentLocalStorage (value) {
  await localStorage.setItem("hash_content", value);

}
// ページロード時の初期化
window.addEventListener("load", init);
