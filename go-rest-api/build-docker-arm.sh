CGO_ENABLED=0 GOOS=linux go build -a -o restapi *.go
docker build . -t togoetha/go-exe-arm
docker save -o ../container/go-exe-arm.tar togoetha/go-exe-arm
