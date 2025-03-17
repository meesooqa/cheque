FROM golang:1.23-alpine AS build
WORKDIR /build
COPY ./common /build/common
COPY ./import /build/import
WORKDIR /build/import
RUN go mod tidy && go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /build/app -ldflags "-s -w"

FROM alpine:latest
COPY --from=build /build/app /usr/local/bin/app
CMD ["/usr/local/bin/app"]
