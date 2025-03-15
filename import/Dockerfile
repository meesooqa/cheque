FROM golang:1.23-alpine AS build
COPY . /build
WORKDIR /build
RUN go get -u . && go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o /build/import -ldflags "-s -w"

FROM alpine:latest
COPY --from=build /build/import /usr/local/bin/import
CMD ["/usr/local/bin/import"]
