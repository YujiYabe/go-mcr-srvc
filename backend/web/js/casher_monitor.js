// グローバル変数
let products;
let soldNo = "1"; //初期値 local storage にない場合


async function init () {
  products = JSON.parse($("#product_list").text());
  await getSoldList();

  if (isDebug()) {
    $("#webcam").css("display", "block");
    $("#canvas").css("display", "block");
    $(".camera_zone").css("display", "block");
  }

}

// local storage にあれば使う
if (localStorage.hasOwnProperty("sold_no")) {
  soldNo = localStorage.getItem("sold_no");
}

async function setSoldNoLocalStorage (soldNo) {
  await localStorage.setItem("sold_no", soldNo);
}




async function toDisplayWithZeroFill (number) {
  $("#sold_no").val(zeroFill(number));
  await setSoldNoLocalStorage(number);
}


function appendToDisplay (value) {
  soldNo = soldNo + value;
  toDisplayWithZeroFill(soldNo);
}

function clearDisplay () {
  soldNo = "0";
  toDisplayWithZeroFill(soldNo);
}

function changeSoldNo (element) {
  toDisplayWithZeroFill(element.innerText);
}

async function buildOnWs (eventData) {

  data = JSON.parse(JSON.parse(eventData));
// console.log(" ==data['jan_code_list']======================= ");
// console.log( data['jan_code_list'] );
// console.log(" ========================= ");


  targetProductList = await convertJANtoProduct(data['jan_code_list']);


  await buildCasherMonitor(targetProductList);
}

async function buildCasherMonitor (targetProductList) {
  if (targetProductList.length == 0) return;

  let historyListElement = $("#history_list");
  let queueContentElement = $("#queue_content");
  let totalPrice = 0;
  let totalCount = 0;

  let divParent = $("<div>")
    .addClass("h-100 overflow-y-scroll")
    ;

  janCodes = [];

  $.each(targetProductList, function (_, targetProduct) {
    if (targetProduct == null || targetProduct == undefined) return;
// console.log(" ==targetProduct======================= ");
// console.log( targetProduct );
// console.log(" ========================= ");

    const janCode = targetProduct.jan_code;

    const price =
      targetProduct.retail_price +
      targetProduct.operation_cost +
      targetProduct.energy_cost
      ;

    totalPrice += (price * targetProduct.count);
    totalCount += targetProduct.count;


    for (let i = 0; i < targetProduct.count; i++) {
      janCodes.push(janCode);
    }


    let divCard = $("<div>")
      .appendTo(divParent)
      .addClass(" mb-3")
      ;

    let divRow = $("<div>")
      .appendTo(divCard)
      .addClass("row g-0")
      ;

    // =============================================
    let divCardLeft = $("<div>")
      .appendTo(divRow)
      .addClass("col-md-4")
      ;

    $("<img>")
      .appendTo(divCardLeft)
      .addClass("img-fluid rounded-start img-thumbnail ")
      .attr("src", "../image/product/" + janCode + ".jpg"
      )
      ;


    // =============================================
    let divCardRight = $("<div>")
      .appendTo(divRow)
      .addClass("col-md-8")
      ;

    let divCardBody = $("<div>")
      .appendTo(divCardRight)
      .addClass("card-body")
      ;

    let ul = $("<ul>")
      .appendTo(divCardBody)
      .addClass("list-group list-group-flush")
      ;

    let liPrice = $("<li>")
      .appendTo(ul)
      .addClass("h1 list-group-item")
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
      .text(" ＝ " + price * targetProduct.count)
      ;

    $("<li>")
      .appendTo(ul)
      .addClass("h6 list-group-item")
      .text(targetProduct.name_ja)
      ;
  });

  queueContentElement.empty();
  divParent.appendTo(queueContentElement);

  $("#predict_sold_no").text(zeroFill(soldNo));
  $("#previous_sold_no").text(zeroFill(soldNo - 1));
  toDisplayWithZeroFill(soldNo);

  $("#total_price").empty();

  discountPrice = 0;
  if (totalCount > 1) {
    discountPrice = 100 * (totalCount - 1)
  }

  $("#total_price").text(
    "計 " + totalCount + " 点 合計金額 " + (totalPrice - discountPrice) + " 円"
  );

  $("#casher_button")
    .addClass("border-danger border-3")
    ;

  historyListElement
    .css("display", "none")
    ;
}

