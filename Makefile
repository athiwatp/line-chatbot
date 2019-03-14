.PHONY : build test docker

build:
	go build -o bin

test:
	go test -v

docker:
	if ! [ -f .env ]; then cp env.sample .env; fi;
	docker build -t line-chatbot:latest .