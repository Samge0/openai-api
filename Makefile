.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w' -o openai-api ./main.go

.PHONY: docker
docker:
	docker build . -t openai-api:v1
