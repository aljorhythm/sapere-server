v=$(git rev-parse --short HEAD)
branch=$(git rev-parse --abbrev-ref HEAD)
tag=$ENV-$v
branchtag=$branch-latest

docker push aljorhythm/sapere-server:$tag
docker image tag aljorhythm/sapere-server:$tag docker.io/aljorhythm/sapere-server:$envtag
docker image tag aljorhythm/sapere-server:$tag docker.io/aljorhythm/sapere-server:$branchtag