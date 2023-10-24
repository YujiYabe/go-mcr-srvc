let allProducts;
let allergies;

const titleIsValid = "有効";
const titlePlace = "場所";
const titleJanCode = "JANコード";
const titleRecipe = "レシピ";
const titleCalorie = "カロリー";

const titleName = "名前";

const titleInformation = "商品情報";

const titleCostPrice = "原価";
const titleRetailPrice = "小売価格";
const titleOperationCost = "作業コスト";
const titleEnergyCost = "光熱費";

const titleAllergy = "アレルギー";

async function init () {
  allProducts = JSON.parse($("#all_product_list").text());
  allergies = JSON.parse($("#allergy_list").text());

  await buildAdminMonitor();
}


async function buildAdminMonitor () {
  try {
    $.each(allProducts, function (_, product) {
      const janCode = product.jan_code;
      printDebug("janCode", janCode); // -----------------------------

      let trRow = $("<tr>")
        .appendTo("#admin_monitor")
        ;

      let content = $("<td>")
        .appendTo(trRow)
        ;

      let divRow = $("<div>")
        .appendTo(content)
        .addClass("alert-info col-12 alert border-0 row ")
        ;

      let divImg = $("<div>")
        .appendTo(divRow)
        .addClass("col-2 alert alert-light border-0")
        ;

      $("<img>")
        .appendTo(divImg)
        .addClass("img-thumbnail")
        .attr("src", "../../image/product/" + janCode + ".jpg")
        ;

      let divContent = $("<div>")
        .appendTo(divRow)
        .addClass("col-10 alert alert-light border-0 row ")
        .attr('name', janCode)
        ;



      //==================
      // base info
      let divBaseInfo = $("<div>").addClass("col-12 row").appendTo(divContent);

      // -----------------
      // 有効  場所  JANコード レシピ  カロリー
      let divIsValid = $("<div>").addClass("col-2").appendTo(divBaseInfo);
      let divPlace = $("<div>").addClass("col-3").appendTo(divBaseInfo);
      let divJanCode = $("<div>").addClass("col-3").appendTo(divBaseInfo);
      let divRecipe = $("<div>").addClass("col-2").appendTo(divBaseInfo);
      let divCalorie = $("<div>").addClass("col-2").appendTo(divBaseInfo);
      // -----------------
      // 名前 
      let divName = $("<div>").addClass("col-12").appendTo(divBaseInfo);

      // -----------------
      // 商品情報
      let divInfo = $("<div>").addClass("col-12").appendTo(divBaseInfo);

      //==================
      // -----------------
      // 小売価格   原価  作業コスト  光熱費
      let divCost = $("<div>").addClass("col-12 row").appendTo(divContent);

      let divCostPrice = $("<div>").addClass("col-3").appendTo(divCost);
      let divRetailPrice = $("<div>").addClass("col-3").appendTo(divCost);
      let divOperationCost = $("<div>").addClass("col-3").appendTo(divCost);
      let divEnergyCost = $("<div>").addClass("col-3").appendTo(divCost);

      //==================
      // -----------------
      let divOther = $("<div>").addClass("col-12 row").appendTo(divContent);
      let divAllergy = $("<div>").addClass("col-12").appendTo(divOther);



      makeCommonDiv(titleName, product.name_ja, divName);
      makeCommonDiv(titleJanCode, janCode, divJanCode);
      makeCommonDiv(titleCalorie, product.calorie, divCalorie);
      makeCommonDiv(titleInformation, product.information_ja, divInfo);
      makeCommonDiv(titleRecipe, product.recipe, divRecipe);
      makeCommonDiv(titleCostPrice, product.cost_price, divCostPrice);
      makeCommonDiv(titleRetailPrice, product.retail_price, divRetailPrice);
      makeCommonDiv(titleOperationCost, product.operation_cost, divOperationCost);
      makeCommonDiv(titleEnergyCost, product.energy_cost, divEnergyCost);

      { // 有効
        $("<span>")
          .appendTo(divIsValid)
          .text(titleIsValid)
          .addClass("col badge bg-success")
          ;

        $("<span>")
          .appendTo(divIsValid)
          .text(" ")
          ;

        let divForm = $("<span>")
          .appendTo(divIsValid)
          .addClass("checkbox")
          ;

        let checkbox = $("<input>")
          .appendTo(divForm)
          .prop({
            type: 'checkbox',
            checked: product.is_valid
          })
          ;

        checkbox.change(function () {
          patchAdminProduct(janCode)
        });

        $("<span>")
          .appendTo(divIsValid)
          .text(" ")
          ;

        $("<span>")
          .appendTo(divIsValid)
          .text(product.is_valid.toString())
          .attr("id", "valid_" + janCode)
          ;
      }

      { // 場所
        $("<span>")
          .appendTo(divPlace)
          .addClass("col badge bg-success")
          .text(titlePlace)
          ;

        $("<span>")
          .appendTo(divPlace)
          .text(" ")
          ;

        let textBox = $("<span>")
          .appendTo(divPlace)
          .addClass("text-left col-9")
          ;

        $("<input>")
          .appendTo(textBox)
          // .addClass("form-control")
          .attr('name', 'place')
          .val(product.place)
          .prop({
            type: 'text',
          })
          ;

        textBox.change(function () {
          patchAdminProduct(janCode)
        });
      }

      { // アレルギー
        $("<span>")
          .appendTo(divAllergy)
          .addClass("col-1 badge bg-success")
          .text(titleAllergy)
          ;

        let divAllergyBody = $("<div>")
          .appendTo(divAllergy)
          .addClass("col-12 row")
          ;

        let cardTop = $("<div>")
          .appendTo(divAllergyBody)
          .addClass("col-12 row ms-0")
          ;

        const allergyList = JSON.parse(product.allergy);
        printDebug("allergyList", allergyList); // -----------------------------

        $.each(allergyList, function (key, boolean) {
          if (boolean == false) return;
          printDebug("allergies[key]", allergies[key]); // -----------------------------

          let card = $("<div>")
            .appendTo(cardTop)
            .addClass("card col-1 ")
            ;

          $("<img>")
            .appendTo(card)
            .addClass("img-thumbnail card-img-top")
            .attr("src", "../../image/allergy/" + key + ".gif")
            ;

          let cardBody = $('<div>')
            .appendTo(card)
            .addClass("text-center")
            ;

          $('<p>')
            .appendTo(cardBody)
            .addClass("text-center")
            .text(allergies[key])
            ;

        });
      }
    });

    $('#my_table').exTableFilter('#my_filter');

  } catch (error) {
    console.error('Error :', error);
  }
}

function makeCommonDiv (name, value, addTarget) {
  let div = $("<div>")
    .appendTo(addTarget)
    .addClass("col-12 list-group-item")
    ;


  $("<span>")
    .appendTo(div)
    .addClass("col badge bg-success")
    .text(name)
    ;

  $("<span>")
    .appendTo(div)
    .addClass("col")
    .text(" " + value)
    ;

}

async function patchAdminProduct (janCode) {
  try {
    let isValid = $(`div[name=${janCode}] input[type='checkbox']`).prop("checked");
    place = $(`div[name=${janCode}] input[name='place']`).val();

    $(`span[id=valid_${janCode}]`).text(isValid.toString());

    await fetch(`/admin/product/${janCode}`, {
      method: "PATCH",
      body: JSON.stringify({
        is_valid: isValid,
        place: place,
      }),
      headers: {
        'Content-type': 'application/json; charset=UTF-8',
      },
    });

  } catch (error) {
    console.error('Error :', error);
  }

}



// ページロード時の初期化
window.addEventListener('load', init);

