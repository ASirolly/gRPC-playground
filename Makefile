.PHONY: run-docker
run-docker:
	docker run -p 8090:8090 grpctest:latest ./main
	
.PHONY: build-docker
build-docker:
	docker build -f Dockerfile -t grpctest:latest --build-arg CGO_ENABLED=0 .
	
.PHONY: build
build:
	go build -o ./bin/grpcTest