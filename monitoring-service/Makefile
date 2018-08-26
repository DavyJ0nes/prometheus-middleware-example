.DEFAULT_TARGET=help
all: help

# VARIABLES
USERNAME = davyj0nes
APP_NAME = monitoring
NAMESPACE = monitoring

PROM_PORT = 9090
GRAFANA_PORT = 3000

# COMMANDS

## get_pods: prints information about runnning kubernetes pods
.PHONY: get_pods
get_pods:
	$(call blue, "# Getting Running Monitoring Pods...")
	@kubectl get pods -n monitoring

## deploy_prom: deploys prometheus
.PHONY: deploy_prom
deploy_prom:
	$(call blue, "# Deploying Prometheus...")
	@kubectl create -f prometheus/namespace.yml
	@kubectl create -f prometheus/clusterrole.yml
	@kubectl create -f prometheus/clusterrolebinding.yml
	@kubectl create -f prometheus/configmap.yml
	@kubectl create -f prometheus/deployment.yml

## access_prom: port forwards to prometheus
.PHONY: access_prom
access_prom:
	$(call blue, "# Port Forwarding Prometheus...")
	@kubectl port-forward -n "${NAMESPACE}" $(shell kubectl get pods -n monitoring | grep prometheus-deployment | awk '{print $$1}') "${PROM_PORT}:${PROM_PORT}"

## access-grafana: port forwards to grafana
.PHONY: access_grafana
access_grafana:
	$(call blue, "# Port Forwarding Grafana...")
	@kubectl port-forward -n "${NAMESPACE}" $(shell kubectl get pods -n monitoring | grep grafana | awk '{print $$1}') "${GRAFANA_PORT}:${GRAFANA_PORT}"

## help: Show this help message
.PHONY: help
help: Makefile
	@echo "${APP_NAME}"
	@echo
	@echo " Choose a command run in "$(APP_NAME)":"
	@echo
	@sed -n 's/^## //p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo

# FUNCTIONS
define blue
	@tput setaf 4
	@echo $1
	@tput sgr0
endef
