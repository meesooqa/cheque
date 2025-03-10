FROM golang:1.23-alpine AS build
COPY . /build
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux go build -o /build/cheque -ldflags "-s -w"

FROM alpine:latest
COPY --from=build /build/cheque /srv/cheque
RUN chmod +x /srv/cheque
RUN apk --no-cache add ca-certificates
WORKDIR /srv
CMD ["/srv/cheque"]