v=$(git rev-parse --short HEAD)
tag=$ENV-$v

docker run -d -p 8080:8080 --name sapere-server --rm aljorhythm/sapere-server:$tag