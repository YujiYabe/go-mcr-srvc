#!/bin/bash

## ----------------------------
# ユーザからのキーボードの入力を受け取り、
# yes と入力されたらスクリプトを実行する、no と入力されたらスクリプトを終了します.
## ----------------------------

function ConfirmExecution() {

  echo "┌──────────────────────────────"
  echo "│ 1 reset"
  echo "│ 2 build"
  echo "│ 9 exit"
  echo "└──────────────────────────────"

  read input

  # ----------------------------------------------
  if [ -z $input ] ; then
    echo "not valid input"
    ConfirmExecution

  # ----------------------------------------------
  elif [ $input = '1' ] ; then
    docker-compose stop
    docker system prune

    sudo rm -rf db/engine/mysql/var_lib_mysql/
    sudo rm -rf db/engine/mysql/data/
    sudo rm -rf db/tool/phpmyadmin/sessions/
    sudo rm -rf go_src/go_api/gin-bin

  # ----------------------------------------------
  elif [ $input = '2' ] ; then
    docker-compose build --no-cache
    docker-compose up

  # ----------------------------------------------
  elif [ $input = '9' ] ; then
    echo "exit "
    exit 1

  # ----------------------------------------------
  else
    echo "not valid input"
    ConfirmExecution
  fi

ConfirmExecution

}

# シェルスクリプトの実行を継続するか確認します。
ConfirmExecution

