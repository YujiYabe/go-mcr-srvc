
// base color 

// "#beb97f"; // rgb(190,185,127) 淡緑 reserving
// "#f1af63"; // rgb(241,175,99)  黄   preparing
// "#e98b67"; // rgb(233,139,103) 赤   completed
// "#8f9c95"; // rgb(143,156,149) 深緑 passed

const colorPreparing = "#f1af63"; // rgb(241,175,99)
const colorCompleted = "#e98b67"; // rgb(233,139,103)
const colorPassed = "#8f9c95"; // rgb(143,156,149)


const preparingStatus = 1;
const completedStatus = 2;
const passedStatus = 3;

const preparingText = "準備中";
const completedText = "準備完了";
const passedText = "お渡し済み";


function parseSold (soldList) {
  let preparingList = [];
  let completedList = [];
  let passedList = [];

  soldList.forEach(value => {
    switch (value.status) {
      case preparingStatus:
        preparingList.push(value);
        break;
      case completedStatus:
        completedList.push(value);
        break;
      case passedStatus:
        passedList.push(value);
        break;
      default:
        console.error("error status : ", value);
        break;
    }
  });

  return [preparingList, completedList, passedList];
}

function zeroFill (number) {
  return ("000" + number).slice(-3);
}

async function getRequestKitchenQueue (responseData) {
  try {
    let targetProductList = [];
    let soldNo = 0;

    if (responseData.hasOwnProperty("status")) {
      soldNo = responseData.sold_no;
    }

    targetProductList = await convertJANtoProduct(
      responseData.jan_code_list
    );



    targetProductList.sort((a, b) => {
      if (a.place < b.place) return -1;
      else if (a.place > b.place) return 1;
      return 0;
    });

    return [targetProductList, soldNo];
  } catch (error) {
    console.error("Error :", error);
  }
}

async function countJANcodes (janCodes) {
  let janCodeMapCount = new Map();
  $.each(janCodes, function (_, janCode) {
    janCodeMapCount.set(janCode, (janCodeMapCount.get(janCode) || 0) + 1);
  });

  return janCodeMapCount;
}

async function convertJANtoProduct (janCodes) {
  let targetProductList = [];

  if (typeof products === "undefined") {
    return targetProductList;
  }

  const janCodeMapCount = await countJANcodes(janCodes);

  for (let [janCode, count] of janCodeMapCount) {
    let targetProduct;
    $.each(products, function (_, product) {
      if (product.jan_code === janCode) {
        targetProduct = product;
        targetProduct.count = count;
        return false;
      }
    });

    targetProductList.push(targetProduct);
  }

  sortProductByPrice(targetProductList);

  return targetProductList;
}

async function sortProductByPrice (products) {
  products.sort((a, b) => {
    let totalPrice1 = a.retailPrice + a.operationCost + a.energyCost;
    let totalPrice2 = b.retailPrice + b.operationCost + b.energyCost;
    if (totalPrice1 < totalPrice2) return 1;
    else if (totalPrice1 > totalPrice2) return -1;
    return 0;
  });
}

function scrollToBottom (timeDelay = 0) {
  var scrollId;
  var height = 0;
  var minScrollHeight = 100;
  scrollId = setInterval(function () {
    if (height <= document.body.scrollHeight) {
      window.scrollBy(0, minScrollHeight);
    } else {
      clearInterval(scrollId);
    }
    height += minScrollHeight;
  }, timeDelay);
}

// 自動垂直スクロールを設定する関数
function autoVerticalScroll () {
  swiper =
    janCodes.length <= 2
      ? new Swiper(".swiper-container", {
        direction: "vertical",
        height: 350,
      })
      : new Swiper(".swiper-container", {
        direction: "vertical",
        slidesPerGroup: 2,
        height: 350,
        speed: 7000,
        autoplay: {
          delay: 1000,
        },
        on: {
          slideChange: function (swiper) {
            // スライドが最後まで進んだら自動再生を停止&&先頭に戻す
            swiper.isEnd ? (swiper.autoplay.stop(), swiper.slideTo(0)) : null;
          },
        },
      }
      );
}

// SHA-256ハッシュを計算する関数
async function calculateSHA256Hash (data) {
  const encoder = new TextEncoder();
  const dataBuffer = encoder.encode(data);
  const hashBuffer = await crypto.subtle.digest("SHA-256", dataBuffer);
  const hashArray = Array.from(new Uint8Array(hashBuffer));
  const hashHex = hashArray
    .map((byte) => byte.toString(16).padStart(2, "0"))
    .join("");
  return hashHex;
}

async function showErrorModal (error) {
  const modal = new bootstrap.Modal(document.getElementById("modal"));
  $("#modal_header_p").text("送信できませんでした。もう一度お願いします");
  $("#modal_body_p1").text(error);
  $("#modal_body_p2").text("");
  $("#modal_body_p3").text("");
  $("#modal_body_p4").text("");
  $("#confirm_button").off("click");
  modal.show();
}



// https://www.relief.jp/docs/001403.html

async function parseAllergyLang (selectedLanguageCode, allergies) {
  let selectedAllergy

  $.each(allergies, function (_, allergy) {
    if (String(selectedLanguageCode) == String(allergy.lang_type)) {
      selectedAllergy = allergy;
      return
    };
  });

  return selectedAllergy
}

function printDebug (title, value) {
  if (isDebug()) {
    console.log(" =" + title + " ======================== ");
    console.log(value);
  }
}



function isDebug () {
  if ('?d' == window.location.search) {
    return true
  }

  return false
}
