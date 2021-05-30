export DOCKER_BUILDKIT=1

v=$(git rev-parse --short HEAD)
tag=$CI_PLAT.$v

docker build --no-cache -t aljorhythm/sapere-server:$tag .