v=$(git rev-parse --short HEAD)
tag=$CI_PLAT-$v

docker run -d -p 8080:8080 --name sapere-server --rm aljorhythm/sapere-server:$tag