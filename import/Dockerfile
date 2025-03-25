FROM golang:1.24-alpine AS build
WORKDIR /build
WORKDIR /build
COPY ./go.work ./
COPY ./common ./common
COPY ./db ./db
COPY ./server ./server
COPY ./api ./api
COPY ./tools/db ./tools/db
COPY ./import ./import
WORKDIR /build/import
RUN go mod tidy && go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /build/app -ldflags "-s -w"

FROM alpine:latest
COPY --from=build /build/app /usr/local/bin/app
CMD ["/usr/local/bin/app"]
