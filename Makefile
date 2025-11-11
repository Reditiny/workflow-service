IMAGE_NAME := myapp
TAG := latest
DOCKERFILE := Dockerfile

SERVICE = workflow-worker

.PHONY: help
help:
	@echo "usage："
	@echo "  make build-worker                # build cmd/main.go"

.PHONY: build
build:
	docker build \
		-t $(IMAGE_NAME)-$(SERVICE):$(TAG) \
		-f $(DOCKERFILE) .


.PHONY: build-bank build-webhook build-sqs
build-worker:
	$(MAKE) build