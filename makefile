# Check to see if we can use ash in Alpine images of default to BASH
SHELL_PATH = /bin/ash
SHELL = $(if $(wildcard $(SHELL_PATHH)),/bin/ash,/bin/bash)

# ==============================================================================
# Define dependencies

GOLANG          := golang:1.21.5
ALPINE          := alpine:3.19
KIND            := kindest/node:v1.29.0@sha256:eaa1450915475849a73a9227b8f201df25e55e268e5d619312131292e324d570
POSTGRES        := postgres:16.1
VAULT           := hashicorp/vault:1.15
GRAFANA         := grafana/grafana:10.2.0
PROMETHEUS      := prom/prometheus:v2.48.0
TEMPO           := grafana/tempo:2.3.0
LOKI            := grafana/loki:2.9.0
PROMTAIL        := grafana/promtail:2.9.0

KIND_CLUSTER    := ardan-starter-cluster
REPO            := iron2.debotjes.nl
NAMESPACE       := sales-system
APP             := sales
BASE_IMAGE_NAME := ardanlabs/service
SERVICE_NAME    := sales-api
VERSION         := 0.0.1
SERVICE_IMAGE   := $(BASE_IMAGE_NAME)/$(SERVICE_NAME):$(VERSION)
METRICS_IMAGE   := $(BASE_IMAGE_NAME)/$(SERVICE_NAME)-metrics:$(VERSION)


run-local:
	go run app/services/sales-api/main.go

tidy:
	go mod tidy
	go mod vendor

# =========================================================================================================
# Running from within k8s

dev-status:
	kubectl get nodes -o wide
	kubectl get svc -o wide
	kubectl get pods -o wide --watch --all-namespaces

dev-apply:
	kustomize build zarf/k8s/dev/sales |kubectl apply -f -
	#kubectl wait pods --namespace=$(NAMESPACE) --selector app=$(APP) --for=conditon=Ready
	kubectl wait --for=condition=ready --namespace=$(NAMESPACE) --selector app=$(APP) pods

dev-logs:
	kubectl logs --namespace=$(NAMESPACE) -l app=$(APP) --all-containers=true -f --tail=100

dev-restart:
	kubectl rollout restart deployment $(APP) --namespace=$(NAMESPACE)

dev-update: all dev-restart

dev-update-apply: all dev-apply



# =========================================================================================================
# Build containers

all: service

service:
	docker build \
		-f zarf/docker/dockerfile.service \
		-t $(REPO)/$(SERVICE_IMAGE) \
		--build-arg BUILD_REF=$(VERSION) \
		--build-arg BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"` \
		.
	cat ~/.docker/.passwd | docker login -u admin --password-stdin $(REPO)
	docker push $(REPO)/$(SERVICE_IMAGE)


