FROM golang:1.18 as build

WORKDIR /go/src/disc-golf-api
COPY . .

RUN go mod download

RUN CGO_ENABLED=0 go build -o /go/bin/disc-golf-api


FROM gcr.io/distroless/static-debian11

COPY --from=build /go/bin/disc-golf-api /
