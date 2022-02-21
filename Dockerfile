ARG build_path=/go/bin/discgolfapi

FROM golang:1.17-alpine as build-env
ARG build_path

WORKDIR /go/src/discgolfapi.com

# cache and install dependencies
COPY go.mod . 
COPY go.sum .
RUN go mod download

COPY . .

# build docs
RUN go get -u github.com/swaggo/swag/cmd/swag
RUN swag init

# build server
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/discgolfapi

FROM scratch
ARG build_path

# run server
COPY --from=build-env $build_path "/go/bin/discgolfapi"
EXPOSE 8080
ENTRYPOINT ["/go/bin/discgolfapi", "-port=8080"]
