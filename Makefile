.PHONY: run

build:
	docker build -t app .
run:
	docker run -p 8080:8080 app
	