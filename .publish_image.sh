v=$(git rev-parse --short HEAD)
branch=$(git rev-parse --abbrev-ref HEAD)
tag=$CI_PLAT-$v
branchtag=$(echo $branch-latest | sed 's/\//-/g')

docker push aljorhythm/sapere-server:$tag
docker image tag aljorhythm/sapere-server:$tag docker.io/aljorhythm/sapere-server:$envtag
docker image tag aljorhythm/sapere-server:$tag docker.io/aljorhythm/sapere-server:$branchtag