default: docker_test

test:
	go test -v

docker_test:
	docker build .

