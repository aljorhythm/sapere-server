ENV ?= localhost

lint:
	go fmt ./... || (echo "lint failed $$?"; exit 1)
run: lint
	go run main.go
test: lint
	go test ./... || (echo "test failed $$?"; exit 1)

# containers

SERVICE_HOST ?= localhost:8080

build-image: test
	ENV=$(ENV) sh .build_image.sh
run-container: build-image
	ENV=$(ENV) sh .run_container.sh
service-test:
	(cd service_tests && make test SERVICE_HOST=$(SERVICE_HOST)) || (echo "service-test failed $$?"; exit 1)
stop-container:
	ENV=$(ENV) sh .stop_container.sh
publish-image:
	ENV=$(ENV) sh .publish_image.sh