docker run --rm -v $(pwd):/app -w /app golang:1.17 go run server.go

docker run --rm -it -v $(pwd):/app -w /app golang:1.17 /bin/bash