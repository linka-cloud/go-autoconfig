REGISTRY := "linkacloud"
IMAGE := "autoconfig"

.PHONY:
docker-build:
	@docker image build -t $(REGISTRY)/$(IMAGE) .

.PHONY:
docker-push:
	@docker image push $(REGISTRY)/$(IMAGE)

.PHONY:
docker: docker-build docker-push