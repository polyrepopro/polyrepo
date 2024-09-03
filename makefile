docker/build:
	docker build -t cli .

docker/run:
	docker run --entrypoint=sh -it cli

up:
	polyrepo commit --message update
	polyrepo push
	go get -u github.com/polyrepopro/api
	go mod tidy
	polyrepo commit --message "bump api version"
	polyrepo push