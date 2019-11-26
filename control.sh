#!/bin/bash

## ----------------------------
# ユーザからのキーボードの入力を受け取り、
# yes と入力されたらスクリプトを実行する、no と入力されたらスクリプトを終了します.
## ----------------------------

function ConfirmExecution() {

  echo "┌──────────────────────────────"
  echo "│"
  echo "│"
  echo "│"
  echo "│"
  echo "スクリプトを実行しますか?"
  echo "  実行する場合は yes、実行をキャンセルする場合は no と入力して下さい."
  echo "└──────────────────────────────"

  read input

  if [ -z $input ] ; then

    echo "  yes または no を入力して下さい."
    ConfirmExecution

  elif [ $input = 'yes' ] || [ $input = 'YES' ] || [ $input = 'y' ] ; then

    echo "  スクリプトを実行します."

  elif [ $input = 'no' ] || [ $input = 'NO' ] || [ $input = 'n' ] ; then

    echo "  スクリプトを終了します."
    exit 1

  else

    echo "  yes または no を入力して下さい."
    ConfirmExecution

  fi

}

# シェルスクリプトの実行を継続するか確認します。
ConfirmExecution

echo "----------------------------"
echo "hello world!"
