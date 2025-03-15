FROM golang:1.23-alpine AS build

WORKDIR /build
COPY ./common /build/common
COPY ./api /build/api
COPY ./api-server /build/api-server

WORKDIR /build/api-server
RUN go mod tidy && go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /build/cheque -ldflags "-s -w"

FROM alpine:latest
COPY --from=build /build/cheque /srv/cheque
RUN chmod +x /srv/cheque
RUN apk --no-cache add ca-certificates
WORKDIR /srv
CMD ["/srv/cheque"]