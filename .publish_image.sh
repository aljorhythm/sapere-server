v=$(git rev-parse --short HEAD)
branch=$(git rev-parse --abbrev-ref HEAD)
ci_tag=$CI_PLAT-$v
branchtag=$(echo $branch-latest | sed 's/\//-/g')

docker push aljorhythm/sapere-server:$ci_tag
docker image tag aljorhythm/sapere-server:$tag docker.io/aljorhythm/sapere-server:$ci_tag
docker image tag aljorhythm/sapere-server:$tag docker.io/aljorhythm/sapere-server:$branchtag