download-dependencies:
	go mod download

run: download-dependencies
	go run main.go

docker-build-image:
	docker build -t go-kit-template .

docker-build-image-tests:
	docker build -f Dockerfile-test -t go-kit-template-test .

docker-run-local: docker-build-image
	docker run --rm --env-file .env -p 8080:8080 go-kit-template

previous-app-running:
	go run cmd/other_apps/main.go $(filepath)

tests: docker-build-image-tests
	docker run --rm go-kit-template