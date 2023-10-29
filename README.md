build memo
docker build -t backend -f backend/build/Dockerfile .

docker run --restart=always --name=backend -v $(pwd)/backend:/go/src/backend -p 5678:5678 -e GOPATH=${GOPATH} -t backend




# windows
GOOS=windows GOARCH=amd64 go build -o myapp.exe


GOOS=linux GOARCH=amd64 go build -o myapp
