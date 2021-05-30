v=$(git rev-parse --short HEAD)
tag=$ENV-$v

docker push aljorhythm/sapere-server:$tag
# docker image tag aljorhythm/sapere-server:$tag docker.io/aljorhythm/sapere-server:latest