FROM golang:1.17-alpine

WORKDIR /go/src/discgolfapi.com
COPY . .
RUN apk add --update make
RUN go get
RUN go get -u github.com/swaggo/swag/cmd/swag
RUN make build

EXPOSE 8080

ENTRYPOINT ["make", "run-build", "--always-make"]