
build-openapi:
	docker run --rm -v "/${PWD}:/disc-golf-api" openapitools/openapi-generator-cli generate \
		-i ./disc-golf-api/openapi.yaml \
		-g go-server \
		--additional-properties=outputAsLibrary=true,sourceFolder=openapi \
		-o ./disc-golf-api/pkg/gen

build-app:
	docker build -t discgolfapi .

run-server: build-app
	docker run -t discgolfapi ./disc-golf-api server