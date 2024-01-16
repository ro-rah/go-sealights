# go-sealights

# build
$ docker build -t go-sealights --no-cache --progress=plain --build-arg SEALIGHTS_TOKEN=<yourtoken> --build-arg APP_NAME=go-sealights --build-arg BRANCH_NAME=main --build-arg BUILD_NAME=<buildno> .

# run
docker run -p 8080:8080 docker.io/library/go-sealights  

# run with SEALIGHTS enabled
docker run -e SEALIGHTS_DISABLE_ON_INIT=false -e SEALIGHTS_LAB_ID=docker_go_sealights -p 8080:8080 docker.io/library/go-sealights

# endpoint access
http://localhost:8080/method<A-J>
ex. for method C:
http://localhost:8080/methodC
