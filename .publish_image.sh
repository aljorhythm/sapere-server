v=$(git rev-parse --short HEAD)
tag=$ENV-$v

docker push aljorhythm/sapere:$tag