async function getSoldList () {
  let soldList = [];
  let historyListElement = $("#history_list")

  soldList = JSON.parse($("#preparing_list").text());

  if (soldList.length == 0) return;

  let divParent = $("<div>")
    .addClass("h-100 overflow-y-scroll")
    ;


  for (let sold of soldList) {
    if (sold.status != "1") return;

    const divRow = $("<div>")
      .appendTo(divParent)
      .addClass("card p-8 ")
      ;

    const divCard = $("<div>")
      .appendTo(divRow)
      .addClass("row")
      ;

    $("<div>")
      .appendTo(divCard)
      .addClass("h2 col-md-4 card-body align-items-center")
      .text("注文番号：" + zeroFill(sold.sold_no))
      ;

    const divCardRight = $("<div>")
      .appendTo(divCard)
      .addClass("col-md-8 card-body")
      ;

    let targetProductList = await convertJANtoProduct(sold.jan_code_list);
    let ul = $("<ul>")
      .appendTo(divCardRight)
      .addClass("list-group list-group-flush")
      ;

    for (let targetProduct of targetProductList) {
      if (targetProduct == null || targetProduct == undefined) continue;

      $("<li>")
        .appendTo(ul)
        .addClass("list-group-item text-start")
        .text(targetProduct.count + " Ｘ " + targetProduct.name_ja + " ")
        ;
    }


    divCard.click(function () {
      showDeleteConfirmModal(
        () => deleteSoldItem(sold.sold_no),
        sold.sold_no
       );
    });
  }

  historyListElement.empty();
  divParent.appendTo(historyListElement);
}

async function showDeleteConfirmModal (callback, soldNo) {
  const modal = new bootstrap.Modal(document.getElementById("modal"));

  $("#modal_header_p")
    .addClass("text-center")
    .text("注文番号：" + zeroFill(soldNo))
    ;

  $("#modal_body_p1").text("削除してもよろしいですか？");

  modal.show();
  $("#confirm_button").off("click");
  $("#confirm_button").on("click", async () => {
    await callback();
  });
}

async function postCasher () {
  try {
    const soldNoString = $("#sold_no").val();

    const res = await fetch(`/v1/casher`, {
      method: "POST",
      body: JSON.stringify({
        sold_no: Number(soldNoString),
        status: preparingStatus,
        jan_code_list: janCodes,
      }),
      headers: {
        "Content-type": "application/json; charset=UTF-8",
      },
    });

    if (!res.ok) {
      throw new Error(res.status);
    }

    soldNo = String(Number(soldNoString) + 1);
    await setSoldNoLocalStorage(soldNo);

    // location.reload();
  } catch (error) {
    showErrorModal(error);
    console.error("Error :", error);
  }
}
async function changePage (element) {
  elName = element.getAttribute("name");

  let historyListElement = $("#history_list");
  let casherListElement = $("#casher_list");

  let historyButtonElement = $("#casher_button");
  let casherButtonElement = $("#history_button");

  let activeClass = "border-danger border-3";
  let tabStatus = elName == "casher_list" ? true : false

  tabStatus ? historyButtonElement.addClass(activeClass) : historyButtonElement.removeClass(activeClass);
  tabStatus ? casherButtonElement.removeClass(activeClass) : casherButtonElement.addClass(activeClass);
  tabStatus ? casherListElement.css("display", "") : casherListElement.css("display", "none");
  tabStatus ? historyListElement.css("display", "none") : historyListElement.css("display", "");
}

async function deleteSoldItem (soldNo) {
  try {
    await fetch(`/v1/casher/${soldNo}`, {
      method: "DELETE",
      body: JSON.stringify({}),
      headers: {
        "Content-type": "application/json; charset=UTF-8",
      },
    });
    location.reload();
  } catch (error) {
    console.error("Error :", error);
    // TODO エラーダイアログ出して 再試行させる
    // location.reload();
  }
}

async function showCompeteConfirmModal (callback) {
  const modal = new bootstrap.Modal(document.getElementById("modal"));

  $("#modal_body_p1").text("確定してもよろしいですか？");

  modal.show();
  $("#confirm_button").off("click");
  $("#confirm_button").on("click", async () => {
    await callback();
  });
}

window.addEventListener("load", init);
