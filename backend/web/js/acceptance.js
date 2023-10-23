// 初期化関数
async function init () {
}

async function buildOnWs (eventData) {
  buildAcceptance(eventData);
}


async function buildAcceptance (eventData) {
  soldList = JSON.parse(JSON.parse(eventData));

  if (soldList.length == 0) return;

  try {
    let preparingList = [];
    let completedList = [];
    let passedList = [];
    let preparingListParent = $('#preparing_list');
    let completedListParent = $('#completed_list');

    [preparingList, completedList, passedList] = parseSold(soldList);


    let divParentPreparing = $("<div>");
    let divParentCompleted = $("<div>");

    $.each(preparingList, function (_, value) {
      $("<div>")
        .appendTo(divParentPreparing)
        .addClass("h3 alert alert-light border-0")
        .text(zeroFill(value['sold_no']))
        ;
    });

    $.each(completedList, function (_, value) {
      $("<div>")
        .appendTo(divParentCompleted)
        .addClass("h3 alert alert-light border-0")
        .text(zeroFill(value['sold_no']))
        ;
    });

    preparingListParent.empty();
    completedListParent.empty();

    divParentPreparing.appendTo(preparingListParent);
    divParentCompleted.appendTo(completedListParent);

  } catch (error) {
    console.error('Error :', error);
  }
}

window.addEventListener('load', init);


