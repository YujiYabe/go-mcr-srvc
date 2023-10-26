let soldList;

async function init () {
  soldList = JSON.parse($("#delivery").text());
  await buildDeliveryMonitor();
}

async function buildDeliveryMonitor () {
  try {
    $.each(soldList, async function (_, sold) {

      let rowColor = "";
      let textStatusPreparing = "";
      let textStatusCompleted = "";
      let textSoldNoPreparing = "";
      let textSoldNoCompleted = "";
      let textSoldNoPassed = "";

      let preparingBorder = "border-1";
      let completedBorder = "border-1";
      let soldNo = "" + ("000" + sold["sold_no"]).slice(-3);



      switch (sold["status"]) {
        case preparingStatus:
          rowColor = " alert-info ";
          textStatusPreparing = preparingText;
          textSoldNoPreparing = soldNo;
          completedBorder = " border-0 ";
          break;

        case completedStatus:
          rowColor = " alert-primary ";
          textStatusCompleted = completedText;
          textSoldNoCompleted = soldNo;
          preparingBorder = " border-0 ";
          break;

        case passedStatus:
          rowColor = " alert-dark ";
          textStatusCompleted = passedText;
          textSoldNoPassed = soldNo;
          PassedBorder = " border-0 ";
          break;

        default:
          console.error("error status : ", sold);
          break;
      }

      if (sold["status"] != passedStatus) {
        {
          let list = $("#sold_list");
          let rowDiv = $("<div>")
            .appendTo(list)
            .addClass(" col-12 row  center-block border-0 " + rowColor)
            ;

          // TODO 見えないonclickボタンの削除
          // 準備中
          let DivPrepare = await createButton(rowColor, preparingBorder, textStatusPreparing, textSoldNoPreparing, sold, soldNo);
          DivPrepare.appendTo(rowDiv);

          // 準備完了
          let DivComplete = await createButton(rowColor, completedBorder, textStatusCompleted, textSoldNoCompleted, sold, soldNo);
          DivComplete.appendTo(rowDiv);

          rowDiv.append(await makeImage(sold["jan_code_list"].sort(), rowColor));

        }
      }
      if (sold["status"] == passedStatus) {
        // お渡し済み
        let list = $("#passed_list");
        let rowDiv = $("<div>")
          .appendTo(list)
          .addClass("col-12 row center-block border-0  " + rowColor)
          ;

        // 準備中
        let DivPrepare = await createButton(rowColor, preparingBorder, passedText, textSoldNoPassed, sold, soldNo);
        DivPrepare.appendTo(rowDiv);


        rowDiv.append(await makeImage(sold["jan_code_list"].sort(), rowColor));
      }
    });
    $("#sold_button")
      .addClass("border-danger border-3")
      ;
  } catch (error) {
    console.error("Error :", error);
  }
}

async function makeImage (janCodeList, rowColor) {
  let divContent = $("<div>")
    .addClass("col border-0 row ps-4 alert ms-0" + rowColor)
    ;

  targetProductList = await countJANcodes(janCodeList);

  for await (let [janCode, count] of targetProductList) {
    let divRow = $("<div>")
      .appendTo(divContent)
      .addClass("col-4 row mb-3 ")
      ;


    let divImg = $("<div>")
      .appendTo(divRow)
      .addClass(" border-0 ms-0 position-relative")
      ;

    $("<div>")
      .appendTo(divImg)
      .addClass(" border-0 badge bg-success align-items-center font-20px position-absolute top-0 start-0")
      .text(count)
      ;


    $("<img>")
      .appendTo(divImg)
      .addClass("img-thumbnail")
      .attr("src", "../../../../image/product/" + janCode + ".jpg")
      ;
  }

  return divContent;
}

async function patchSold (soldNo, status) {
  try {
    const res = await fetch(`/v1/delivery`, {
      method: "PATCH",
      body: JSON.stringify({
        sold_no: soldNo,
        status: status,
      }),
      headers: {
        "Content-type": "application/json; charset=UTF-8",
      },
    });
    if (!res.ok) {
      throw new Error(res.status);
    }
    location.reload();
  } catch (error) {res
    showErrorModal(error);
    console.error("Error :", error);
  }
}

async function showConfirmModal (callback, soldNo, status) {
  const modal = new bootstrap.Modal(document.getElementById("modal"));
  $("#modal_header_p")
    .addClass("text-center")
    .text(soldNo)
    ;

  let modalBodyP1 = $("#modal_body_p1");
  let modalBodyP3 = $("#modal_body_p3");

  $("#modal_body_p2").text("⇒")

  switch (status) {
    case preparingStatus:

      modalBodyP1
        .text(passedText)
        .css({ "background-color": "#ced4da" })

      modalBodyP3
        .text(preparingText)
        .css({ "background-color": "#cff4fc" })
      break;

    case completedStatus:
      modalBodyP1
        .text(preparingText)
        .css({ "background-color": "#cff4fc" })
      modalBodyP3
        .text(completedText)
        .css({ "background-color": "#cfe2ff" })
      break;

    case passedStatus:
      modalBodyP1
        .text(completedText)
        .css({ "background-color": "#cfe2ff" })
      modalBodyP3
        .text(passedText)
        .css({ "background-color": "#ced4da" })
      break;

    default:
      console.error("error status : ", sold);
      break;
  }



  modal.show();
  $("#confirm_button").off("click");
  $("#confirm_button").on("click", async () => await callback());
}

function changePage (element) {
  elName = element.getAttribute("name");

  let soldButtonElement = $("#sold_button");
  let passedButtonElement = $("#passed_button");

  let soldListElement = $("#sold_list");
  let passedListElement = $("#passed_list");

  let activeClass = "border-danger border-3";
  let tabStatus = elName == "sold_list" ? true : false

  tabStatus ? soldButtonElement.addClass(activeClass) : soldButtonElement.removeClass(activeClass);
  tabStatus ? passedButtonElement.removeClass(activeClass) : passedButtonElement.addClass(activeClass);
  tabStatus ? soldListElement.css("display", "") : soldListElement.css("display", "none");
  tabStatus ? passedListElement.css("display", "none") : passedListElement.css("display", "");

}


async function createButton (rowColor, border, textSoldNo, textStatus, sold, soldNo) {
  let newDiv = $("<div>")
    .addClass(
      "col-2 d-flex align-items-center border-success alert" +
      rowColor +
      border
    )
    ;



  if (sold["status"] == preparingStatus && textStatus != "") {
    newDiv.click(function () {
      showConfirmModal(
        async () => {
          await patchSold(sold["sold_no"], completedStatus);
        },
        soldNo,
        completedStatus
      );
    });
  }

  if (sold["status"] == completedStatus && textStatus != "") {
    newDiv.click(function () {
      showConfirmModal(
        async () => {
          await patchSold(sold["sold_no"], passedStatus);
        },
        soldNo,
        passedStatus
      );
    });
  }

  if (sold["status"] == passedStatus) {
    newDiv.click(function () {
      showConfirmModal(
        async () => {
          await patchSold(sold["sold_no"], preparingStatus);
        },
        soldNo,
        preparingStatus
      );
    });
  }

  let contentDiv = $("<div>")
    .appendTo(newDiv)
    .addClass("m-auto")
    ;

  $("<div>")
    .appendTo(contentDiv)
    .addClass("h1")
    .text(textStatus)
    ;

  $("<div>")
    .appendTo(contentDiv)
    .text(textSoldNo)
    ;

  return newDiv

}

window.addEventListener("load", init);
