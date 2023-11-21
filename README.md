# go-sealights

# build
$ docker build -t go-sealights --build-arg SEALIGHTS_TOKEN=<yourtoken> --build-arg APP_NAME=go-sealights --build-arg BRANCH_NAME=main --build-arg BUILD_NAME=<buildno> .

# run
docker run -p 8080:8080 docker.io/library/go-sealights  
