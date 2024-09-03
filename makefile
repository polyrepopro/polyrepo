docker/build:
	docker build -t cli .

docker/run:
	docker run --entrypoint=sh -it cli

up:
	polyrepo workspace commit --message update
	polyrepo workspace push
	go get -u github.com/polyrepopro/api
	go mod tidy
	polyrepo workspace commit --message "bump api version"
	polyrepo workspace push