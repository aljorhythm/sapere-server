CI_PLAT ?= local

lint:
	go fmt ./... || (echo "lint failed $$?"; exit 1)
run: lint
	go run main.go
test: lint
	go test ./... || (echo "test failed $$?"; exit 1)

# containers

SERVICE_HOST ?= localhost:8080

build-image: test
	CI_PLAT=$(CI_PLAT) sh .build_image.sh
run-container:
	CI_PLAT=$(CI_PLAT) sh .run_container.sh
service-test:
	(cd service_tests && make test SERVICE_HOST=$(SERVICE_HOST)) || (echo "service-test failed $$?"; exit 1)
stop-container:
	CI_PLAT=$(CI_PLAT) sh .stop_container.sh
publish-image:
	CI_PLAT=$(CI_PLAT) sh .publish_image.sh
all: build-image run-container service-test stop-container publish-image