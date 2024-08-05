docker/build:
	docker build -t cli .

docker/run:
	docker run --entrypoint=sh -it cli