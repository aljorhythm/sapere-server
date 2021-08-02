CI_PLAT ?= local

lint:
	go fmt ./... || (echo "lint failed $$?"; exit 1)
run: lint
	go run .
test: lint
	go test $$(go list ./... | grep -v /servicetests) || (echo "test failed $$?"; exit 1)

# containers

SERVICE_HOST ?= http://localhost:8080

build-info:
	printf $$(git rev-parse --short HEAD) > .commit
	printf $$(git rev-parse --abbrev-ref HEAD) > .branch
	printf $(CI_PLAT) > .ci_plat
build-image: test
	CI_PLAT=$(CI_PLAT) sh .build_image.sh || (echo "build-image failed $$?"; exit 1)
run-container:
	CI_PLAT=$(CI_PLAT) sh .run_container.sh || (echo "run-container failed $$?"; exit 1)
service-test:
	(cd servicetests && make test SERVICE_HOST=$(SERVICE_HOST)) || (echo "service-test failed $$?"; exit 1)
	SERVICE_HOST=$(SERVICE_HOST) go test $$(go list ./... | grep /servicetests) || (echo "test failed $$?"; exit 1)
stop-container:
	CI_PLAT=$(CI_PLAT) sh .stop_container.sh || (echo "stop-container failed $$?"; exit 1)
publish-image:
	CI_PLAT=$(CI_PLAT) sh .publish_image.sh || (echo "publish-image failed $$?"; exit 1)
all: build-image run-container service-test stop-container publish-image