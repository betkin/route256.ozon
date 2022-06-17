
.PHONY: load-example
load-example:
	docker run --network device-api_ompnw -v $(shell pwd):/var/loadtest direvius/yandex-tank -c ./load.yaml

.PHONY: grafana-up
grafana-up:
	docker run --network device-api_ompnw -d -p 3000:3000 grafana/grafana

.PHONY: prometheus-up
prometheus-up:
	docker run --network device-api_ompnw -p 9090:9090 prom/prometheus