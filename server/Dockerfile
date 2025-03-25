FROM golang:1.24-alpine AS build
WORKDIR /build
COPY ./common /build/common
COPY ./api /build/api
COPY ./db /build/db
COPY ./server /build/server
WORKDIR /build/server
RUN go mod tidy && go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /build/app -ldflags "-s -w"

FROM alpine:latest
COPY --from=build /build/app /srv/app
RUN chmod +x /srv/app
RUN apk --no-cache add ca-certificates
WORKDIR /srv
CMD ["/srv/app"]