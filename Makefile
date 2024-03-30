.PHONY: build
build:
	docker compose up --build

.PHONY: start
start:
	docker compose up -d

.PHONY: rmi_files
rmi_files:
	rm -rf ./elastic.datas