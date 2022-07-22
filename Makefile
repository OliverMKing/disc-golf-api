
build-openapi:
	docker run --rm -v "/${PWD}:/curr" openapitools/openapi-generator-cli generate \
		-i ./curr/openapi.yaml \
		-g go-server \
		-o ./curr/gen
