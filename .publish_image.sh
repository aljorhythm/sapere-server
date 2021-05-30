v=$(git rev-parse --short HEAD)
branch=$(git rev-parse --abbrev-ref HEAD)
tag=$CI_PLAT.$v
branch_latest_tag=$(echo $CI_PLAT.$branch.latest | sed 's/\//-/g')
branch_commit_tag=$(echo $CI_PLAT.$branch.$v | sed 's/\//-/g')

docker tag aljorhythm/sapere-server:$tag docker.io/aljorhythm/sapere-server:$branch_latest_tag
docker tag aljorhythm/sapere-server:$tag docker.io/aljorhythm/sapere-server:$branch_commit_tag

docker push aljorhythm/sapere-server:$tag
docker push aljorhythm/sapere-server:$branch_latest_tag
docker push aljorhythm/sapere-server:$branch_commit_tag
