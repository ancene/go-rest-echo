# variable
PORT					= 40000
DOCKER_IMAGE_NAME		= "goapp"
DOCKER_CONTAINER_NAME	= "goapp"

# command docker
build: clear
	@docker build --build-arg TAGGED=builder-${DOCKER_IMAGE_NAME} -f Dockerfile --tag $(DOCKER_IMAGE_NAME) .

run: clear
	@docker run --restart=always -d --name $(DOCKER_CONTAINER_NAME) -p $(PORT):$(PORT) $(DOCKER_IMAGE_NAME)

# development
clear:
	@clear

dev: clear
	@go run cmd/main.go

watch: clear
	@reflex -r '\.go' -s -- sh -c "go run cmd/main.go"

test: clear
	@go test ./internal/... ./pkg/...

cover: clear
	@go test ./internal/... ./pkg/... -coverprofile=cover.out
	@go tool cover -html=cover.out

lint: clear
	@staticcheck ./...

gen-mock:
	@rm -rf mocks/*.go
	@mockery --dir=internal/application/port --all

# @docker build --build-arg TAGGED=builder-${DOCKER_IMAGE_NAME} --file Dockerfile --tag $(DOCKER_IMAGE_NAME) .