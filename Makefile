CURRENT_DIR=$(shell pwd)

APP=$(shell basename ${CURRENT_DIR})
APP_CMD_DIR=${CURRENT_DIR}/cmd

TAG=latest
ENV_TAG=latest
DOCKERFILE=Dockerfile

proto-gen:
	bash scripts/gen-proto.sh ${CURRENT_DIR} && bash scripts/remove_omitempty_tag.sh

copy-protos:
	rm -rf protos/* && cp -R eld_protos/* protos

pull-proto-module:
	git submodule update --init --recursive

update-proto-module:
	git submodule update ``--remote --merge

build:
	go build -mod=vendor -a -installsuffix cgo -o ${CURRENT_DIR}/bin/${APP} ${APP_CMD_DIR}/main.go

build-image:
	docker build --rm -t ${REGISTRY}/${PROJECT_NAME}/${APP}:${TAG} . -f ${DOCKERFILE}
	docker tag ${REGISTRY}/${PROJECT_NAME}/${APP}:${TAG} ${REGISTRY}/${PROJECT_NAME}/${APP}:${ENV_TAG}

clear-image:
	docker rmi ${REGISTRY}/${PROJECT_NAME}/${APP}:${TAG}
	docker rmi ${REGISTRY}/${PROJECT_NAME}/${APP}:${ENV_TAG}

push-image:
	docker push ${REGISTRY}/${PROJECT_NAME}/${APP}:${TAG}
	docker push ${REGISTRY}/${PROJECT_NAME}/${APP}:${ENV_TAG}

run:
	go run cmd/main.go

linter:
	golangci-lint run

swag_init:
	swag init -g api/main.go -o api/docs

stop: 
	docker-compose down

start: stop
	docker-compose up -d --build