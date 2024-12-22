# setup
### 下記インストール
- docker (version 20.10.5)
- docker-compose (version 1.24.0)
### 立ち上げ
```
make build
make up
```

# air
cd backend/   
curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b ./bin


# gosec
cd backend/   
curl -sfL https://raw.githubusercontent.com/securego/gosec/master/install.sh | sh -s v2.18.2 ./bin

# xo
brew install xo/xo/xo
